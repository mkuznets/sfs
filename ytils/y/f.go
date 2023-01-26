package y

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/constraints"
	"mkuznets.com/go/sfs/ytils/yerr"

	"io"
)

func Min[T constraints.Ordered](a T, b ...T) T {
	m := a
	for _, v := range b {
		if v < m {
			m = v
		}
	}
	return m
}

func Max[T constraints.Ordered](a T, b ...T) T {
	m := a
	for _, v := range b {
		if v > m {
			m = v
		}
	}
	return m
}

func Close(r io.Closer) {
	if err := r.Close(); err != nil {
		log.Debug().Stack().Err(yerr.New("Close error").Err(err)).Send()
	}
}
