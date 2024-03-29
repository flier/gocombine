package char

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/token"
)

// Cmp parses the string `s`, using `cmp` to compare each character.
func Cmp(s string, cmp func(l, r rune) bool) parser.Func[rune, []rune] {
	return token.Tokens(cmp, []rune(s), []rune(s)).Expected("cmp")
}

// String parses the string `s`.
func String(s string) parser.Func[rune, []rune] {
	return Cmp(s, func(l, r rune) bool { return l == r }).Expected("string")
}

// StringFold parses the string `s`, are equal under Unicode case-folding.
func Fold(s string) parser.Func[rune, []rune] {
	return Cmp(s, func(l, r rune) bool { return unicode.ToLower(l) == unicode.ToLower(r) }).Expected("char fold")
}
