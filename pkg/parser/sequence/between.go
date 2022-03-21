package sequence

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Between parses `open` followed by `parser` followed by `closing`.
/// Returns the value of `parser`.
func Between[
	S stream.Stream[T],
	T stream.Token,
	O1, O2, O3 any,
](
	open parser.Func[S, T, O1],
	closing parser.Func[S, T, O2],
	parser parser.Func[S, T, O3],
) parser.Func[S, T, O3] {
	return combinator.Attempt(func(input S) (out O3, remaining S, err error) {
		if _, remaining, err = open.Parse(input); err != nil {
			return
		}

		if out, remaining, err = parser(remaining); err != nil {
			return
		}

		if _, remaining, err = closing.Parse(remaining); err != nil {
			return
		}

		return
	})
}
