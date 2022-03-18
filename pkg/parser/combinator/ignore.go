package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Ignore discards the value of the `parser`.
func Ignore[S stream.Stream[T], T stream.Token, I any](parser parser.Func[S, T, I]) parser.Func[S, T, any] {
	return func(input S) (ignored any, remaining S, err error) {
		_, remaining, err = parser.Parse(input)
		return
	}
}
