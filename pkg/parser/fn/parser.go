package fn

import (
	"sync"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

type FnParser[S stream.Stream[T], T stream.Token, O any] struct {
	init   sync.Once
	factor func() parser.Func[S, T, O]
	parser parser.Func[S, T, O]
}

func (p *FnParser[S, T, O]) Parse(input S) (out O, remaining S, err error) {
	return p.parser.Parse(input)
}

// Parser convert a `parser` to function.
func Parser[
	S stream.Stream[T],
	T stream.Token,
	O any](
	parser parser.Parser[S, T, O],
) parser.Func[S, T, O] {
	return func(input S) (out O, remaining S, err error) {
		return parser.Parse(input)
	}
}
