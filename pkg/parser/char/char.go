package char

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
)

// Char parses a character and succeeds if the character is equal to `c`.
func Char[S stream.Stream[rune]](c rune) parser.Func[S, rune, rune] {
	return token.Token[S](c)
}

// Digit parses a base-10 digit.
func Digit[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Satisfy[S](unicode.IsDigit).Expected("digit")
}

// Space parse a single whitespace
func Space[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Satisfy[S](unicode.IsSpace).Expected("whitespace")
}

// Spaces parse zero or more spaces
func Spaces[S stream.Stream[rune]]() parser.Func[S, rune, []rune] {
	return repeat.Many(Space[S]()).Expected("whitespaces")
}

// NewLine parses a newline character (`'\n'`).
func NewLine[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Token[S]('\n').Expected("newline")
}

// CrLf parses carriage return and newline (`"\r\n"`), returning the newline character.
func CrLf[S stream.Stream[rune]]() parser.Func[S, rune, []rune] {
	return combinator.And[S, rune, rune](
		token.Token[S]('\r'),
		token.Token[S]('\n'),
	).Expected("crlf")
}

// Tab parses a tab character (`'\t'`).
func Tab[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Token[S]('\t').Expected("tab")
}

// Upper parses an uppercase letter
func Upper[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Satisfy[S](unicode.IsUpper).Expected("uppercase letter")
}

// Upper parses a lowercase letter
func Lower[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Satisfy[S](unicode.IsLower).Expected("lowercase letter")
}

// Letter parses an alphabet letter
func Letter[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Satisfy[S](unicode.IsLetter).Expected("letter")
}

// AlphaNum parses either an alphabet letter or digit
func AlphaNum[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return combinator.Or[S, rune, rune](
		Letter[S](),
		Digit[S](),
	).Expected("letter or digit")
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
