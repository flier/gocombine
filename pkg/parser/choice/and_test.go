package choice_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
)

func ExampleAnd() {
	p := choice.And(char.Digit(), char.Char('i'))

	fmt.Println(p.Parse([]rune("9i456")))

	// Output:
	// [57 105] [52 53 54] <nil>
}
