package combinator_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
)

func ExampleMap() {
	p := combinator.Map(char.Digit(), func(r rune) bool { return r == '9' })

	fmt.Println(p.Parse([]rune("9i")))

	// Output:
	// true [105] <nil>
}
