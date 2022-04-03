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
func Expected[T stream.Token, O any](parser Func[T, O], msg string) Func[T, O] {
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
func Message[T stream.Token, O any](parser Func[T, O], msg string) Func[T, O] {
	return func(input []T) (parsed O, remaining []T, err error) {
		parsed, remaining, err = parser(input)
		if err != nil {
			err = multierror.Append(err, errors.New(msg))
		}

		return
	}
}

// Map uses `fn` to map over the parsed value.
func (f Func[T, O]) Map(fn func(O) O) Func[T, O] {
	return func(input []T) (parsed O, remaining []T, err error) {
		var o O

		if o, remaining, err = f(input); err != nil {
			return
		}

		parsed = fn(o)

		return
	}
}

// MapErr uses `f` to map over the parser `p` error.
func MapErr[T stream.Token, O any](p Func[T, O], f func(error) error) Func[T, O] {
	return func(input []T) (parsed O, remaining []T, err error) {
		parsed, remaining, err = p(input)
		if err != nil {
			err = f(err)
		}

		return
	}
}

// MapErr uses `fn` to map over the error.
func (f Func[T, O]) MapErr(fn func(error) error) Func[T, O] {
	return MapErr(f, fn)
}

// AndThen parses with `f` and applies `fn` on the result if `parser` parses successfully.
// `fn` may optionally fail with an error.
func (f Func[T, O]) AndThen(fn func(O) (O, error)) Func[T, O] {
	return func(input []T) (parsed O, remaining []T, err error) {
		var o O

		if o, remaining, err = f(input); err != nil {
			return
		}

		if parsed, err = fn(o); err != nil {
			remaining = input
		}

		return
	}
}
