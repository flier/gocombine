package ranges

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Take reads a range of length `n`.
func Take[S stream.Stream[T], T stream.Token](n int) parser.Func[S, T, []T] {
	return func(input S) (out []T, remaining S, err error) {
		return stream.UnconsRange(input, n)
	}
}

// TakeWhile reads a range of 0 or more tokens which satisfy `f`.
func TakeWhile[S stream.Stream[T], T stream.Token](f func(T) bool) parser.Func[S, T, []T] {
	return func(input S) (out []T, remaining S, err error) {
		return stream.UnconsWhile(input, f)
	}
}

// TakeWhile1 reads a range of 1 or more tokens which satisfy `f`.
func TakeWhile1[S stream.Stream[T], T stream.Token](f func(T) bool) parser.Func[S, T, []T] {
	return func(input S) (out []T, remaining S, err error) {
		out, remaining, err = stream.UnconsWhile(input, f)
		if err == nil && len(out) == 0 {
			err = fmt.Errorf("one or more elements, %w", parser.ErrExpected)
		}

		return
	}
}
