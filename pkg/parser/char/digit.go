package char

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
)

// Digit parses a base-10 digit.
func Digit[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Satisfy[S](unicode.IsDigit).Expected("digit")
}

// OctDigit parses an octal digit.
func OctDigit[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Satisfy[S](func(c rune) bool {
		return c >= '0' && c <= '7'
	}).Expected("octal digit")
}

// HexDigit parses a hexdecimal digit with uppercase and lowercase.
func HexDigit[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Satisfy[S](func(c rune) bool {
		return unicode.IsDigit(c) || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
	}).Expected("octal digit")
}
