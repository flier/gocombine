package choice_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/to"
)

func ExampleOptional() {
	p := choice.Optional(to.String(char.String("hello")))

	fmt.Println(p([]rune("hello")))
	fmt.Println(p([]rune("world")))

	// Output:
	// hello [] <nil>
	// <none> [119 111 114 108 100] <nil>
}
