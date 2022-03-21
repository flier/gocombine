package bytes_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/bytes"
)

func ExampleAny() {
	p := bytes.Any()

	fmt.Println(p([]byte("apple")))
	fmt.Println(p(nil))

	// Output:
	// 97 [112 112 108 101] <nil>
	// 0 [] unexpected EOF
}

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

func ExampleSpace() {
	p := bytes.Space()

	fmt.Println(p([]byte(" ")))
	fmt.Println(p([]byte("  ")))
	fmt.Println(p([]byte("!")))
	fmt.Println(p([]byte("")))

	// Output:
	// 32 [] <nil>
	// 32 [32] <nil>
	// 33 [33] whitespace, satisfy, expected
	// 0 [] whitespace, unexpected EOF
}

func ExampleSpaces() {
	p := bytes.Spaces()

	fmt.Println(p([]byte(" ")))
	fmt.Println(p([]byte("  ")))
	fmt.Println(p([]byte("!")))
	fmt.Println(p([]byte("")))

	// Output:
	// [32] [] <nil>
	// [32 32] [] <nil>
	// [] [33] <nil>
	// [] [] <nil>
}

func ExampleNewLine() {
	p := bytes.NewLine()

	fmt.Println(p([]byte("\r")))
	fmt.Println(p([]byte("\n")))

	// Output:
	// 13 [13] newline, expected `[10]`, actual `[13]`, unexpected
	// 10 [] <nil>
}

func ExampleCrLf() {
	p := bytes.CrLf()

	fmt.Println(p([]byte("\r\n")))
	fmt.Println(p([]byte("\r")))
	fmt.Println(p([]byte("\n")))

	// Output:
	// [13 10] [] <nil>
	// [13] [13] crlf, unexpected EOF
	// [10] [10] crlf, expected `[13]`, actual `[10]`, unexpected
}

func ExampleTab() {
	p := bytes.Tab()

	fmt.Println(p([]byte("\t")))
	fmt.Println(p([]byte(" ")))

	// Output:
	// 9 [] <nil>
	// 32 [32] tab, expected `[9]`, actual `[32]`, unexpected
}

func ExampleUpper() {
	p := bytes.Upper()

	fmt.Println(p([]byte("a")))
	fmt.Println(p([]byte("A")))

	// Output:
	// 97 [97] uppercase letter, satisfy, expected
	// 65 [] <nil>
}

func ExampleLower() {
	p := bytes.Lower()

	fmt.Println(p([]byte("a")))
	fmt.Println(p([]byte("A")))

	// Output:
	// 97 [] <nil>
	// 65 [65] lowercase letter, satisfy, expected
}

func ExampleLetter() {
	p := bytes.Letter()

	fmt.Println(p([]byte("a")))
	fmt.Println(p([]byte("A")))
	fmt.Println(p([]byte("9")))

	// Output:
	// 97 [] <nil>
	// 65 [] <nil>
	// 57 [57] letter, satisfy, expected
}

func ExampleAlphaNum() {
	p := bytes.AlphaNum()

	fmt.Println(p([]byte("a")))
	fmt.Println(p([]byte("9")))
	fmt.Println(p([]byte("!")))

	// Output:
	// 97 [] <nil>
	// 57 [] <nil>
	// 33 [33] letter or digit, 2 errors occurred:
	// 	* letter, satisfy, expected
	// 	* digit, satisfy, expected
}
