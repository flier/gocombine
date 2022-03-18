package bytes

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
)

// Digit parses a base-10 digit.
func Digit[S stream.Stream[byte]]() parser.Func[S, byte, byte] {
	return token.Satisfy[S](func(b byte) bool {
		return '0' <= b && b <= '9'
	}).Expected("digit")
}

// OctDigit parses an octal digit.
func OctDigit[S stream.Stream[byte]]() parser.Func[S, byte, byte] {
	return token.Satisfy[S](func(b byte) bool {
		return b >= '0' && b <= '7'
	}).Expected("octal digit")
}

// HexDigit parses a hexdecimal digit with uppercase and lowercase.
func HexDigit[S stream.Stream[byte]]() parser.Func[S, byte, byte] {
	return token.Satisfy[S](func(b byte) bool {
		return unicode.IsDigit(rune(b)) || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
	}).Expected("octal digit")
}
