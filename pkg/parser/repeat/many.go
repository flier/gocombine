package repeat

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Many parses `parser` zero or more times returning a collection with the values from `parser`.
func Many[
	S stream.Stream[T],
	T stream.Token,
	O any,
	P parser.Func[S, T, O],
](parser P) parser.Func[S, T, []O] {
	return func(input S) (parsed []O, remaining S, err error) {
		remaining = input

		for len(remaining) > 0 {
			var o O
			if o, remaining, err = parser(remaining); err != nil {
				err = nil
				break
			}
			parsed = append(parsed, o)
		}

		return
	}
}

// Parses `parser` one or more times returning a collection with the values from `parser`.
func Many1[
	S stream.Stream[T],
	T stream.Token,
	O any,
	P parser.Func[S, T, O],
](parser P) parser.Func[S, T, []O] {
	return func(input S) (parsed []O, remaining S, err error) {
		remaining = input

		var o O
		if o, remaining, err = parser(remaining); err != nil {
			err = nil
		} else {
			parsed = append(parsed, o)
		}

		return
	}
}
