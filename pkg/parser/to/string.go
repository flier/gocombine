package to

import (
	"fmt"
	"unicode/utf16"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
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

// String convert the result of `parser` to a string.
func String[
	S stream.Stream[T],
	T stream.Token,
	O IntoString,
](parser parser.Func[S, T, O]) parser.Func[S, T, string] {
	return combinator.Map(parser, ToString[O])
}

// StringSlice convert the result of `parser` to a string slice.
func StringSlice[
	S stream.Stream[T],
	T stream.Token,
	O IntoString,
](parser parser.Func[S, T, []O]) parser.Func[S, T, []string] {
	return combinator.Map(parser, func(s []O) (r []string) {
		r = make([]string, len(s))
		for i, v := range s {
			r[i] = ToString(v)
		}
		return
	})
}
