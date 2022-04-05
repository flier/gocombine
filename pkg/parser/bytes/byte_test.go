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
	// 0 [] any, unexpected EOF
}

func ExampleByte() {
	p := bytes.Byte(byte('!'))

	fmt.Println(p([]byte("!")))
	fmt.Println(p([]byte("a")))
	fmt.Println(p([]byte("")))

	// Output:
	// 33 [] <nil>
	// 97 [] byte, expected '0x21', actual '0x61', unexpected
	// 0 [] byte, unexpected EOF
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
	// 33 [] whitespace, satisfy, actual '0x21', unexpected
	// 0 [] whitespace, satisfy, unexpected EOF
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
	// 13 [] newline, expected '0x0a', actual '0x0d', unexpected
	// 10 [] <nil>
}

func ExampleCrLf() {
	p := bytes.CrLf()

	fmt.Println(p([]byte("\r\n")))
	fmt.Println(p([]byte("\r")))
	fmt.Println(p([]byte("\n")))

	// Output:
	// [13 10] [] <nil>
	// [13] [] crlf, and, unexpected EOF
	// [10] [] crlf, and, expected '0x0d', actual '0x0a', unexpected
}

func ExampleTab() {
	p := bytes.Tab()

	fmt.Println(p([]byte("\t")))
	fmt.Println(p([]byte(" ")))

	// Output:
	// 9 [] <nil>
	// 32 [] tab, expected '0x09', actual '0x20', unexpected
}

func ExampleUpper() {
	p := bytes.Upper()

	fmt.Println(p([]byte("a")))
	fmt.Println(p([]byte("A")))

	// Output:
	// 97 [] uppercase letter, satisfy, actual '0x61', unexpected
	// 65 [] <nil>
}

func ExampleLower() {
	p := bytes.Lower()

	fmt.Println(p([]byte("a")))
	fmt.Println(p([]byte("A")))

	// Output:
	// 97 [] <nil>
	// 65 [] lowercase letter, satisfy, actual '0x41', unexpected
}

func ExampleLetter() {
	p := bytes.Letter()

	fmt.Println(p([]byte("a")))
	fmt.Println(p([]byte("A")))
	fmt.Println(p([]byte("9")))

	// Output:
	// 97 [] <nil>
	// 65 [] <nil>
	// 57 [] letter, satisfy, actual '0x39', unexpected
}

func ExampleAlphaNum() {
	p := bytes.AlphaNum()

	fmt.Println(p([]byte("a")))
	fmt.Println(p([]byte("9")))
	fmt.Println(p([]byte("!")))

	// Output:
	// 97 [] <nil>
	// 57 [] <nil>
	// 33 [] letter or digit, or, 2 errors occurred:
	// 	* letter, satisfy, actual '0x21', unexpected
	// 	* digit, satisfy, actual '0x21', unexpected
}
