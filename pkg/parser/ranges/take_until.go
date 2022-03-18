package ranges

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// TakeUntil reads a range of 0 or more tokens until `b` is found.
func TakeUntil[S stream.Stream[T], T stream.Token](b ...T) parser.Func[S, T, []T] {
	return func(input S) (out []T, remaining S, err error) {
		if i := stream.Index(input, b); i >= 0 {
			out, remaining = input[:i], input[i:]
		} else {
			err = parser.Unexpected(b, input)
		}

		return
	}
}
