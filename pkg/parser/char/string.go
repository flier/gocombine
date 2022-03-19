package char

import (
	"fmt"
	"unicode"
	"unicode/utf16"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
)

type IntoString interface {
	byte | rune | []byte | []rune | []uint16 | string
}

func ToString[S IntoString](s S) string {
	switch v := interface{}(s).(type) {
	case byte:
		return string([]byte{v})
	case rune:
		return string([]rune{v})
	case []byte:
		return string(v)
	case []rune:
		return string(v)
	case []uint16:
		return string(utf16.Decode(v))
	case string:
		return v
	default:
		panic(fmt.Errorf("unexpected %T", v))
	}
}

func AsString[
	S stream.Stream[T],
	T stream.Token,
	O IntoString,
](parser parser.Func[S, T, O]) parser.Func[S, T, string] {
	return combinator.Map(parser, ToString[O])
}

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
