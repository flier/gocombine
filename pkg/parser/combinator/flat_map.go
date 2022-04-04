package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// FlatMap uses `f` to map over the output of `parser`. If `f` returns an error the parser fails.
func FlatMap[T stream.Token, O any](parser parser.Func[T, []T], f func([]T) (O, error)) parser.Func[T, O] {
	return Attempt(func(input []T) (out O, remaining []T, err error) {
		var parsed []T

		if parsed, remaining, err = parser(input); err != nil {
			return
		}

		out, err = f(parsed)

		return
	}).Expected("flat map")
}
