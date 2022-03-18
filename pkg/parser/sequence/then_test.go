package sequence_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
	"github.com/flier/gocombine/pkg/parser/token"
)

func ExampleThen() {
	comment := sequence.With(char.Spaces[[]rune](),
		repeat.Many(token.Satisfy[[]rune](func(c rune) bool {
			return c != '\n'
		})))

	p := combinator.Map(sequence.Then(token.Any[[]rune](), func(c rune) parser.Func[[]rune, rune, []rune] {
		if c == '#' {
			return comment
		}

		return combinator.Map(repeat.Many1(char.Letter[[]rune]()), func(s []rune) []rune {
			return append([]rune{c}, s...)
		})
	}), func(s []rune) string { return string(s) })

	fmt.Println(p([]rune("ac2")))
	fmt.Println(p([]rune("# ac2")))

	// Output:
	// ac [50] <nil>
	// ac2 [] <nil>
}
