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
](parser parser.Func[S, T, O]) parser.Func[S, T, S] {
	return func(input S) (out S, remaining S, err error) {
		_, remaining, err = parser(input)
		if err == nil {
			out = input[:stream.Len(input)-stream.Len(remaining)]
		}
		return
	}
}
