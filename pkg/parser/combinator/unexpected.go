package combinator

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Unexpected always fails with `msg` as an unexpected error. Never consumes any input.
func Unexpected[S stream.Stream[T], T stream.Token, O any](msg string) parser.Func[S, T, O] {
	return func(input S) (out O, remaining S, err error) {
		remaining, err = input, fmt.Errorf("%s, %w", msg, parser.ErrUnexpected)
		return
	}
}
