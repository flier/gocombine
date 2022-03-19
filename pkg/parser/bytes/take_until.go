package bytes

import (
	"bytes"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// TakeUntil reads a range of 0 or more tokens until `b` is found.
func TakeUntil[S stream.Stream[byte]](b ...byte) parser.Func[S, byte, []byte] {
	return func(input S) (out []byte, remaining S, err error) {
		if i := bytes.Index(input, b); i >= 0 {
			out, remaining, err = stream.UnconsRange(input, i)
		} else {
			err = parser.Unexpected(b, input)
		}

		return
	}
}
