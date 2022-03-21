package to

import (
	"strconv"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Bool convert the result of `parser` to a bool.
func Bool[
	S stream.Stream[T],
	T stream.Token,
	O StringLike,
](parser parser.Func[S, T, O]) parser.Func[S, T, bool] {
	return combinator.AndThen(String(parser), func(s string) (bool, error) {
		return strconv.ParseBool(s)
	})
}
