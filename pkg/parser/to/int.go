package to

import (
	"strconv"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Atoi parses a string-like data and converts it to an integer.
func Int[
	S stream.Stream[T],
	T stream.Token,
	O IntoString,
](parser parser.Func[S, T, O]) parser.Func[S, T, int] {
	return func(input S) (out int, remaining S, err error) {
		var o O
		o, remaining, err = parser(input)
		if err != nil {
			return
		}

		s := ToString(o)
		out, err = strconv.Atoi(s)
		return
	}
}
