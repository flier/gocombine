package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// AndThen parses with `parser` and applies `f` on the result if `parser` parses successfully.
// `f` may optionally fail with an error.
func AndThen[T stream.Token, I, O any](parser parser.Func[T, I], f func(I) (O, error)) parser.Func[T, O] {
	return Attempt(func(input []T) (parsed O, remaining []T, err error) {
		var i I
		if i, remaining, err = parser(input); err != nil {
			return
		}

		parsed, err = f(i)

		return
	})
}
