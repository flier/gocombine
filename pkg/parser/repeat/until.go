package repeat

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Until parses `parser` zero or more times until `end` is encountered
// or `end` indicates that it has committed input before failing.
func Until[T stream.Token, O, E any](parser parser.Func[T, O], end parser.Func[T, E]) parser.Func[T, []O] {
	return func(input []T) (out []O, remaining []T, err error) {
		remaining = input

		for {
			if _, _, err = end(remaining); err == nil {
				break
			}

			var o O

			var rest []T

			if o, rest, err = parser(remaining); err != nil {
				return
			}

			out = append(out, o)
			remaining = rest
		}

		return
	}
}
