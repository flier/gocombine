package to

import (
	"strconv"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Int parses a string-like data and converts it to an integer.
func Int[T stream.Token, O StringLike](p parser.Func[T, O]) parser.Func[T, int] {
	return parser.Expected(func(input []T) (out int, remaining []T, err error) {
		var o O

		if o, remaining, err = p(input); err != nil {
			return
		}

		s := Str(o)
		out, err = strconv.Atoi(s)

		return
	}, "int")
}
