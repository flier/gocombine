package bytes_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/bytes"
)

func ExampleDigit() {
	p := bytes.Digit()

	fmt.Println(p([]byte("9")))
	fmt.Println(p([]byte("A")))

	// Output:
	// 57 [] <nil>
	// 65 [65] digit, satisfy, actual '0x41', unexpected
}

func ExampleOctDigit() {
	p := bytes.OctDigit()

	fmt.Println(p([]byte("7")))
	fmt.Println(p([]byte("8")))

	// Output:
	// 55 [] <nil>
	// 56 [56] octal digit, satisfy, actual '0x38', unexpected
}

func ExampleHexDigit() {
	p := bytes.HexDigit()

	fmt.Println(p([]byte("7")))
	fmt.Println(p([]byte("F")))
	fmt.Println(p([]byte("h")))

	// Output:
	// 55 [] <nil>
	// 70 [] <nil>
	// 104 [104] octal digit, satisfy, actual '0x68', unexpected
}
