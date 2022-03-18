package repeat

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Many parses `parser` zero or more times returning a collection with the values from `parser`.
func Many[
	S stream.Stream[T],
	T stream.Token,
	O any,
](parser parser.Func[S, T, O]) parser.Func[S, T, []O] {
	return func(input S) (parsed []O, remaining S, err error) {
		remaining = input

		for !stream.Empty(remaining) {
			var o O
			var rest S

			if o, rest, err = parser.Parse(remaining); err != nil {
				err = nil
				break
			}

			parsed = append(parsed, o)
			remaining = rest
		}

		return
	}
}

// Parses `parser` one or more times returning a collection with the values from `parser`.
func Many1[
	S stream.Stream[T],
	T stream.Token,
	O any,
](parser parser.Func[S, T, O]) parser.Func[S, T, []O] {
	return func(input S) (parsed []O, remaining S, err error) {
		remaining = input

		var o O
		if o, remaining, err = parser.Parse(remaining); err != nil {
			remaining = input
			return
		}

		parsed = []O{o}

		for !stream.Empty(remaining) {
			var o O
			var rest S

			if o, rest, err = parser.Parse(remaining); err != nil {
				err = nil
				break
			}

			parsed = append(parsed, o)
			remaining = rest
		}

		return
	}
}

// Parses `p` zero or more times ignoring the result.
func SkipMany[
	S stream.Stream[T],
	T stream.Token,
	O any,
](parser parser.Func[S, T, O]) parser.Func[S, T, any] {
	return combinator.Ignore(Many(parser))
}

// Parses `p` one or more times ignoring the result.
func SkipMany1[
	S stream.Stream[T],
	T stream.Token,
	O any,
](parser parser.Func[S, T, O]) parser.Func[S, T, any] {
	return combinator.Ignore(Many1(parser))
}
