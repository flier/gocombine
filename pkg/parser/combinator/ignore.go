package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Ignore discards the value of the `parser`.
func Ignore[T stream.Token, I any](parser parser.Func[T, I]) parser.Func[T, any] {
	return func(input []T) (ignored any, remaining []T, err error) {
		_, remaining, err = parser(input)

		return
	}
}
