package parser

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"

	"github.com/flier/gocombine/pkg/stream"
)

// Parser is a type that it can be used to parse an input stream `S` of token `T` into a `O` value.
type Parser[S stream.Stream[T], T stream.Token, O any] interface {
	Parse(input S) (out O, remaining S, err error)
}

// Func is a function that it can be used to parse an input stream `S` of token `T` into a `O` value.
type Func[S stream.Stream[T], T stream.Token, O any] func(input S) (out O, remaining S, err error)

func (f Func[S, T, O]) Parse(input S) (out O, remaining S, err error) { return f(input) }

// Expected parses with `f` and if it fails without consuming any input any expected errors are replaced by `msg`.
// `msg` is then used in error messages as "Expected `msg`".
func (f Func[S, T, O]) Expected(msg string) Func[S, T, O] {
	return Expected(f, msg)
}

// Expected parses with `parser` and if it fails without consuming any input any expected errors are replaced by `msg`.
// `msg` is then used in error messages as "Expected `msg`".
func Expected[
	S stream.Stream[T],
	T stream.Token,
	O any,
](
	parser Func[S, T, O], msg string,
) Func[S, T, O] {
	return func(input S) (parsed O, remaining S, err error) {
		parsed, remaining, err = parser(input)
		if err != nil {
			err = fmt.Errorf("%s, %w", msg, err)
		}

		return
	}
}

// Message parses with `f` and if it fails, adds the message `msg` to the error.
func (f Func[S, T, O]) Message(msg string) Func[S, T, O] {
	return Message(f, msg)
}

// Message parses with `parser` and if it fails, adds the message `msg` to the error.
func Message[
	S stream.Stream[T],
	T stream.Token,
	O any,
](
	parser Func[S, T, O], msg string,
) Func[S, T, O] {
	return func(input S) (parsed O, remaining S, err error) {
		parsed, remaining, err = parser(input)
		if err != nil {
			err = multierror.Append(err, errors.New(msg))
		}

		return
	}
}
