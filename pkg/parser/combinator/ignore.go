package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Ignore discards the value of the `p`.
func Ignore[T stream.Token, I any](p parser.Func[T, I]) parser.Func[T, any] {
	return parser.Expected(func(input []T) (ignored any, remaining []T, err error) {
		_, remaining, err = p(input)

		return
	}, "ignore")
}
