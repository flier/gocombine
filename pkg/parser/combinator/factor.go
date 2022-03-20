package combinator

import (
	"sync"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Constructs the `parser` lazily.
//
// This is similar to `Lazy` but it takes `Input` as an argument and allows different parsers to be returned on each call
// to `p` while still reporting the correct errors.
func Factor[
	S stream.Stream[T],
	T stream.Token,
	O any](
	f func(S) parser.Func[S, T, O],
) parser.Func[S, T, O] {
	var p parser.Func[S, T, O]
	var init sync.Once

	return func(input S) (out O, remaining S, err error) {
		init.Do(func() {
			p = f(input)
		})

		return p(input)
	}
}
