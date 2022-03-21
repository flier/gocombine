package combinator

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Unexpected always fails with `msg` as an unexpected error. Never consumes any input.
func Unexpected[T stream.Token, O any](msg string) parser.Func[T, O] {
	return func(input []T) (out O, remaining []T, err error) {
		remaining, err = input, fmt.Errorf("%s, %w", msg, parser.ErrUnexpected)

		return
	}
}
