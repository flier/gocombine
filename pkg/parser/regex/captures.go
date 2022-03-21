package regex

import (
	"fmt"
	"regexp"
	"unicode/utf16"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

const numOfLocs = 2

// Captures matches `re` on the input by running `Find[String]SubmatchIndex` on the input.
/// Returns the captures of the first match and consumes the input up until the end of that match.
func Captures[T stream.Token](re *regexp.Regexp) parser.Func[T, [][]T] {
	return func(input []T) (captured [][]T, remaining []T, err error) {
		var loc []int

		switch v := interface{}(input).(type) {
		case []byte:
			loc = re.FindSubmatchIndex(v)

		case []rune:
			loc = re.FindStringSubmatchIndex(string(v))

		case []uint16:
			s := string(utf16.Decode(v))
			loc = re.FindStringSubmatchIndex(s)

		default:
			err = fmt.Errorf("unsupported type, %T, %w", v, parser.ErrUnexpected)
		}

		if loc == nil {
			remaining = input
		} else {
			captured = make([][]T, len(loc)/numOfLocs)

			for i := 0; i < len(captured); i++ {
				captured[i], remaining, _ = stream.UnconsRange(input, loc[i*numOfLocs+1])
				_, captured[i], _ = stream.UnconsRange(captured[i], loc[i*numOfLocs])
			}
		}

		return
	}
}

// CapturesMany matches `re` on the input by running `FindAll[String]SubmatchIndex` on the input.
/// Returns all captures until the end of the last match.
func CapturesMany[T stream.Token](re *regexp.Regexp) parser.Func[T, [][][]T] {
	return func(input []T) (captured [][][]T, remaining []T, err error) {
		var locs [][]int

		switch v := interface{}(input).(type) {
		case []byte:
			locs = re.FindAllSubmatchIndex(v, -1)

		case []rune:
			locs = re.FindAllStringSubmatchIndex(string(v), -1)

		case []uint16:
			s := string(utf16.Decode(v))
			locs = re.FindAllStringSubmatchIndex(s, -1)

		default:
			err = fmt.Errorf("unsupported type, %T, %w", v, parser.ErrUnexpected)
		}

		if locs == nil {
			remaining = input
		} else {
			captured = make([][][]T, len(locs))

			for i, loc := range locs {
				captured[i] = make([][]T, len(loc)/numOfLocs)

				for j := 0; j < len(captured[i]); j++ {
					if loc[j*numOfLocs+1] >= 0 && loc[j*numOfLocs] >= 0 {
						captured[i][j], remaining, _ = stream.UnconsRange(input, loc[j*numOfLocs+1])
						_, captured[i][j], _ = stream.UnconsRange(captured[i][j], loc[j*numOfLocs])
					}
				}
			}
		}

		return
	}
}
