package char

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
)

// Any parses any char.
func Any() parser.Func[rune, rune] {
	return parser.Expected(func(input []rune) (rune, []rune, error) {
		return stream.Uncons(input)
	}, "any")
}

// Char parses a character and succeeds if the character is equal to `c`.
func Char(c rune) parser.Func[rune, rune] {
	return token.Token(c).Expected("char")
}

// Space parse a single whitespace.
func Space() parser.Func[rune, rune] {
	return token.Satisfy(unicode.IsSpace).Expected("whitespace")
}

// Spaces parse zero or more spaces.
func Spaces() parser.Func[rune, []rune] {
	return repeat.Many(Space()).Expected("whitespaces")
}

// NewLine parses a newline character (`'\n'`).
func NewLine() parser.Func[rune, rune] {
	return token.Token('\n').Expected("newline")
}

// CrLf parses carriage return and newline (`"\r\n"`), returning the newline character.
func CrLf() parser.Func[rune, []rune] {
	return choice.And(
		token.Token('\r'),
		token.Token('\n'),
	).Expected("crlf")
}

// Tab parses a tab character (`'\t'`).
func Tab() parser.Func[rune, rune] {
	return token.Token('\t').Expected("tab")
}

// Upper parses an uppercase letter.
func Upper() parser.Func[rune, rune] {
	return token.Satisfy(unicode.IsUpper).Expected("uppercase letter")
}

// Lower parses a lowercase letter.
func Lower() parser.Func[rune, rune] {
	return token.Satisfy(unicode.IsLower).Expected("lowercase letter")
}

// Letter parses an alphabet letter.
func Letter() parser.Func[rune, rune] {
	return token.Satisfy(unicode.IsLetter).Expected("letter")
}

// AlphaNum parses either an alphabet letter or digit.
func AlphaNum() parser.Func[rune, rune] {
	return choice.Or(
		Letter(),
		Digit(),
	).Expected("letter or digit")
}

// OneOf extract one token and succeeds if it is part of `tokens`.
func OneOf(tokens string) parser.Func[rune, rune] {
	return token.OneOf([]rune(tokens)).Expected("one of")
}

// NoneOf extract one token and succeeds if it is not part of `tokens`.
func NoneOf(tokens string) parser.Func[rune, rune] {
	return token.NoneOf([]rune(tokens)).Expected("none of")
}
