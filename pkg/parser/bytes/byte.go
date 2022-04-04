package bytes

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
)

// Any parses any byte.
func Any() parser.Func[byte, byte] {
	return parser.Expected(func(input []byte) (byte, []byte, error) {
		return stream.Uncons(input)
	}, "any")
}

// Byte parses a byte and succeeds if the byte is equal to `b`.
func Byte(b byte) parser.Func[byte, byte] {
	return token.Token(b).Expected("byte")
}

// Space parse a single whitespace.
func Space() parser.Func[byte, byte] {
	return token.Satisfy(func(b byte) bool {
		return unicode.IsSpace(rune(b))
	}).Expected("whitespace")
}

// Spaces parse zero or more spaces.
func Spaces() parser.Func[byte, []byte] {
	return repeat.Many(Space()).Expected("whitespaces")
}

// NewLine parses a newline character (`'\n'`).
func NewLine() parser.Func[byte, byte] {
	return token.Token[byte]('\n').Expected("newline")
}

// CrLf parses carriage return and newline (`"\r\n"`), returning the newline character.
func CrLf() parser.Func[byte, []byte] {
	return choice.And(
		token.Token[byte]('\r'),
		token.Token[byte]('\n'),
	).Expected("crlf")
}

// Tab parses a tab character (`'\t'`).
func Tab() parser.Func[byte, byte] {
	return token.Token[byte]('\t').Expected("tab")
}

// Upper parses an uppercase letter.
func Upper() parser.Func[byte, byte] {
	return token.Satisfy(func(b byte) bool {
		return unicode.IsUpper(rune(b))
	}).Expected("uppercase letter")
}

// Lower parses a lowercase letter.
func Lower() parser.Func[byte, byte] {
	return token.Satisfy(func(b byte) bool {
		return unicode.IsLower(rune(b))
	}).Expected("lowercase letter")
}

// Letter parses an alphabet letter.
func Letter() parser.Func[byte, byte] {
	return token.Satisfy(func(b byte) bool {
		return unicode.IsLetter(rune(b))
	}).Expected("letter")
}

// AlphaNum parses either an alphabet letter or digit.
func AlphaNum() parser.Func[byte, byte] {
	return choice.Or(
		Letter(),
		Digit(),
	).Expected("letter or digit")
}
