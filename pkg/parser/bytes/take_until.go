package bytes

import (
	"bytes"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// TakeUntil reads a range of 0 or more tokens until `b` is found.
func TakeUntil(b ...byte) parser.Func[byte, []byte] {
	return func(input []byte) (out []byte, remaining []byte, err error) {
		if i := bytes.Index(input, b); i >= 0 {
			out, remaining, err = stream.UnconsRange(input, i)
		} else {
			err = parser.UnexpectedRange(b, input)
		}

		return
	}
}
