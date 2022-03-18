package repeat

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

var ErrExpected = parser.ErrExpected

// Count parses `parser` from zero up to `count` times.
func Count[
	S stream.Stream[T],
	T stream.Token,
	O any,
](count int, parser parser.Func[S, T, O]) parser.Func[S, T, []O] {
	return CountMinMax(0, count, parser)
}

// CountMinMax parses `parser` from `min` to `max` times (including `min` and `max`).
func CountMinMax[
	S stream.Stream[T],
	T stream.Token,
	O any,
](min, max int, parser parser.Func[S, T, O]) parser.Func[S, T, []O] {
	return func(input S) (parsed []O, remaining S, err error) {
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
			err = fmt.Errorf("%d more elements, %w", min-len(parsed), ErrExpected)
		} else {
			err = nil
		}

		return
	}
}
