package choice_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/to"
)

func ExampleOptional() {
	p := choice.Optional(to.String(char.String[[]rune]("hello")))

	o, remaining, err := p([]rune("hello"))
	fmt.Println(o.Unwrap(), remaining, err)

	o, remaining, err = p([]rune("world"))
	fmt.Println(o.HasSome(), remaining, err)

	// Output:
	// hello [] <nil>
	// false [119 111 114 108 100] <nil>
}
