package ylog

import (
	"github.com/pkg/errors"
	"mkuznets.com/go/sfs/internal/ytils/yerr"
)

var (
	StackSourceFileName     = "source"
	StackSourceLineName     = "line"
	StackSourceFunctionName = "func"
)

type state struct {
	b []byte
}

// Write implement fmt.Formatter interface.
func (s *state) Write(b []byte) (n int, err error) {
	s.b = b
	return len(b), nil
}

// Width implement fmt.Formatter interface.
func (s *state) Width() (wid int, ok bool) {
	return 0, false
}

// Precision implement fmt.Formatter interface.
func (s *state) Precision() (prec int, ok bool) {
	return 0, false
}

// Flag implement fmt.Formatter interface.
func (s *state) Flag(c int) bool {
	return false
}

func frameField(f errors.Frame, s *state, c rune) string {
	f.Format(s, c)
	return string(s.b)
}

// MarshalStack implements pkg/errors stack trace marshaling.
func MarshalStack(err error) interface{} {
	e := err
	var st errors.StackTrace
	for e != nil {
		if errS, ok := e.(yerr.StackTracer); ok {
			st = errS.StackTrace()
		}
		e = errors.Unwrap(e)
	}
	if st == nil {
		return nil
	}

	s := &state{}
	out := make([]map[string]string, 0, len(st))
	for _, frame := range st {
		out = append(out, map[string]string{
			StackSourceFileName:     frameField(frame, s, 's'),
			StackSourceLineName:     frameField(frame, s, 'd'),
			StackSourceFunctionName: frameField(frame, s, 'n'),
		})
	}
	return out
}
