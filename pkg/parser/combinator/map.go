package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Map uses `f` to map over the `p` parsed value.
func Map[T stream.Token, I, O any](p parser.Func[T, I], f func(I) O) parser.Func[T, O] {
	return parser.Expected(func(input []T) (parsed O, remaining []T, err error) {
		var i I

		if i, remaining, err = p(input); err != nil {
			return
		}

		parsed = f(i)

		return
	}, "map")
}
