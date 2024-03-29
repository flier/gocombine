package bytes

import (
	"unicode"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/token"
)

// Cmp parses the bytes `s`, using `cmp` to compare each character.
func Cmp(s []byte, cmp func(l, r byte) bool) parser.Func[byte, []byte] {
	return token.Tokens(cmp, []byte(s), []byte(s)).Expected("cmp")
}

// Bytes parses the bytes `s`.
func Bytes(s []byte) parser.Func[byte, []byte] {
	return Cmp(s, func(l, r byte) bool { return l == r }).Expected("bytes")
}

// Fold parses the bytes `s`, are equal under Unicode case-folding.
func Fold(s []byte) parser.Func[byte, []byte] {
	return Cmp(s, func(l, r byte) bool {
		return unicode.ToLower(rune(l)) == unicode.ToLower(rune(r))
	}).Expected("bytes fold")
}
