package regex

import (
	"fmt"
	"regexp"
	"unicode/utf16"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Match matches `re` on the input returning the entire input if it matches.
// Never consumes any input.
func Match[
	S stream.Stream[T],
	T stream.Token,
](
	re *regexp.Regexp,
) parser.Func[S, T, bool] {
	return func(input S) (matched bool, remaining S, err error) {
		switch v := interface{}(input).(type) {
		case []byte:
			matched = re.Match(v)

		case []rune:
			matched = re.MatchString(string(v))

		case []uint16:
			s := string(utf16.Decode(v))
			matched = re.MatchString(s)

		case string:
			matched = re.MatchString(v)

		default:
			err = fmt.Errorf("unsupported type, %T, %w", v, parser.ErrUnexpected)
		}

		remaining = input

		return
	}
}
