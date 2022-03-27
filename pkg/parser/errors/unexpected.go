package errors

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Unexpected always fails with `message` as an unexpected error.
// Never consumes any input.
func Unexpected[T stream.Token](message string) parser.Func[T, any] {
	return func(input []T) (r any, remaining []T, err error) {
		remaining, err = input, parser.UnexpectedMessage(message)

		return
	}
}
