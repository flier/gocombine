package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Map uses `f` to map over the parsed value.
func Map[
	S stream.Stream[T],
	T stream.Token,
	I, O any,
](
	parser parser.Func[S, T, I], f func(I) O,
) parser.Func[S, T, O] {
	return func(input S) (parsed O, remaining S, err error) {
		var i I

		if i, remaining, err = parser(input); err != nil {
			return
		}

		parsed = f(i)

		return
	}
}
