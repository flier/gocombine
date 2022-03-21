package token

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Value always returns the value `v` without consuming any input.
func Value[T stream.Token, O any](v O) parser.Func[T, O] {
	return func(input []T) (O, []T, error) {
		return v, input, nil
	}
}

// Produce always returns the value produced by calling `f`.
func Produce[T stream.Token, O any](f func() O) parser.Func[T, O] {
	return func(input []T) (O, []T, error) {
		return f(), input, nil
	}
}
