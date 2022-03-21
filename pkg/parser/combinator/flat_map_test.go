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
			out, _, err = repeat.Many(char.Digit()).Parse(input)

			return
		},
	)

	fmt.Println(p([]rune("12abcd")))

	// Output:
	// [49 50] [99 100] <nil>
}
