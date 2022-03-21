package ranges

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Recognize returns committed input range.
func Recognize[

	T stream.Token,
	O any,
](
	parser parser.Func[T, O],
) parser.Func[T, []T] {
	return func(input []T) (out []T, remaining []T, err error) {
		_, remaining, err = parser(input)
		if err == nil {
			out, remaining, err = stream.UnconsRange(input, stream.Len(input)-stream.Len(remaining))
		}

		return
	}
}
