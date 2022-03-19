package repeat

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Until parses `parser` zero or more times until `end` is encountered or `end` indicates that it has committed input before failing.
func Until[
	S stream.Stream[T],
	T stream.Token,
	O, E any,
](
	parser parser.Func[S, T, O],
	end parser.Func[S, T, E],
) parser.Func[S, T, []O] {
	return func(input S) (out []O, remaining S, err error) {
		remaining = input

		for {
			if _, _, err = end(remaining); err == nil {
				break
			}

			var o O
			var rest S
			if o, rest, err = parser(remaining); err != nil {
				return
			}

			out = append(out, o)
			remaining = rest
		}

		return
	}
}
