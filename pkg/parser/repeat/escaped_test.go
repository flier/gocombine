package repeat_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/ranges"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/to"
	"github.com/flier/gocombine/pkg/parser/token"
)

func ExampleEscaped() {
	p := to.String(repeat.Escaped(
		ranges.TakeWhile1[[]rune](func(c rune) bool { return c != '"' && c != '\\' }),
		'\\',
		token.OneOf([]rune(`nrt\"`)),
	))

	fmt.Println(p([]rune(`ab\"12\n\rc"`)))

	// Output:
	// ab\"12\n\rc [34] <nil>
}
