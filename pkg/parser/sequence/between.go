package sequence

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Between parses `open` followed by `parser` followed by `closing`.
/// Returns the value of `parser`.
func Between[
	T stream.Token,
	O1, O2, O3 any,
](
	open parser.Func[T, O1],
	closing parser.Func[T, O2],
	parser parser.Func[T, O3],
) parser.Func[T, O3] {
	return combinator.Attempt(func(input []T) (out O3, remaining []T, err error) {
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
