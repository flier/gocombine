package choice_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
)

func ExampleAnd() {
	p := choice.And(char.Digit(), char.Char('i'))

	fmt.Println(p([]rune("9i456")))
	fmt.Println(p([]rune("123")))

	// Output:
	// [57 105] [52 53 54] <nil>
	// [49 50] [49 50 51] and, char, expected 'i', actual '2', unexpected
}
