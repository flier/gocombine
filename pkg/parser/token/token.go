package token

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
	"golang.org/x/exp/slices"
)

// Any parses any token.
func Any[T stream.Token]() parser.Func[T, T] {
	return func(input []T) (T, []T, error) {
		return stream.Uncons(input)
	}
}

// Token parses a character and succeeds if the character is equal to `tok`.
func Token[T stream.Token](tok T) parser.Func[T, T] {
	return func(input []T) (actual T, remaining []T, err error) {
		if actual, remaining, err = stream.Uncons(input); err != nil {
			// pass
		} else if actual != tok {
			err = parser.UnexpectedToken(actual, tok)
		}

		return
	}
}

// Tokens parses multiple tokens.
//
// Consumes items from the input and compares them to the values from `tokens` using the
/// comparison function `cmp`. Succeeds if all the items from `tokens` are matched in the input
/// stream and fails otherwise with `expected` used as part of the error.
func Tokens[T stream.Token](cmp func(lhs, rhs T) bool, expected, tokens []T) parser.Func[T, []T] {
	return func(input []T) (actual []T, remaining []T, err error) {
		actual = make([]T, len(tokens))
		remaining = input

		for i, tok := range tokens {
			if actual[i], remaining, err = stream.Uncons(remaining); err != nil {
				actual = actual[:i]

				break
			}

			if !cmp(actual[i], tok) {
				actual = actual[:i+1]
				err = parser.UnexpectedRange(expected, actual)

				break
			}
		}

		return
	}
}

// OneOf extract one token and succeeds if it is part of `tokens`.
func OneOf[T stream.Token](tokens []T) parser.Func[T, T] {
	return Satisfy(func(t T) bool { return slices.Contains(tokens, t) }).
		Expected(fmt.Sprintf("one of %s", parser.FormatRange(tokens)))
}

// NoneOf extract one token and succeeds if it is not part of `tokens`.
func NoneOf[T stream.Token](tokens []T) parser.Func[T, T] {
	return Satisfy(func(t T) bool { return !slices.Contains(tokens, t) }).
		Expected(fmt.Sprintf("none of %s", parser.FormatRange(tokens)))
}

// EOF succeeds only if the stream is at end of input, fails otherwise.
func EOF[T stream.Token]() parser.Func[T, bool] {
	return parser.Expected(func(input []T) (bool, []T, error) {
		if stream.Empty(input) {
			return true, input, nil
		}

		return false, input, parser.UnexpectedRange(nil, input)
	}, "eof")
}
