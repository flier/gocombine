package bytes

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
)

// BytesCmp parses the bytes `s`, using `cmp` to compare each character.
func BytesCmp[S stream.Stream[byte]](s []byte, cmp func(l, r byte) bool) parser.Func[S, byte, []byte] {
	p := token.Tokens(cmp, []byte(s), []byte(s))

	return func(input S) (out []byte, remaining S, err error) {
		var bytes []byte
		if bytes, remaining, err = p(input); err != nil {
			remaining = input
		} else {
			out = bytes
		}
		return
	}
}

// Bytes parses the bytes `s`.
func Bytes[S stream.Stream[byte]](s []byte) parser.Func[S, byte, []byte] {
	return BytesCmp[S](s, func(l, r byte) bool { return l == r })
}

// Bytes parses the bytes `s`, are equal under Unicode case-folding.
func BytesFold[S stream.Stream[byte]](s []byte) parser.Func[S, byte, []byte] {
	return BytesCmp[S](s, func(l, r byte) bool { return unicode.ToLower(rune(l)) == unicode.ToLower(rune(r)) })
}
