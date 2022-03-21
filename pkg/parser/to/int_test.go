package to_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/to"
)

func ExampleInt() {
	p := to.Int(repeat.Many1(char.Digit()))

	fmt.Println(p([]rune("123abc")))

	// Output:
	// 123 [97 98 99] <nil>
}
