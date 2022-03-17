package parser

import (
	"fmt"

	"github.com/flier/gocombine/pkg/stream"
)

type Parser[S stream.Stream[T], T stream.Token, O any] interface {
	Parse(input S) (out O, remaining S, err error)

	Expected(msg string) Func[S, T, O]
}

type Func[S stream.Stream[T], T stream.Token, O any] func(input S) (out O, remaining S, err error)

func (f Func[S, T, O]) Parse(input S) (out O, remaining S, err error) { return f(input) }

// Expected parses with `f` and if it fails without consuming any input any expected errors are replaced by `msg`.
// `msg` is then used in error messages as "Expected `msg`".
func (f Func[S, T, O]) Expected(msg string) Func[S, T, O] {
	return Expected[S, T, O](f, msg)
}

// Expected parses with `parser` and if it fails without consuming any input any expected errors are replaced by `msg`.
// `msg` is then used in error messages as "Expected `msg`".
func Expected[
	S stream.Stream[T],
	T stream.Token,
	O any,
	P Parser[S, T, O],
](parser P, msg string) Func[S, T, O] {
	return func(input S) (parsed O, remaining S, err error) {
		parsed, remaining, err = parser.Parse(input)
		if err != nil {
			err = fmt.Errorf("%s, %w", msg, err)
		}
		return
	}
}
