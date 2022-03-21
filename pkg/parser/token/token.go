package token

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Any parses any token.
func Any[S stream.Stream[T], T stream.Token]() parser.Func[S, T, T] {
	return func(input S) (T, S, error) {
		return stream.Uncons(input)
	}
}

// Token parses a character and succeeds if the character is equal to `tok`.
func Token[S stream.Stream[T], T stream.Token](tok T) parser.Func[S, T, T] {
	return combinator.Attempt(func(input S) (actual T, remaining S, err error) {
		if actual, remaining, err = stream.Uncons(input); err != nil {
		} else if actual != tok {
			err = parser.Unexpected([]T{tok}, []T{actual})
		}

		return
	})
}

// Tokens parses multiple tokens.
//
// Consumes items from the input and compares them to the values from `tokens` using the
/// comparison function `cmp`. Succeeds if all the items from `tokens` are matched in the input
/// stream and fails otherwise with `expected` used as part of the error.
func Tokens[
	S stream.Stream[T],
	T stream.Token,
](
	cmp func(lhs, rhs T) bool,
	expected, tokens []T,
) parser.Func[S, T, []T] {
	return combinator.Attempt(func(input S) (actual []T, remaining S, err error) {
		actual = make([]T, len(tokens))
		remaining = input

		for i, tok := range tokens {
			if actual[i], remaining, err = stream.Uncons(remaining); err != nil {
				actual = actual[:i]

				break
			}

			if !cmp(actual[i], tok) {
				actual = actual[:i+1]
				err = parser.Unexpected(expected, actual)

				break
			}
		}

		return
	})
}

// OneOf extract one token and succeeds if it is part of `tokens`.
func OneOf[S stream.Stream[T], T stream.Token](tokens S) parser.Func[S, T, T] {
	return combinator.Attempt(func(input S) (actual T, remaining S, err error) {
		if actual, remaining, err = stream.Uncons(input); err != nil {
			return
		}

		for _, tok := range tokens {
			if actual == tok {
				return
			}
		}

		err = parser.Unexpected(tokens, []T{actual})

		return
	})
}

// NoneOf extract one token and succeeds if it is not part of `tokens`.
func NoneOf[S stream.Stream[T], T stream.Token](tokens S) parser.Func[S, T, T] {
	return combinator.Attempt(func(input S) (actual T, remaining S, err error) {
		if actual, remaining, err = stream.Uncons(input); err != nil {
			return
		}

		for _, tok := range tokens {
			if actual == tok {
				err = parser.Unexpected(nil, []T{actual})

				return
			}
		}

		return
	})
}

// EOF succeeds only if the stream is at end of input, fails otherwise.
func EOF[S stream.Stream[T], T stream.Token]() parser.Func[S, T, bool] {
	return func(input S) (bool, S, error) {
		if stream.Empty(input) {
			return true, input, nil
		}

		return false, input, parser.Unexpected(nil, input)
	}
}
