package combinator

import (
	"sync"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Lazy constructs the `parser` lazily.
//
// Can be used to effectively reduce the size of deeply nested parsers as only
// the function producing the parser is stored.
func Lazy[

	T stream.Token,
	O any](
	f func() parser.Func[T, O],
) parser.Func[T, O] {
	var p parser.Func[T, O]

	var init sync.Once

	return func(input []T) (out O, remaining []T, err error) {
		init.Do(func() { p = f() })

		return p(input)
	}
}
