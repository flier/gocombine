package char

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/token"
)

// StringCmp parses the string `s`, using `cmp` to compare each character.
func StringCmp(s string, cmp func(l, r rune) bool) parser.Func[rune, []rune] {
	p := token.Tokens(cmp, []rune(s), []rune(s))

	return func(input []rune) (out []rune, remaining []rune, err error) {
		var chars []rune

		if chars, remaining, err = p(input); err != nil {
			remaining = input
		} else {
			out = chars
		}

		return
	}
}

// String parses the string `s`.
func String(s string) parser.Func[rune, []rune] {
	return StringCmp(s, func(l, r rune) bool { return l == r })
}

// StringFold parses the string `s`, are equal under Unicode case-folding.
func StringFold(s string) parser.Func[rune, []rune] {
	return StringCmp(s, func(l, r rune) bool { return unicode.ToLower(l) == unicode.ToLower(r) })
}
