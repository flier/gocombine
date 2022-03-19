package repeat

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// TakeUntil takes input until `end` is encountered or `end` indicates
// that it has committed input before failing.
func TakeUntil[
	S stream.Stream[T],
	T stream.Token,
	O any,
](
	end parser.Func[S, T, O],
) parser.Func[S, T, []T] {
	return func(input S) (out []T, remaining S, err error) {
		remaining = input

		for {
			if _, _, err = end(remaining); err == nil {
				break
			}

			var tok T
			var rest S
			tok, rest, err = stream.Uncons(remaining)
			if err != nil {
				return
			}

			out = append(out, tok)
			remaining = rest
		}

		return
	}
}

// SkipUntil skips input until `end` is encountered or `end` indicates
// that it has committed input before failing.
func SkipUntil[
	S stream.Stream[T],
	T stream.Token,
	O any,
](
	end parser.Func[S, T, O],
) parser.Func[S, T, any] {
	return combinator.Ignore(TakeUntil(end))
}
