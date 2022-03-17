package token

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Satisfy parses a token and succeeds depending on the result of `predicate`.
func Satisfy[S stream.Stream[T], T stream.Token](predicate func(T) bool) parser.Func[S, T, T] {
	return func(input S) (actual T, remaining S, err error) {
		if actual, remaining, err = stream.Uncons(input); err != nil {
			remaining = input
		} else if !predicate(actual) {
			remaining, err = input, fmt.Errorf("satisfy, %w", parser.ErrExpected)
		}

		return
	}
}

// SatisfyMap parses a token and passes it to `predicate`. If `predicate` succeeds and returns the value.
// If `predicate` returns error the parser fails without consuming any input.
func SatisfyMap[S stream.Stream[T], T stream.Token, O any](predicate func(T) (O, error)) parser.Func[S, T, O] {
	return func(input S) (out O, remaining S, err error) {
		var tok T
		if tok, remaining, err = stream.Uncons(input); err != nil {
			remaining = input
		} else if out, err = predicate(tok); err != nil {
			remaining, err = input, fmt.Errorf("satisfy and map, %w", parser.ErrExpected)
		}

		return
	}
}
