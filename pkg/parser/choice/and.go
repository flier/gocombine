package choice

import (
	"errors"
	"io"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// And parses with `parsers`.
// Succeeds if all parsers succeed, otherwise fails.
// Returns a slice with all values on success.
func And[T stream.Token, O any](parsers ...parser.Func[T, O]) parser.Func[T, []O] {
	return parser.Expected(func(input []T) (out []O, remaining []T, err error) {
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
	}, "and")
}
