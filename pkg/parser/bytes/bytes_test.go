package bytes_test

import (
	"fmt"
	"unicode"

	"github.com/flier/gocombine/pkg/parser/bytes"
)

func ExampleBytes() {
	p := bytes.Bytes([]byte("golang"))

	fmt.Println(p([]byte("golang")))

	// Output:
	// [103 111 108 97 110 103] [] <nil>
}

func ExampleFold() {
	p := bytes.Fold([]byte("golang"))

	fmt.Println(p([]byte("Golang")))

	// Output:
	// [71 111 108 97 110 103] [] <nil>
}

func ExampleCmp() {
	p := bytes.Cmp([]byte("golang"), func(l, r byte) bool {
		return unicode.ToLower(rune(l)) == unicode.ToLower(rune(r))
	})

	fmt.Println(p([]byte("Golang")))

	// Output:
	// [71 111 108 97 110 103] [] <nil>
}
