package yctx

import (
	"context"
	"github.com/dlsniper/debugger"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

type Contexts interface {
	Normal() context.Context
	Critical() context.Context
}

type contextsImpl struct {
	normal   context.Context
	critical context.Context
}

func NewContexts() Contexts {
	parent := context.Background()
	normal, normalCancel := context.WithCancel(parent)
	critical, criticalCancel := context.WithCancel(parent)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		debugger.SetLabels(func() []string {
			return []string{
				"pkg", "ytils/yctx",
				"func", "cotexts signal handler",
			}
		})

		cnt := 0
		for range signalChan {
			switch cnt {
			case 0:
				log.Debug().Msg("graceful exit")
				normalCancel()
			case 1:
				log.Debug().Msg("send one more for hard exit")
				criticalCancel()
			default:
				log.Debug().Msg("hard exit")
				os.Exit(1)
			}
			cnt++
		}
	}()

	return &contextsImpl{
		normal:   normal,
		critical: critical,
	}
}

func (c *contextsImpl) Normal() context.Context {
	return c.normal
}

func (c *contextsImpl) Critical() context.Context {
	return c.critical
}
