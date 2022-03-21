package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Attempt behaves as `parser` except it always acts as `parser` peeked instead of committed on its parse.
func Attempt[T stream.Token, O any](parser parser.Func[T, O]) parser.Func[T, O] {
	return func(input []T) (parsed O, remaining []T, err error) {
		parsed, remaining, err = parser(input)
		if err != nil {
			remaining = input
		}

		return
	}
}
