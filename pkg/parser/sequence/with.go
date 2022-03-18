package sequence

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// With discards the value of the `p1` parser and returns the value of `p2`. Fails if any of the parsers fails.
func With[
	S stream.Stream[T],
	T stream.Token,
	O1, O2 any,
](p1 parser.Func[S, T, O1], p2 parser.Func[S, T, O2]) parser.Func[S, T, O2] {
	return func(input S) (parse O2, remaining S, err error) {
		_, remaining, err = p1.Parse(input)
		if err != nil {
			return
		}

		return p2.Parse(remaining)
	}
}
