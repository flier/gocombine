package regex

import (
	"fmt"
	"regexp"
	"unicode/utf16"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Find matches `re` on the input by running `find` on the input and returns the first match.
// Consumes all input up until the end of the first match.
func Find[T stream.Token](re *regexp.Regexp) parser.Func[T, []T] {
	return parser.Expected(func(input []T) (matched, remaining []T, err error) {
		var loc []int

		switch v := interface{}(input).(type) {
		case []byte:
			loc = re.FindIndex(v)

		case []rune:
			loc = re.FindStringIndex(string(v))

		case []uint16:
			s := string(utf16.Decode(v))
			loc = re.FindStringIndex(s)

		default:
			err = fmt.Errorf("unsupported type, %T, %w", v, parser.ErrUnexpected)
		}

		if loc == nil {
			remaining = input
		} else {
			matched, remaining, _ = stream.UnconsRange(input, loc[1])
			_, matched, _ = stream.UnconsRange(matched, loc[0])
		}

		return
	}, "find")
}

// FindMany matches `re` on the input by running `FindAll` on the input.
/// Returns all matches until the end of the last match.
func FindMany[T stream.Token](re *regexp.Regexp) parser.Func[T, [][]T] {
	return parser.Expected(func(input []T) (matched [][]T, remaining []T, err error) {
		var locs [][]int

		switch v := interface{}(input).(type) {
		case []byte:
			locs = re.FindAllIndex(v, -1)

		case []rune:
			locs = re.FindAllStringIndex(string(v), -1)

		case []uint16:
			s := string(utf16.Decode(v))
			locs = re.FindAllStringIndex(s, -1)

		default:
			err = fmt.Errorf("unsupported type, %T, %w", v, parser.ErrUnexpected)
		}

		if locs == nil {
			remaining = input
		} else {
			matched = make([][]T, len(locs))

			for i, loc := range locs {
				matched[i], remaining, _ = stream.UnconsRange(input, loc[1])
				_, matched[i], _ = stream.UnconsRange(matched[i], loc[0])
			}
		}

		return
	}, "find many")
}
