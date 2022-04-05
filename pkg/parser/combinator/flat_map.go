package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// FlatMap uses `f` to map over the output of `p`. If `f` returns an error the parser fails.
func FlatMap[T stream.Token, O any](p parser.Func[T, []T], f func([]T) (O, error)) parser.Func[T, O] {
	return parser.Expected(func(input []T) (out O, remaining []T, err error) {
		var parsed []T

		if parsed, remaining, err = p(input); err != nil {
			return
		}

		out, err = f(parsed)

		return
	}, "flat map")
}
