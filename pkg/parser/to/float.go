package to

import (
	"strconv"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

const bitsofFloat64 = 64

// Float convert the result of `parser` to a float64.
func Float[T stream.Token, O StringLike](parser parser.Func[T, O]) parser.Func[T, float64] {
	return combinator.AndThen(String(parser), func(s string) (float64, error) {
		return strconv.ParseFloat(s, bitsofFloat64)
	}).Expected("float")
}
