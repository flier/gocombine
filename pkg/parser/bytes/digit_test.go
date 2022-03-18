package bytes_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/bytes"
)

func ExampleDigit() {
	p := bytes.Digit[[]byte]()

	fmt.Println(p([]byte("9")))
	fmt.Println(p([]byte("A")))

	// Output:
	// 57 [] <nil>
	// 65 [65] digit, satisfy, expected
}

func ExampleOctDigit() {
	p := bytes.OctDigit[[]byte]()

	fmt.Println(p([]byte("7")))
	fmt.Println(p([]byte("8")))

	// Output:
	// 55 [] <nil>
	// 56 [56] octal digit, satisfy, expected
}

func ExampleHexDigit() {
	p := bytes.HexDigit[[]byte]()

	fmt.Println(p([]byte("7")))
	fmt.Println(p([]byte("F")))
	fmt.Println(p([]byte("h")))

	// Output:
	// 55 [] <nil>
	// 70 [] <nil>
	// 104 [104] octal digit, satisfy, expected
}
