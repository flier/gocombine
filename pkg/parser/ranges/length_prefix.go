package ranges

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
	"golang.org/x/exp/constraints"
)

// LengthPrefix takes a parser which parses a `length` then extracts a range of that length and returns it.
// Commonly used in binary formats.
func LengthPrefix[T stream.Token, N constraints.Integer](length parser.Func[T, N]) parser.Func[T, []T] {
	return func(input []T) (out []T, remaining []T, err error) {
		var n N

		if n, remaining, err = length(input); err != nil {
			return
		}

		out, remaining, err = stream.UnconsRange(remaining, int(n))

		return
	}
}
