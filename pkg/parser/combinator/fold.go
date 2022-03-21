package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Fold every element into an accumulator by applying an operation, returning the final result.
func Fold[T stream.Token, I, B any](parser parser.Func[T, []I], init func() B, f func(B, I)) parser.Func[T, B] {
	return func(input []T) (acc B, remaining []T, err error) {
		var items []I

		if items, remaining, err = parser(input); err != nil {
			remaining = input

			return
		}

		acc = init()
		for _, item := range items {
			f(acc, item)
		}

		return
	}
}
