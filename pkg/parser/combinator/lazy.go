package combinator

import (
	"sync"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Lazy constructs the `parser` lazily.
//
// Can be used to effectively reduce the size of deeply nested parsers as only the function producing the parser is stored.
func Lazy[
	S stream.Stream[T],
	T stream.Token,
	O any](
	f func() parser.Func[S, T, O],
) parser.Func[S, T, O] {
	var p parser.Func[S, T, O]
	var init sync.Once

	return func(input S) (out O, remaining S, err error) {
		init.Do(func() { p = f() })

		return p(input)
	}
}
