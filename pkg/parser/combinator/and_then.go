package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// AndThen parses with `p` and applies `f` on the result if `p` parses successfully.
// `f` may optionally fail with an error.
func AndThen[T stream.Token, I, O any](p parser.Func[T, I], f func(I) (O, error)) parser.Func[T, O] {
	return parser.Expected(func(input []T) (parsed O, remaining []T, err error) {
		var i I
		if i, remaining, err = p(input); err != nil {
			return
		}

		parsed, err = f(i)

		return
	}, "and then")
}
