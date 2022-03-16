package token

import (
	"fmt"
	"io"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

type UnexpectedErr[T stream.Token] struct {
	Expected T
	Actual   T
}

func (e *UnexpectedErr[T]) Error() string {
	return fmt.Sprintf("expected `%c`, got `%c`", e.Expected, e.Actual)
}

func Unexpected[T stream.Token](expected, actual T) error {
	return &UnexpectedErr[T]{expected, actual}
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

			err = Unexpected(tok, actual)
		}
		return
	}
}

func Tokens[S stream.Stream[T], T stream.Token, P parser.ParseFunc[S, T, []T]](cmp func(lhs, rhs T) bool, tokens S) P {
	return func(input S) (actual []T, remaining S, err error) {
		n := len(tokens)
		if len(input) < n {
			err = io.ErrUnexpectedEOF
		} else {
			for i, tok := range tokens {
				if !cmp(input[i], tok) {
					err = Unexpected(tok, input[i])
					return
				}
			}

			actual, remaining = input[:n], input[n:]
		}
		return
	}
}
