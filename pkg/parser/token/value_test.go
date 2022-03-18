package token_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/token"
)

func ExampleValue() {
	p := token.Value[[]rune]('a')

	fmt.Println(p([]rune("pple")))

	// Output:
	// 97 [112 112 108 101] <nil>
}

func ExampleProduce() {
	p := token.Produce[[]rune](func() string { return "foo" })

	fmt.Println(p([]rune("bar")))

	// Output:
	// foo [98 97 114] <nil>
}
