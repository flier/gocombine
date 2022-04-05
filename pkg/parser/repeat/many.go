package repeat

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Many parses `p` zero or more times returning a collection with the values from `parser`.
func Many[T stream.Token, O any](p parser.Func[T, O]) parser.Func[T, []O] {
	return parser.Expected(func(input []T) (parsed []O, remaining []T, err error) {
		remaining = input

		for !stream.Empty(remaining) {
			var o O

			var rest []T

			if o, rest, err = p(remaining); err != nil {
				err = nil

				break
			}

			parsed = append(parsed, o)
			remaining = rest
		}

		return
	}, "many")
}

// Many1 parses `p` one or more times returning a collection with the values from `parser`.
func Many1[T stream.Token, O any](p parser.Func[T, O]) parser.Func[T, []O] {
	return parser.Expected(func(input []T) (parsed []O, remaining []T, err error) {
		remaining = input

		var o O

		if o, remaining, err = p(remaining); err != nil {
			return
		}

		parsed = []O{o}

		for !stream.Empty(remaining) {
			var o O

			var rest []T

			if o, rest, err = p(remaining); err != nil {
				err = nil

				break
			}

			parsed = append(parsed, o)
			remaining = rest
		}

		return
	}, "many1")
}

// SkipMany parses `p` zero or more times ignoring the result.
func SkipMany[T stream.Token, O any](parser parser.Func[T, O]) parser.Func[T, any] {
	return combinator.Ignore(Many(parser)).Expected("skip")
}

// SkipMany1 parses `p` one or more times ignoring the result.
func SkipMany1[T stream.Token, O any](parser parser.Func[T, O]) parser.Func[T, any] {
	return combinator.Ignore(Many1(parser)).Expected("skip")
}
