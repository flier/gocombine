package char_test

import (
	"fmt"
	"unicode"

	"github.com/flier/gocombine/pkg/parser/char"
)

func ExampleString() {
	p := char.String[[]rune]("golang")

	fmt.Println(p([]rune("golang")))

	// Output:
	// golang [] <nil>
}

func ExampleStringFold() {
	p := char.StringFold[[]rune]("golang")

	fmt.Println(p([]rune("Golang")))

	// Output:
	// Golang [] <nil>
}

func ExampleStringCmp() {
	p := char.StringCmp[[]rune]("golang", func(l, r rune) bool { return unicode.ToLower(l) == unicode.ToLower(r) })

	fmt.Println(p([]rune("Golang")))

	// Output:
	// Golang [] <nil>
}
