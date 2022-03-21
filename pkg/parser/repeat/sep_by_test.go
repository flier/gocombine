package repeat_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/repeat"
)

func ExampleSepBy() {
	p := repeat.SepBy(char.Digit(), char.Char(','))

	fmt.Println(p([]rune("1,2,3")))
	fmt.Println(p([]rune("")))

	// Output:
	// [49 50 51] [] <nil>
	// [] [] <nil>
}

func ExampleSepBy1() {
	p := repeat.SepBy1(char.Digit(), char.Char(','))

	fmt.Println(p([]rune("1,2,3")))
	fmt.Println(p([]rune("")))

	// Output:
	// [49 50 51] [] <nil>
	// [] [] digit, unexpected EOF
}

func ExampleSepEndBy() {
	p := repeat.SepEndBy(char.Digit(), char.Char(';'))

	fmt.Println(p([]rune("1;2;3;")))
	fmt.Println(p([]rune("")))

	// Output:
	// [49 50 51] [] <nil>
	// [] [] <nil>
}

func ExampleSepEndBy1() {
	p := repeat.SepEndBy1(char.Digit(), char.Char(';'))

	fmt.Println(p([]rune("1;2;3;")))
	fmt.Println(p([]rune("")))

	// Output:
	// [49 50 51] [] <nil>
	// [] [] digit, unexpected EOF
}
