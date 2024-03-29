package choice_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/to"
)

func ExampleOr() {
	p := choice.Or(
		to.String(char.String("let")),
		combinator.Map(char.Digit(), func(r rune) string { return "digit" }),
		to.String(char.String("led")),
	)

	fmt.Println(p([]rune("let")))
	fmt.Println(p([]rune("1")))
	fmt.Println(p([]rune("lost")))

	// Output:
	// let [] <nil>
	// digit [] <nil>
	//  [115 116] or, 3 errors occurred:
	// 	* map, string, cmp, expected "let", actual "lo", unexpected
	// 	* map, digit, satisfy, actual 'l', unexpected
	// 	* map, string, cmp, expected "led", actual "lo", unexpected
}
