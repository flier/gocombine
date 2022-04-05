package sequence

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Between parses `open` followed by `p` followed by `closing`.
/// Returns the value of `p`.
func Between[
	T stream.Token,
	O1, O2, O3 any,
](
	open parser.Func[T, O1],
	closing parser.Func[T, O2],
	p parser.Func[T, O3],
) parser.Func[T, O3] {
	return parser.Expected(func(input []T) (out O3, remaining []T, err error) {
		if _, remaining, err = open(input); err != nil {
			return
		}

		if out, remaining, err = p(remaining); err != nil {
			return
		}

		if _, remaining, err = closing(remaining); err != nil {
			return
		}

		return
	}, "between")
}
