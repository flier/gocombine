package to

import (
	"strconv"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Float convert the result of `parser` to a float64.
func Float[
	S stream.Stream[T],
	T stream.Token,
	O IntoString,
](parser parser.Func[S, T, O]) parser.Func[S, T, float64] {
	return combinator.AndThen(String(parser), func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	})
}
