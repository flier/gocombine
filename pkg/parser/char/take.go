package char

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Take reads a chars of length `n`.
func Take(n int) parser.Func[rune, []rune] {
	return parser.Expected(func(input []rune) (out []rune, remaining []rune, err error) {
		return stream.UnconsRange(input, n)
	}, "take")
}
