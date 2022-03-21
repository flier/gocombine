package bytes

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/token"
)

// Digit parses a base-10 digit.
func Digit() parser.Func[byte, byte] {
	return token.Satisfy(func(b byte) bool {
		return '0' <= b && b <= '9'
	}).Expected("digit")
}

// OctDigit parses an octal digit.
func OctDigit() parser.Func[byte, byte] {
	return token.Satisfy(func(b byte) bool {
		return b >= '0' && b <= '7'
	}).Expected("octal digit")
}

// HexDigit parses a hexdecimal digit with uppercase and lowercase.
func HexDigit() parser.Func[byte, byte] {
	return token.Satisfy(func(b byte) bool {
		return unicode.IsDigit(rune(b)) || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
	}).Expected("octal digit")
}
