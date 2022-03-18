package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Attempt behaves as `parser` except it always acts as `parser` peeked instead of committed on its parse.
func Attempt[S stream.Stream[T], T stream.Token, O any](parser parser.Func[S, T, O]) parser.Func[S, T, O] {
	return func(input S) (parsed O, remaining S, err error) {
		parsed, remaining, err = parser.Parse(input)
		if err != nil {
			remaining = input
		}
		return
	}
}
