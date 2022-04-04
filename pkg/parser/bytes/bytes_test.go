package bytes_test

import (
	"fmt"
	"unicode"

	"github.com/flier/gocombine/pkg/parser/bytes"
)

func ExampleBytes() {
	p := bytes.Bytes([]byte("golang"))

	fmt.Println(p([]byte("golang")))
	fmt.Println(p([]byte("goto")))

	// Output:
	// [103 111 108 97 110 103] [] <nil>
	// [] [103 111 116 111] bytes, bytes cmp, expected "golang", actual "got", unexpected
}

func ExampleFold() {
	p := bytes.Fold([]byte("golang"))

	fmt.Println(p([]byte("Golang")))
	fmt.Println(p([]byte("goto")))

	// Output:
	// [71 111 108 97 110 103] [] <nil>
	// [] [103 111 116 111] bytes fold, bytes cmp, expected "golang", actual "got", unexpected
}

func ExampleCmp() {
	p := bytes.Cmp([]byte("golang"), func(l, r byte) bool {
		return unicode.ToLower(rune(l)) == unicode.ToLower(rune(r))
	})

	fmt.Println(p([]byte("Golang")))
	fmt.Println(p([]byte("goto")))

	// Output:
	// [71 111 108 97 110 103] [] <nil>
	// [] [103 111 116 111] bytes cmp, expected "golang", actual "got", unexpected
}
