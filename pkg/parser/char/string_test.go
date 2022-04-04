package char_test

import (
	"fmt"
	"unicode"

	"github.com/flier/gocombine/pkg/parser/char"
)

func ExampleString() {
	p := char.String("apple")

	fmt.Println(p([]rune("apple")))
	fmt.Println(p(nil))

	// Output:
	// [97 112 112 108 101] [] <nil>
	// [] [] string, char cmp, unexpected EOF
}

func ExampleFold() {
	p := char.Fold("golang")

	fmt.Println(p([]rune("Golang")))
	fmt.Println(p([]rune("goto")))

	// Output:
	// [71 111 108 97 110 103] [] <nil>
	// [] [103 111 116 111] char fold, char cmp, expected "golang", actual "got", unexpected
}

func ExampleCmp() {
	p := char.Cmp("golang", func(l, r rune) bool {
		return unicode.ToLower(l) == unicode.ToLower(r)
	})

	fmt.Println(p([]rune("Golang")))
	fmt.Println(p([]rune("goto")))

	// Output:
	// [71 111 108 97 110 103] [] <nil>
	// [] [103 111 116 111] char cmp, expected "golang", actual "got", unexpected
}
