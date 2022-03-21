package sequence

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Skip discards the value of the `p2` parser and returns the value of `p1`.
/// Fails if any of the parsers fails.
func Skip[
	S stream.Stream[T],
	T stream.Token,
	O1, O2 any,
](p1 parser.Func[S, T, O1], p2 parser.Func[S, T, O2],
) parser.Func[S, T, O1] {
	return combinator.Attempt(func(input S) (out O1, remaining S, err error) {
		if out, remaining, err = p1.Parse(input); err != nil {
			return
		}

		if _, remaining, err = p2.Parse(remaining); err != nil {
			return
		}

		return
	})
}
