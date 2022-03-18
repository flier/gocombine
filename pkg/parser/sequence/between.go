package sequence

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Between parses `open` followed by `parser` followed by `close`.
/// Returns the value of `parser`.
func Between[
	S stream.Stream[T],
	T stream.Token,
	O1, O2, O3 any,
](
	open parser.Func[S, T, O1],
	close parser.Func[S, T, O2],
	parser parser.Func[S, T, O3],
) parser.Func[S, T, O3] {
	return func(input S) (out O3, remaining S, err error) {
		if _, remaining, err = open.Parse(input); err != nil {
			remaining = input
			return
		}

		if out, remaining, err = parser.Parse(remaining); err != nil {
			remaining = input
			return
		}

		if _, remaining, err = close.Parse(remaining); err != nil {
			remaining = input
			return
		}

		return
	}
}
