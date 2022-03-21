package to_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/to"
	"github.com/flier/gocombine/pkg/parser/token"
)

func ExampleBool() {
	p := to.Bool(choice.Or(
		to.String(char.StringFold("true")),
		to.String(char.StringFold("false")),
		to.String(token.OneOf([]rune("tTfF10"))),
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
