package repeat

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// TakeUntil takes input until `end` is encountered or `end` indicates
// that it has committed input before failing.
func TakeUntil[T stream.Token, O any](end parser.Func[T, O]) parser.Func[T, []T] {
	return func(input []T) (out []T, remaining []T, err error) {
		remaining = input

		for {
			if _, _, err = end(remaining); err == nil {
				break
			}

			var tok T

			var rest []T

			if tok, rest, err = stream.Uncons(remaining); err != nil {
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
func SkipUntil[T stream.Token, O any](end parser.Func[T, O]) parser.Func[T, any] {
	return combinator.Ignore(TakeUntil(end))
}
