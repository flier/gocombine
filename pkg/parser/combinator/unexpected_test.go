package combinator_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
)

func ExampleUnexpected() {
	p := choice.Or(char.Char[[]rune]('a'), combinator.Unexpected[[]rune, rune, rune]("token"))

	fmt.Println(p([]rune("b")))

	// Output:
	// 0 [98] 2 errors occurred:
	// 	* expected `[97]`, actual `[98]`, unexpected
	// 	* token, unexpected
}
