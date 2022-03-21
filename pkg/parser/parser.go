package parser

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"

	"github.com/flier/gocombine/pkg/stream"
)

// Parser is a type that it can be used to parse an input stream `S` of token `T` into a `O` value.
type Parser[T stream.Token, O any] interface {
	Parse(input []T) (out O, remaining []T, err error)
}

// Func is a function that it can be used to parse an input stream `S` of token `T` into a `O` value.
type Func[T stream.Token, O any] func(input []T) (out O, remaining []T, err error)

func (f Func[T, O]) Parse(input []T) (out O, remaining []T, err error) { return f(input) }

// Expected parses with `f` and if it fails without consuming any input any expected errors are replaced by `msg`.
// `msg` is then used in error messages as "Expected `msg`".
func (f Func[T, O]) Expected(msg string) Func[T, O] {
	return Expected(f, msg)
}

// Expected parses with `parser` and if it fails without consuming any input any expected errors are replaced by `msg`.
// `msg` is then used in error messages as "Expected `msg`".
func Expected[
	T stream.Token,
	O any,
](
	parser Func[T, O], msg string,
) Func[T, O] {
	return func(input []T) (parsed O, remaining []T, err error) {
		parsed, remaining, err = parser(input)
		if err != nil {
			err = fmt.Errorf("%s, %w", msg, err)
		}

		return
	}
}

// Message parses with `f` and if it fails, adds the message `msg` to the error.
func (f Func[T, O]) Message(msg string) Func[T, O] {
	return Message(f, msg)
}

// Message parses with `parser` and if it fails, adds the message `msg` to the error.
func Message[

	T stream.Token,
	O any,
](
	parser Func[T, O], msg string,
) Func[T, O] {
	return func(input []T) (parsed O, remaining []T, err error) {
		parsed, remaining, err = parser(input)
		if err != nil {
			err = multierror.Append(err, errors.New(msg))
		}

		return
	}
}
