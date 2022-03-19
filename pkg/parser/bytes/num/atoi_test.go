package num_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/bytes/num"
	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/repeat"
)

func ExampleAtoi() {
	p := num.Atoi(repeat.Many1(char.Digit[[]rune]()))

	fmt.Println(p([]rune("123abc")))

	// Output:
	// 123 [97 98 99] <nil>
}
