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

// CountMinMax parses `parser` from `min` to `max` times (including `min` and `max`).
func CountMinMax[T stream.Token, O any](min, max int, parser parser.Func[T, O]) parser.Func[T, []O] {
	return func(input []T) (parsed []O, remaining []T, err error) {
		remaining = input
		parsed = make([]O, 0, max)

		for i := 0; i < max; i++ {
			var o O

			if o, remaining, err = parser(remaining); err != nil {
				break
			}

			parsed = append(parsed, o)
		}

		if len(parsed) < min {
			err = fmt.Errorf("%d more elements, %w", min-len(parsed), errExpected)
		} else {
			err = nil
		}

		return
	}
}

// SkipCount parses `parser` from zero up to `count` times skipping the output of `parser`.
func SkipCount[T stream.Token, O any](count int, parser parser.Func[T, O]) parser.Func[T, any] {
	return combinator.Ignore(Count(count, parser))
}

// SkipCountMinMax parses `parser` from `min` to `max` times (including `min` and `max`)
// skipping the output of `parser`.
func SkipCountMinMax[T stream.Token, O any](min, max int, parser parser.Func[T, O]) parser.Func[T, any] {
	return combinator.Ignore(CountMinMax(min, max, parser))
}
