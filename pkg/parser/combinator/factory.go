package combinator

import (
	"sync"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Factory constructs the `parser` lazily.
//
// This is similar to `Lazy` but it takes `input` as an argument and allows different parsers to be returned
// on each call to `p` while still reporting the correct errors.
func Factory[T stream.Token, O any](f func([]T) parser.Func[T, O]) parser.Func[T, O] {
	var p parser.Func[T, O]

	var init sync.Once

	return func(input []T) (out O, remaining []T, err error) {
		init.Do(func() {
			p = f(input)
		})

		return p(input)
	}
}
