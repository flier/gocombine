package errors_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/errors"
)

func ExampleUnexpected() {
	p := choice.Or(combinator.Ignore(char.Char('a')), errors.Unexpected[rune]("token"))

	fmt.Println(p([]rune("b")))

	// Output:
	// <nil> [98] 2 errors occurred:
	// 	* expected 'a', actual 'b', unexpected
	// 	* token, unexpected
}
