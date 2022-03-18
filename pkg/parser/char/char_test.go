package char_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
)

func ExampleChar() {
	p := char.Char[[]rune]('!')

	fmt.Println(p([]rune("!")))
	fmt.Println(p([]rune("A")))

	// Output:
	// 33 [] <nil>
	// 65 [65] expected `[33]`, got `[65]`
}

func ExampleSpace() {
	p := char.Space[[]rune]()

	fmt.Println(p([]rune(" ")))
	fmt.Println(p([]rune("  ")))
	fmt.Println(p([]rune("!")))
	fmt.Println(p([]rune("")))

	// Output:
	// 32 [] <nil>
	// 32 [32] <nil>
	// 33 [33] whitespace, satisfy, expected
	// 0 [] whitespace, unexpected EOF
}

func ExampleSpaces() {
	p := char.Spaces[[]rune]()

	fmt.Println(p([]rune(" ")))
	fmt.Println(p([]rune("  ")))
	fmt.Println(p([]rune("!")))
	fmt.Println(p([]rune("")))

	// Output:
	// [32] [] <nil>
	// [32 32] [] <nil>
	// [] [33] <nil>
	// [] [] <nil>
}

func ExampleNewLine() {
	p := char.NewLine[[]rune]()

	fmt.Println(p([]rune("\r")))
	fmt.Println(p([]rune("\n")))

	// Output:
	// 13 [13] newline, expected `[10]`, got `[13]`
	// 10 [] <nil>
}

func ExampleCrLf() {
	p := char.CrLf[[]rune]()

	fmt.Println(p([]rune("\r\n")))
	fmt.Println(p([]rune("\r")))
	fmt.Println(p([]rune("\n")))

	// Output:
	// [13 10] [] <nil>
	// [13] [13] crlf, unexpected EOF
	// [10] [10] crlf, expected `[13]`, got `[10]`
}

func ExampleTab() {
	p := char.Tab[[]rune]()

	fmt.Println(p([]rune("\t")))
	fmt.Println(p([]rune(" ")))

	// Output:
	// 9 [] <nil>
	// 32 [32] tab, expected `[9]`, got `[32]`
}

func ExampleUpper() {
	p := char.Upper[[]rune]()

	fmt.Println(p([]rune("a")))
	fmt.Println(p([]rune("A")))

	// Output:
	// 97 [97] uppercase letter, satisfy, expected
	// 65 [] <nil>
}

func ExampleLower() {
	p := char.Lower[[]rune]()

	fmt.Println(p([]rune("a")))
	fmt.Println(p([]rune("A")))

	// Output:
	// 97 [] <nil>
	// 65 [65] lowercase letter, satisfy, expected
}

func ExampleLetter() {
	p := char.Letter[[]rune]()

	fmt.Println(p([]rune("a")))
	fmt.Println(p([]rune("A")))
	fmt.Println(p([]rune("9")))

	// Output:
	// 97 [] <nil>
	// 65 [] <nil>
	// 57 [57] letter, satisfy, expected
}

func ExampleAlphaNum() {
	p := char.AlphaNum[[]rune]()

	fmt.Println(p([]rune("a")))
	fmt.Println(p([]rune("9")))
	fmt.Println(p([]rune("!")))

	// Output:
	// 97 [] <nil>
	// 57 [] <nil>
	// 33 [33] letter or digit, 2 errors occurred:
	// 	* letter, satisfy, expected
	// 	* digit, satisfy, expected
}
