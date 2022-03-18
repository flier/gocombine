package repeat_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/repeat"
)

func ExampleMany() {
	p := repeat.Many(char.Digit[[]rune]())

	fmt.Println(p([]rune("123A")))

	// Output:
	// [49 50 51] [65] <nil>
}
