package token

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Satisfy parses a token and succeeds depending on the result of `predicate`.
func Satisfy[T stream.Token](predicate func(T) bool) parser.Func[T, T] {
	return combinator.Attempt(func(input []T) (actual T, remaining []T, err error) {
		if actual, remaining, err = stream.Uncons(input); err != nil {
			// pass
		} else if !predicate(actual) {
			err = parser.UnexpectedToken(actual)
		}

		return
	}).Expected("satisfy")
}

// SatisfyMap parses a token and passes it to `predicate`. If `predicate` succeeds and returns the value.
// If `predicate` returns error the parser fails without consuming any input.
func SatisfyMap[T stream.Token, O any](predicate func(T) (O, error)) parser.Func[T, O] {
	return combinator.Attempt(func(input []T) (out O, remaining []T, err error) {
		var actual T
		if actual, remaining, err = stream.Uncons(input); err != nil {
			// pass
		} else if out, err = predicate(actual); err != nil {
			err = parser.UnexpectedToken(actual)
		}

		return
	}).Expected("satisfy map")
}
