package fn

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Parser convert a `parser` to function.
func Parser[
	S stream.Stream[T],
	T stream.Token,
	O any](
	parser parser.Parser[S, T, O],
) parser.Func[S, T, O] {
	return func(input S) (out O, remaining S, err error) {
		return parser.Parse(input)
	}
}
