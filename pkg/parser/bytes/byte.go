package bytes

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
)

// Byte parses a byte and succeeds if the byte is equal to `b`.
func Byte[S stream.Stream[byte]](b byte) parser.Func[S, byte, byte] {
	return token.Token[S](b)
}

// Space parse a single whitespace
func Space[S stream.Stream[byte]]() parser.Func[S, byte, byte] {
	return token.Satisfy[S](func(b byte) bool {
		return unicode.IsSpace(rune(b))
	}).Expected("whitespace")
}

// Spaces parse zero or more spaces
func Spaces[S stream.Stream[byte]]() parser.Func[S, byte, []byte] {
	return repeat.Many(Space[S]()).Expected("whitespaces")
}

// NewLine parses a newline character (`'\n'`).
func NewLine[S stream.Stream[byte]]() parser.Func[S, byte, byte] {
	return token.Token[S]('\n').Expected("newline")
}

// CrLf parses carriage return and newline (`"\r\n"`), returning the newline character.
func CrLf[S stream.Stream[byte]]() parser.Func[S, byte, []byte] {
	return choice.And(
		token.Token[S]('\r'),
		token.Token[S]('\n'),
	).Expected("crlf")
}

// Tab parses a tab character (`'\t'`).
func Tab[S stream.Stream[byte]]() parser.Func[S, byte, byte] {
	return token.Token[S]('\t').Expected("tab")
}

// Upper parses an uppercase letter
func Upper[S stream.Stream[byte]]() parser.Func[S, byte, byte] {
	return token.Satisfy[S](func(b byte) bool {
		return unicode.IsUpper(rune(b))
	}).Expected("uppercase letter")
}

// Lower parses a lowercase letter
func Lower[S stream.Stream[byte]]() parser.Func[S, byte, byte] {
	return token.Satisfy[S](func(b byte) bool {
		return unicode.IsLower(rune(b))
	}).Expected("lowercase letter")
}

// Letter parses an alphabet letter
func Letter[S stream.Stream[byte]]() parser.Func[S, byte, byte] {
	return token.Satisfy[S](func(b byte) bool {
		return unicode.IsLetter(rune(b))
	}).Expected("letter")
}

// AlphaNum parses either an alphabet letter or digit
func AlphaNum[S stream.Stream[byte]]() parser.Func[S, byte, byte] {
	return choice.Or(
		Letter[S](),
		Digit[S](),
	).Expected("letter or digit")
}
