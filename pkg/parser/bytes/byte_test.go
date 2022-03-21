package bytes_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/bytes"
)

func ExampleByte() {
	p := bytes.Byte(byte('!'))

	fmt.Println(p([]byte("!")))
	fmt.Println(p([]byte("a")))
	fmt.Println(p([]byte("")))

	// Output:
	// 33 [] <nil>
	// 97 [97] expected `[33]`, actual `[97]`, unexpected
	// 0 [] unexpected EOF
}
