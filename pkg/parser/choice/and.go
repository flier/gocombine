package choice

import (
	"errors"
	"io"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// And parses with `parsers`.
// Succeeds if all parsers succeed, otherwise fails.
// Returns a slice with all values on success.
func And[
	S stream.Stream[T],
	T stream.Token,
	O any,
](parsers ...parser.Func[S, T, O]) parser.Func[S, T, []O] {
	return combinator.Attempt(func(input S) (out []O, remaining S, err error) {
		out = make([]O, len(parsers))

		remaining = input
		for i, p := range parsers {
			out[i], remaining, err = p(remaining)
			if err != nil {
				if errors.Is(err, io.ErrUnexpectedEOF) {
					out = out[:i]
				} else {
					out = out[:i+1]
				}
				break
			}
		}

		return
	})
}
