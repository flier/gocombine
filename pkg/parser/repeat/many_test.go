package repeat_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/repeat"
)

func ExampleMany() {
	p := repeat.Many(char.Digit())

	fmt.Println(p([]rune("123A")))

	// Output:
	// [49 50 51] [65] <nil>
}

func ExampleMany1() {
	p := repeat.Many1(char.Digit())

	fmt.Println(p([]rune("1")))
	fmt.Println(p([]rune("123A")))
	fmt.Println(p([]rune("A")))

	// Output:
	// [49] [] <nil>
	// [49 50 51] [65] <nil>
	// [] [] many1, digit, satisfy, actual 'A', unexpected
}

func ExampleSkipMany() {
	p := repeat.SkipMany(char.Digit())

	fmt.Println(p([]rune("123A")))

	// Output:
	// <nil> [65] <nil>
}

func ExampleSkipMany1() {
	p := repeat.SkipMany1(char.Digit())

	fmt.Println(p([]rune("1")))
	fmt.Println(p([]rune("123A")))
	fmt.Println(p([]rune("A")))

	// Output:
	// <nil> [] <nil>
	// <nil> [65] <nil>
	// <nil> [] skip, ignore, many1, digit, satisfy, actual 'A', unexpected
}
