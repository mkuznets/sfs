package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

func foo() error {
	var f interface{}
	err := json.Unmarshal([]byte(`"`), &f)
	if err != nil {
		return errors.Wrap(err, "invalid json")
	}
	return nil
}

func bar() error {
	err := foo()
	if err != nil {
		return errors.Wrap(err, "foo failed")
	}
	return nil
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	err := bar()

	e := err
	var deepestStack errors.StackTrace
	for e != nil {
		if st, ok := e.(stackTracer); ok {
			deepestStack = st.StackTrace()
		}
		e = errors.Unwrap(e)
	}

	for _, f := range deepestStack {
		fmt.Printf("%+s:%d\n", f, f)
	}
}
