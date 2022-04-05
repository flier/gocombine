package repeat

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Count parses `parser` from zero up to `count` times.
func Count[T stream.Token, O any](count int, parser parser.Func[T, O]) parser.Func[T, []O] {
	return CountMinMax(0, count, parser)
}

var errExpected = parser.ErrExpected

// CountMinMax parses `p` from `min` to `max` times (including `min` and `max`).
func CountMinMax[T stream.Token, O any](min, max int, p parser.Func[T, O]) parser.Func[T, []O] {
	return parser.Expected(func(input []T) (parsed []O, remaining []T, err error) {
		remaining = input
		parsed = make([]O, 0, max)

		for i := 0; i < max; i++ {
			var o O

			var rest []T
			if o, rest, err = p(remaining); err != nil {
				break
			}

			parsed = append(parsed, o)
			remaining = rest
		}

		if len(parsed) < min {
			err = fmt.Errorf("%d more elements, %w", min-len(parsed), errExpected)
		} else {
			err = nil
		}

		return
	}, "count min max")
}

// SkipCount parses `parser` from zero up to `count` times skipping the output of `parser`.
func SkipCount[T stream.Token, O any](count int, parser parser.Func[T, O]) parser.Func[T, any] {
	return combinator.Ignore(Count(count, parser)).Expected("skip count")
}

// SkipCountMinMax parses `parser` from `min` to `max` times (including `min` and `max`)
// skipping the output of `parser`.
func SkipCountMinMax[T stream.Token, O any](min, max int, parser parser.Func[T, O]) parser.Func[T, any] {
	return combinator.Ignore(CountMinMax(min, max, parser)).Expected("skip count min max")
}
