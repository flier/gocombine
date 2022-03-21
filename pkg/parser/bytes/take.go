package bytes

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Take reads a bytes of length `n`.
func Take(n int) parser.Func[byte, []byte] {
	return func(input []byte) (out []byte, remaining []byte, err error) {
		return stream.UnconsRange(input, n)
	}
}
