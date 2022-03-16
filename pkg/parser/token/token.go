package token

import (
	"fmt"
	"io"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

type UnexpectedErr[T stream.Token] struct {
	Expected []T
	Actual   []T
}

func (e *UnexpectedErr[T]) Error() string {
	if e.Expected != nil {
		return fmt.Sprintf("expected `%c`, got `%c`", e.Expected, e.Actual)
	}

	return fmt.Sprintf("unexpected `%c`", e.Actual)
}

func Unexpected[T stream.Token](expected, actual []T) error {
	return &UnexpectedErr[T]{expected, actual}
}

// Any parses any token.
func Any[S stream.Stream[T], T stream.Token, P parser.ParseFunc[S, T, T]]() P {
	return func(input S) (actual T, remaining S, err error) {
		if len(input) == 0 {
			err = io.ErrUnexpectedEOF
		} else {
			actual, remaining = input[0], input[1:]
		}

		return
	}
}

// Token parses a character and succeeds if the character is equal to `tok`.
func Token[S stream.Stream[T], T stream.Token, P parser.ParseFunc[S, T, T]](tok T) P {
	return func(input S) (actual T, remaining S, err error) {
		if len(input) == 0 {
			err = io.ErrUnexpectedEOF
		} else if input[0] == tok {
			actual, remaining = input[0], input[1:]
		} else {
			actual, remaining = input[0], input

			err = Unexpected([]T{tok}, []T{actual})
		}
		return
	}
}

// Tokens parses multiple tokens.
//
// Consumes items from the input and compares them to the values from `tokens` using the
/// comparison function `cmp`. Succeeds if all the items from `tokens` are matched in the input
/// stream and fails otherwise with `expected` used as part of the error.
func Tokens[S stream.Stream[T], T stream.Token, P parser.ParseFunc[S, T, []T]](cmp func(lhs, rhs T) bool, expected, tokens S) P {
	return func(input S) (actual []T, remaining S, err error) {
		n := len(tokens)
		if len(input) < n {
			err = io.ErrUnexpectedEOF
		} else {
			for i, tok := range tokens {
				if !cmp(input[i], tok) {
					err = Unexpected(expected, input[:n])
					return
				}
			}

			actual, remaining = input[:n], input[n:]
		}
		return
	}
}

// OneOf extract one token and succeeds if it is part of `tokens`.
func OneOf[S stream.Stream[T], T stream.Token, P parser.ParseFunc[S, T, T]](tokens S) P {
	return func(input S) (actual T, remaining S, err error) {
		if len(input) == 0 {
			err = io.ErrUnexpectedEOF
		} else {
			for _, tok := range tokens {
				if input[0] == tok {
					actual, remaining = input[0], input[1:]
					return
				}
			}

			actual, remaining = input[0], input

			err = Unexpected(tokens, input[:1])
		}
		return
	}
}

// NoneOf extract one token and succeeds if it is not part of `tokens`.
func NoneOf[S stream.Stream[T], T stream.Token, P parser.ParseFunc[S, T, T]](tokens S) P {
	return func(input S) (actual T, remaining S, err error) {
		if len(input) == 0 {
			err = io.ErrUnexpectedEOF
		} else {
			for _, tok := range tokens {
				if input[0] == tok {
					actual, remaining = input[0], input
					err = Unexpected(nil, input[:1])
					return
				}
			}

			actual, remaining = input[0], input[1:]
		}
		return
	}
}

// Value always returns the value `v` without consuming any input.
func Value[S stream.Stream[T], T stream.Token, P parser.ParseFunc[S, T, T]](v T) P {
	return func(input S) (T, S, error) {
		return v, input, nil
	}
}

// Produce always returns the value produced by calling `f`.
func Produce[S stream.Stream[T], T stream.Token, O any, P parser.ParseFunc[S, T, O]](f func() O) P {
	return func(input S) (O, S, error) {
		return f(), input, nil
	}
}

// Eof succeeds only if the stream is at end of input, fails otherwise.
func Eof[S stream.Stream[T], T stream.Token, P parser.ParseFunc[S, T, bool]]() P {
	return func(input S) (bool, S, error) {
		if len(input) == 0 {
			return true, input, nil
		}

		return false, input, Unexpected(nil, input)
	}
}
