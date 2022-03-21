package repeat

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Count parses `parser` from zero up to `count` times.
func Count[
	S stream.Stream[T],
	T stream.Token,
	O any,
](
	count int,
	parser parser.Func[S, T, O],
) parser.Func[S, T, []O] {
	return CountMinMax(0, count, parser)
}

var errExpected = parser.ErrExpected

// CountMinMax parses `parser` from `min` to `max` times (including `min` and `max`).
func CountMinMax[
	S stream.Stream[T],
	T stream.Token,
	O any,
](
	min, max int,
	parser parser.Func[S, T, O],
) parser.Func[S, T, []O] {
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
			err = fmt.Errorf("%d more elements, %w", min-len(parsed), errExpected)
		} else {
			err = nil
		}

		return
	}
}

// SkipCount parses `parser` from zero up to `count` times skipping the output of `parser`.
func SkipCount[
	S stream.Stream[T],
	T stream.Token,
	O any,
](
	count int,
	parser parser.Func[S, T, O],
) parser.Func[S, T, any] {
	return combinator.Ignore(Count(count, parser))
}

// SkipCountMinMax parses `parser` from `min` to `max` times (including `min` and `max`)
// skipping the output of `parser`.
func SkipCountMinMax[
	S stream.Stream[T],
	T stream.Token,
	O any,
](
	min, max int,
	parser parser.Func[S, T, O],
) parser.Func[S, T, any] {
	return combinator.Ignore(CountMinMax(min, max, parser))
}
