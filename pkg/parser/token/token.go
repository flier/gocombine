package token

import (
	"fmt"
	"io"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

type PeekErr[T stream.Token] struct {
	Expected T
	Actual   T
}

func (e *PeekErr[T]) Error() string {
	return fmt.Sprintf("expected `%c`, got `%c`", e.Expected, e.Actual)
}

func Token[S stream.Stream[T], T stream.Token, P parser.ParseFunc[S, T, T]](t T) P {
	return func(s S) (c T, r S, err error) {
		if len(s) == 0 {
			err = io.ErrUnexpectedEOF
		} else if s[0] == t {
			c, r = s[0], s[1:]
		} else {
			r = s
			err = &PeekErr[T]{t, s[0]}
		}
		return
	}
}
