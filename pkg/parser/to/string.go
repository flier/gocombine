package to

import (
	"fmt"
	"unicode/utf16"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// StringLike constraints for the types that can be converted to string.
type StringLike interface {
	byte | rune | []byte | []rune | []uint16 | string
}

// Str convert a `StringLike` type to a string.
func Str[S StringLike](s S) string {
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
	O StringLike,
](parser parser.Func[S, T, O]) parser.Func[S, T, string] {
	return combinator.Map(parser, Str[O])
}

// StringSlice convert the result of `parser` to a string slice.
func StringSlice[
	S stream.Stream[T],
	T stream.Token,
	O StringLike,
](parser parser.Func[S, T, []O]) parser.Func[S, T, []string] {
	return combinator.Map(parser, func(s []O) (r []string) {
		r = make([]string, len(s))
		for i, v := range s {
			r[i] = Str(v)
		}
		return
	})
}
