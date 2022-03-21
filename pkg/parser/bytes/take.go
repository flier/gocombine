package bytes

import (
	"unsafe"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Take reads a bytes of length `n`.
func Take(n int) parser.Func[byte, []byte] {
	return func(input []byte) (out []byte, remaining []byte, err error) {
		return stream.UnconsRange(input, n)
	}
}

// Take reads a bytes of length `T`.
func TakeOf[T any]() parser.Func[byte, []byte] {
	return func(input []byte) (out []byte, remaining []byte, err error) {
		var t T

		return stream.UnconsRange(input, int(unsafe.Sizeof(t)))
	}
}
