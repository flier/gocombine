package sequence

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Then parses using `p` and then passes the value to `f`
// which returns a parser used to parse the rest of the input.
func Then[T stream.Token, I, O any](p parser.Func[T, I], f func(I) parser.Func[T, O]) parser.Func[T, O] {
	return parser.Expected(func(input []T) (out O, remaining []T, err error) {
		var in I

		if in, remaining, err = p(input); err != nil {
			return
		}

		p := f(in)

		out, remaining, err = p(remaining)

		return
	}, "then")
}
