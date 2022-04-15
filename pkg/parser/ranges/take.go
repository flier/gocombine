package ranges

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Take reads a range of length `n`.
func Take[T stream.Token](n int) parser.Func[T, []T] {
	return parser.Expected(func(input []T) (out []T, remaining []T, err error) {
		return stream.UnconsRange(input, n)
	}, "take")
}

// TakeWhile reads a range of 0 or more tokens which satisfy `f`.
func TakeWhile[T stream.Token](f func(T) bool) parser.Func[T, []T] {
	return parser.Expected(func(input []T) (out []T, remaining []T, err error) {
		return stream.UnconsWhile(input, f)
	}, "take while")
}

// TakeWhile1 reads a range of 1 or more tokens which satisfy `f`.
func TakeWhile1[T stream.Token](f func(T) bool) parser.Func[T, []T] {
	return parser.Expected(func(input []T) (out []T, remaining []T, err error) {
		out, remaining, err = stream.UnconsWhile(input, f)
		if err == nil && len(out) == 0 {
			err = fmt.Errorf("one or more elements, %w", parser.ErrExpected)
		}

		return
	}, "take while1")
}

// TakeUntil reads a range of 0 or more tokens until token `r` which satisfy `f`.
//
// The range `r` will not be committed. If `r` is not found, the parser will return an error.
func TakeUntil[T stream.Token](f func(T) bool) parser.Func[T, []T] {
	return parser.Expected(func(input []T) (out []T, remaining []T, err error) {
		return stream.UnconsUntil(input, f)
	}, "take until")
}

// TakeUntil reads a range of 1 or more tokens until token `r` which satisfy `f`.
//
// The range `r` will not be committed. If `r` is not found, the parser will return an error.
func TakeUntil1[T stream.Token](f func(T) bool) parser.Func[T, []T] {
	return parser.Expected(func(input []T) (out []T, remaining []T, err error) {
		out, remaining, err = stream.UnconsUntil(input, f)
		if err == nil && len(out) == 0 {
			err = fmt.Errorf("one or more elements, %w", parser.ErrExpected)
		}

		return
	}, "take until")
}
