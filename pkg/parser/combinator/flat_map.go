package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// FlatMap uses `f` to map over the output of `parser`. If `f` returns an error the parser fails.
func FlatMap[

	T stream.Token,
	O, P any,
](
	parser parser.Func[T, O],
	f func(O) (P, error),
) parser.Func[T, P] {
	return func(input []T) (out P, remaining []T, err error) {
		var o O

		if o, remaining, err = parser(input); err != nil {
			return
		}

		out, err = f(o)

		return
	}
}
