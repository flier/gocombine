package char

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
)

// Char parses a character and succeeds if the character is equal to `c`.
func Char[S stream.Stream[rune]](c rune) parser.Func[S, rune, rune] {
	return token.Token[S](c)
}

// Space parse a single whitespace.
func Space[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Satisfy[S](unicode.IsSpace).Expected("whitespace")
}

// Spaces parse zero or more spaces.
func Spaces[S stream.Stream[rune]]() parser.Func[S, rune, []rune] {
	return repeat.Many(Space[S]()).Expected("whitespaces")
}

// NewLine parses a newline character (`'\n'`).
func NewLine[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Token[S]('\n').Expected("newline")
}

// CrLf parses carriage return and newline (`"\r\n"`), returning the newline character.
func CrLf[S stream.Stream[rune]]() parser.Func[S, rune, []rune] {
	return choice.And(
		token.Token[S]('\r'),
		token.Token[S]('\n'),
	).Expected("crlf")
}

// Tab parses a tab character (`'\t'`).
func Tab[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Token[S]('\t').Expected("tab")
}

// Upper parses an uppercase letter.
func Upper[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Satisfy[S](unicode.IsUpper).Expected("uppercase letter")
}

// Lower parses a lowercase letter.
func Lower[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Satisfy[S](unicode.IsLower).Expected("lowercase letter")
}

// Letter parses an alphabet letter.
func Letter[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return token.Satisfy[S](unicode.IsLetter).Expected("letter")
}

// AlphaNum parses either an alphabet letter or digit.
func AlphaNum[S stream.Stream[rune]]() parser.Func[S, rune, rune] {
	return choice.Or(
		Letter[S](),
		Digit[S](),
	).Expected("letter or digit")
}
