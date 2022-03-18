package token

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Value always returns the value `v` without consuming any input.
func Value[S stream.Stream[T], T stream.Token](v T) parser.Func[S, T, T] {
	return func(input S) (T, S, error) {
		return v, input, nil
	}
}

// Produce always returns the value produced by calling `f`.
func Produce[S stream.Stream[T], T stream.Token, O any](f func() O) parser.Func[S, T, O] {
	return func(input S) (O, S, error) {
		return f(), input, nil
	}
}
