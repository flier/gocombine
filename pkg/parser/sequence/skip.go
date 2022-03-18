package sequence

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Skip discards the value of the `p2` parser and returns the value of `p1`.
/// Fails if any of the parsers fails.
func Skip[
	S stream.Stream[T],
	T stream.Token,
	O1, O2 any,
](p1 parser.Func[S, T, O1], p2 parser.Func[S, T, O2]) parser.Func[S, T, O1] {
	return func(input S) (out O1, remaining S, err error) {
		out, remaining, err = p1.Parse(input)
		if err != nil {
			remaining = input
			return
		}

		_, remaining, err = p2.Parse(remaining)
		if err != nil {
			remaining = input
			return
		}

		return
	}
}
