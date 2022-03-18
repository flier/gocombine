package char

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
)

// StringCmp parses the string `s`, using `cmp` to compare each character.
func StringCmp[S stream.Stream[rune]](s string, cmp func(l, r rune) bool) parser.Func[S, rune, string] {
	p := token.Tokens(cmp, []rune(s), []rune(s))

	return func(input S) (out string, remaining S, err error) {
		var chars []rune
		if chars, remaining, err = p(input); err != nil {
			remaining = input
		} else {
			out = string(chars)
		}
		return
	}
}

// String parses the string `s`.
func String[S stream.Stream[rune]](s string) parser.Func[S, rune, string] {
	return StringCmp[S](s, func(l, r rune) bool { return l == r })
}

// String parses the string `s`, are equal under Unicode case-folding.
func StringFold[S stream.Stream[rune]](s string) parser.Func[S, rune, string] {
	return StringCmp[S](s, func(l, r rune) bool { return unicode.ToLower(l) == unicode.ToLower(r) })
}