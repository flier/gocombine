package sequence

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Skip discards the value of the `p2` parser and returns the value of `p1`.
/// Fails if any of the parsers fails.
func Skip[T stream.Token, O1, O2 any](p1 parser.Func[T, O1], p2 parser.Func[T, O2]) parser.Func[T, O1] {
	return parser.Expected(func(input []T) (out O1, remaining []T, err error) {
		if out, remaining, err = p1(input); err != nil {
			return
		}

		if _, remaining, err = p2(remaining); err != nil {
			return
		}

		return
	}, "skip")
}
