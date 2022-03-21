package combinator

import (
	"github.com/flier/gocombine/pkg/pair"
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Pair parses two heterogeneous value.
func Pair[T stream.Token, O1, O2 any](p1 parser.Func[T, O1], P2 parser.Func[T, O2]) parser.Func[T, pair.Pair[O1, O2]] {
	return Attempt(func(input []T) (out pair.Pair[O1, O2], remaining []T, err error) {
		var o1 O1
		if o1, remaining, err = p1.Parse(input); err != nil {
			return
		}

		var o2 O2
		if o2, remaining, err = P2.Parse(remaining); err != nil {
			return
		}

		out = pair.New(o1, o2)

		return
	})
}
