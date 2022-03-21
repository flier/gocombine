package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// AndThen parses with `parser` and applies `f` on the result if `parser` parses successfully.
// `f` may optionally fail with an error.
func AndThen[
	S stream.Stream[T],
	T stream.Token,
	I, O any,
](
	parser parser.Func[S, T, I], f func(I) (O, error),
) parser.Func[S, T, O] {
	return Attempt(func(input S) (parsed O, remaining S, err error) {
		var i I
		if i, remaining, err = parser(input); err != nil {
			return
		}

		parsed, err = f(i)

		return
	})
}
