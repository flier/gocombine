package repeat

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Many parses `parser` zero or more times returning a collection with the values from `parser`.
func Many[

	T stream.Token,
	O any,
](
	parser parser.Func[T, O],
) parser.Func[T, []O] {
	return func(input []T) (parsed []O, remaining []T, err error) {
		remaining = input

		for !stream.Empty(remaining) {
			var o O

			var rest []T

			if o, rest, err = parser(remaining); err != nil {
				err = nil

				break
			}

			parsed = append(parsed, o)
			remaining = rest
		}

		return
	}
}

// Many1 parses `parser` one or more times returning a collection with the values from `parser`.
func Many1[

	T stream.Token,
	O any,
](
	parser parser.Func[T, O],
) parser.Func[T, []O] {
	return func(input []T) (parsed []O, remaining []T, err error) {
		remaining = input

		var o O

		if o, remaining, err = parser(remaining); err != nil {
			remaining = input

			return
		}

		parsed = []O{o}

		for !stream.Empty(remaining) {
			var o O

			var rest []T

			if o, rest, err = parser(remaining); err != nil {
				err = nil

				break
			}

			parsed = append(parsed, o)
			remaining = rest
		}

		return
	}
}

// SkipMany parses `p` zero or more times ignoring the result.
func SkipMany[

	T stream.Token,
	O any,
](
	parser parser.Func[T, O],
) parser.Func[T, any] {
	return combinator.Ignore(Many(parser))
}

// SkipMany1 parses `p` one or more times ignoring the result.
func SkipMany1[

	T stream.Token,
	O any,
](
	parser parser.Func[T, O],
) parser.Func[T, any] {
	return combinator.Ignore(Many1(parser))
}
