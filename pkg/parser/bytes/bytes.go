package bytes

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
)

// Cmp parses the bytes `s`, using `cmp` to compare each character.
func Cmp[S stream.Stream[byte]](s []byte, cmp func(l, r byte) bool) parser.Func[S, byte, []byte] {
	p := token.Tokens[S](cmp, []byte(s), []byte(s))

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
	return Cmp[S](s, func(l, r byte) bool { return l == r })
}

// Fold parses the bytes `s`, are equal under Unicode case-folding.
func Fold[S stream.Stream[byte]](s []byte) parser.Func[S, byte, []byte] {
	return Cmp[S](s, func(l, r byte) bool { return unicode.ToLower(rune(l)) == unicode.ToLower(rune(r)) })
}
