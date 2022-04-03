package char_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
)

func ExampleAny() {
	p := char.Any()

	fmt.Println(p([]rune("apple")))
	fmt.Println(p(nil))

	// Output:
	// 97 [112 112 108 101] <nil>
	// 0 [] unexpected EOF
}

func ExampleChar() {
	p := char.Char('!')

	fmt.Println(p([]rune("!")))
	fmt.Println(p([]rune("A")))

	// Output:
	// 33 [] <nil>
	// 65 [65] expected '!', actual 'A', unexpected
}

func ExampleSpace() {
	p := char.Space()

	fmt.Println(p([]rune(" ")))
	fmt.Println(p([]rune("  ")))
	fmt.Println(p([]rune("!")))
	fmt.Println(p([]rune("")))

	// Output:
	// 32 [] <nil>
	// 32 [32] <nil>
	// 33 [33] whitespace, satisfy, actual '!', unexpected
	// 0 [] whitespace, unexpected EOF
}

func ExampleSpaces() {
	p := char.Spaces()

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
	p := char.NewLine()

	fmt.Println(p([]rune("\r")))
	fmt.Println(p([]rune("\n")))

	// Output:
	// 13 [13] newline, expected '\n', actual '\r', unexpected
	// 10 [] <nil>
}

func ExampleCrLf() {
	p := char.CrLf()

	fmt.Println(p([]rune("\r\n")))
	fmt.Println(p([]rune("\r")))
	fmt.Println(p([]rune("\n")))

	// Output:
	// [13 10] [] <nil>
	// [13] [13] crlf, unexpected EOF
	// [10] [10] crlf, expected '\r', actual '\n', unexpected
}

func ExampleTab() {
	p := char.Tab()

	fmt.Println(p([]rune("\t")))
	fmt.Println(p([]rune(" ")))

	// Output:
	// 9 [] <nil>
	// 32 [32] tab, expected '\t', actual ' ', unexpected
}

func ExampleUpper() {
	p := char.Upper()

	fmt.Println(p([]rune("a")))
	fmt.Println(p([]rune("A")))

	// Output:
	// 97 [97] uppercase letter, satisfy, actual 'a', unexpected
	// 65 [] <nil>
}

func ExampleLower() {
	p := char.Lower()

	fmt.Println(p([]rune("a")))
	fmt.Println(p([]rune("A")))

	// Output:
	// 97 [] <nil>
	// 65 [65] lowercase letter, satisfy, actual 'A', unexpected
}

func ExampleLetter() {
	p := char.Letter()

	fmt.Println(p([]rune("a")))
	fmt.Println(p([]rune("A")))
	fmt.Println(p([]rune("9")))

	// Output:
	// 97 [] <nil>
	// 65 [] <nil>
	// 57 [57] letter, satisfy, actual '9', unexpected
}

func ExampleAlphaNum() {
	p := char.AlphaNum()

	fmt.Println(p([]rune("a")))
	fmt.Println(p([]rune("9")))
	fmt.Println(p([]rune("!")))

	// Output:
	// 97 [] <nil>
	// 57 [] <nil>
	// 33 [33] letter or digit, or, 2 errors occurred:
	// 	* letter, satisfy, actual '!', unexpected
	// 	* digit, satisfy, actual '!', unexpected
}
