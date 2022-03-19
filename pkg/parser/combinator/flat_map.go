package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// FlatMap uses `f` to map over the output of `self`. If `f` returns an error the parser fails.
func FlatMap[
	S stream.Stream[T],
	T stream.Token,
	O, P any,
](
	parser parser.Func[S, T, O],
	f func(O) (P, error),
) parser.Func[S, T, P] {
	return func(input S) (out P, remaining S, err error) {
		var o O
		o, remaining, err = parser(input)
		if err != nil {
			return
		}

		out, err = f(o)
		return
	}
}
