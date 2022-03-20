package ranges

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Recognize returns committed input range.
func Recognize[
	S stream.Stream[T],
	T stream.Token,
	O any,
](parser parser.Func[S, T, O]) parser.Func[S, T, []T] {
	return func(input S) (out []T, remaining S, err error) {
		_, remaining, err = parser(input)
		if err == nil {
			out, remaining, err = stream.UnconsRange(input, stream.Len(input)-stream.Len(remaining))
		}
		return
	}
}
