package combinator_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
)

func ExampleFlatMap() {
	p := combinator.FlatMap(
		char.Take(4),
		func(input []rune) (out []rune, err error) {
			out, _, err = repeat.Many1(char.Digit())(input)

			return
		},
	)

	fmt.Println(p([]rune("12abcd")))
	fmt.Println(p([]rune("123")))
	fmt.Println(p([]rune("foobar")))

	// Output:
	// [49 50] [99 100] <nil>
	// [] [49 50 51] flat map, take, unexpected EOF
	// [] [97 114] flat map, many1, digit, satisfy, actual 'f', unexpected
}
