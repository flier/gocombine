package ranges

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
	"golang.org/x/exp/constraints"
)

// LengthPrefix takes a parser which parses a `length` then extracts a range of that length and returns it.
// Commonly used in binary formats.
func LengthPrefix[
	S stream.Stream[T],
	T stream.Token,
	N constraints.Integer,
](
	length parser.Func[S, T, N],
) parser.Func[S, T, []T] {
	return func(input S) (out []T, remaining S, err error) {
		var n N

		if n, remaining, err = length(input); err != nil {
			return
		}

		out, remaining, err = stream.UnconsRange(remaining, int(n))

		return
	}
}
