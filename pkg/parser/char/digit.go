package char

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/token"
)

// Digit parses a base-10 digit.
func Digit() parser.Func[rune, rune] {
	return token.Satisfy(unicode.IsDigit).Expected("digit")
}

// OctDigit parses an octal digit.
func OctDigit() parser.Func[rune, rune] {
	return token.Satisfy(func(c rune) bool {
		return c >= '0' && c <= '7'
	}).Expected("octal digit")
}

// HexDigit parses a hexdecimal digit with uppercase and lowercase.
func HexDigit() parser.Func[rune, rune] {
	return token.Satisfy(func(c rune) bool {
		return unicode.IsDigit(c) || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
	}).Expected("hex digit")
}
