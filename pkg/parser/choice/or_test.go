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
		to.String(char.String[[]rune]("let")),
		combinator.Map(char.Digit[[]rune](), func(r rune) string { return "digit" }),
		to.String(char.String[[]rune]("led")),
	)

	fmt.Println(p.Parse([]rune("let")))
	fmt.Println(p.Parse([]rune("1")))
	fmt.Println(p.Parse([]rune("lost")))

	// Output:
	// let [] <nil>
	// digit [] <nil>
	//  [108 111 115 116] 3 errors occurred:
	// 	* expected `[108 101 116]`, actual `[108 111]`, unexpected
	// 	* digit, satisfy, expected
	// 	* expected `[108 101 100]`, actual `[108 111]`, unexpected

}
