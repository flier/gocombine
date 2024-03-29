package repeat

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Escaped parses an escaped string by first applying `p`
// which accept the normal characters which do not need escaping.
// Once `p` can not consume any more input it checks if the next token is `escape`.
// If it is then `escapeParser` is used to parse the escaped character and then resumes parsing using `p`.
// If `escape` was not found then the parser finishes successfully.
func Escaped[T stream.Token](p parser.Func[T, []T], escape T, escapeParser parser.Func[T, T]) parser.Func[T, []T] {
	return parser.Expected(func(input []T) (parsed []T, remaining []T, err error) {
		remaining = input

		for {
			var outs []T

			var rest []T

			if outs, rest, err = p(remaining); err != nil {
				var tok T

				if tok, rest, err = stream.Uncons(remaining); err != nil {
					return
				}

				if tok != escape {
					break
				}

				outs = []T{tok}

				if tok, rest, err = escapeParser(rest); err != nil {
					return
				}

				outs = append(outs, tok)
			}

			parsed = append(parsed, outs...)
			remaining = rest
		}

		return
	}, "escaped")
}
