package repeat_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/ranges"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/to"
)

func ExampleEscaped() {
	p := to.String(repeat.Escaped(
		ranges.TakeWhile1(func(c rune) bool { return c != '"' && c != '\\' }),
		'\\',
		char.OneOf(`nrt\"`),
	))

	fmt.Println(p([]rune(`ab\"12\n\rc"`)))
	fmt.Println(p([]rune(`\a`)))

	// Output:
	// ab\"12\n\rc [34] <nil>
	//  [92 97] map, escaped, one of, one of "nrt\\\"", satisfy, actual 'a', unexpected
}
