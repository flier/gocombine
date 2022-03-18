package sequence

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Then parses using `parser` and then passes the value to `f`
// which returns a parser used to parse the rest of the input.
func Then[
	S stream.Stream[T],
	T stream.Token,
	I, O any,
](parser parser.Func[S, T, I], f func(I) parser.Func[S, T, O]) parser.Func[S, T, O] {
	return combinator.Attempt(func(input S) (out O, remaining S, err error) {
		var in I
		in, remaining, err = parser(input)
		if err != nil {
			return
		}

		out, remaining, err = f(in).Parse(remaining)
		return
	})
}
