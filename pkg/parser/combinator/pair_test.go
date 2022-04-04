package combinator_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
)

func ExamplePair() {
	p := combinator.Pair(choice.Optional(char.OneOf("+-")), repeat.Many1(char.Digit()))

	fmt.Println(p([]rune("123")))
	fmt.Println(p([]rune("+123")))
	fmt.Println(p([]rune("foobar")))

	// Output:
	// {<none> [49 50 51]} [] <nil>
	// {43 [49 50 51]} [] <nil>
	// {<none> []} [102 111 111 98 97 114] pair, many1, digit, satisfy, actual 'f', unexpected
}
