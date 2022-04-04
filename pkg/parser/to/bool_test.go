package to_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/to"
)

func ExampleBool() {
	p := to.Bool(choice.Or(
		to.String(char.Fold("true")),
		to.String(char.Fold("false")),
		to.String(char.OneOf("tTfF10")),
	))

	fmt.Println(p([]rune("t")))
	fmt.Println(p([]rune("0")))
	fmt.Println(p([]rune("True")))
	fmt.Println(p([]rune("FALSE")))

	// Output:
	// true [] <nil>
	// false [] <nil>
	// true [] <nil>
	// false [] <nil>
}
