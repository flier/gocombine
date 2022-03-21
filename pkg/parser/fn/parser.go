package fn

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Parser convert a `parser` to function.
func Parser[T stream.Token, O any](parser parser.Parser[T, O]) parser.Func[T, O] {
	return func(input []T) (out O, remaining []T, err error) {
		return parser.Parse(input)
	}
}
