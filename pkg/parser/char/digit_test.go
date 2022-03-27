package char_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
)

func ExampleDigit() {
	p := char.Digit()

	fmt.Println(p([]rune("9")))
	fmt.Println(p([]rune("A")))

	// Output:
	// 57 [] <nil>
	// 65 [65] digit, satisfy, actual 'A', unexpected
}

func ExampleOctDigit() {
	p := char.OctDigit()

	fmt.Println(p([]rune("7")))
	fmt.Println(p([]rune("8")))

	// Output:
	// 55 [] <nil>
	// 56 [56] octal digit, satisfy, actual '8', unexpected
}

func ExampleHexDigit() {
	p := char.HexDigit()

	fmt.Println(p([]rune("7")))
	fmt.Println(p([]rune("F")))
	fmt.Println(p([]rune("h")))

	// Output:
	// 55 [] <nil>
	// 70 [] <nil>
	// 104 [104] octal digit, satisfy, actual 'h', unexpected
}
