package char_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
)

func ExampleTake() {
	p := char.Take(2)

	fmt.Println(p([]rune("let")))
	fmt.Println(p([]rune("1")))

	// Output:
	// [108 101] [116] <nil>
	// [] [49] take, unexpected EOF
}
