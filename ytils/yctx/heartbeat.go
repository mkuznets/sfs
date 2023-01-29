package yctx

import (
	"context"
	"github.com/dlsniper/debugger"
	"github.com/rs/zerolog/log"
	"mkuznets.com/go/sfs/ytils/y"
	"mkuznets.com/go/sfs/ytils/ytime"
	"time"
)

const (
	DefaultCheckInterval = 10 * time.Second
	DefaultLeftWarning   = time.Minute
)

type Heartbeat struct {
	beatC         chan bool
	ctx           context.Context
	cancel        context.CancelFunc
	timeout       time.Duration
	checkInterval time.Duration
	leftWarning   time.Duration
	name          string
}

func NewHeartbeat(ctx context.Context, timeout time.Duration) *Heartbeat {
	ctx, cancel := context.WithCancel(ctx)
	return &Heartbeat{
		ctx:           ctx,
		cancel:        cancel,
		timeout:       timeout,
		checkInterval: DefaultCheckInterval,
		leftWarning:   DefaultLeftWarning,
		beatC:         make(chan bool),
	}
}

func (h *Heartbeat) Context() context.Context {
	return h.ctx
}

func (h *Heartbeat) Beat() {
	select {
	case h.beatC <- true:
	default:
	}
}

func (h *Heartbeat) Close() {
	h.cancel()
	close(h.beatC)
}

func (h *Heartbeat) WithName(name string) *Heartbeat {
	h.name = name
	return h
}

func (h *Heartbeat) Start() *Heartbeat {
	lastBeat := time.Now()

	logger := log.Logger
	if h.name != "" {
		logger = logger.With().Str("ctx", h.name).Logger()
	}

	go func(last *time.Time) {
		debugger.SetLabels(func() []string {
			return []string{
				"pkg", "ytils/yctx",
				"name", h.name,
				"func", "beats monitor",
			}
		})

		for {
			if h.ctx.Err() != nil {
				return
			}

			idle := time.Since(*last)
			logger.Debug().Stringer("elapsed", idle.Round(time.Millisecond)).Msg("last heartbeat")

			if idle >= h.timeout {
				logger.Error().Msg("idle context cancelled")
				h.cancel()
				return
			}

			left := h.timeout - idle
			if left <= DefaultLeftWarning {
				logger.Warn().Stringer("left", left.Round(time.Millisecond)).Msg("idle context")
			}

			ytime.Sleep(h.ctx, y.Min(DefaultCheckInterval, left))
		}
	}(&lastBeat)

	go func() {
		debugger.SetLabels(func() []string {
			return []string{
				"pkg", "ytils/yctx",
				"name", h.name,
				"func", "beats consumer",
			}
		})

		for {
			select {
			case <-h.beatC:
				lastBeat = time.Now()
			case <-h.ctx.Done():
				return
			}
		}
	}()

	return h
}
