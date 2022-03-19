package repeat_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/bytes"
	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/repeat"
)

func ExampleTakeUntil() {
	char_parser := repeat.TakeUntil(char.Digit[[]rune]())

	fmt.Println(char_parser([]rune("abc123")))

	byte_parser := repeat.TakeUntil(bytes.Bytes[[]byte]([]byte("TAG")))

	fmt.Println(byte_parser([]byte("123TAG")))

	// Output:
	// [97 98 99] [49 50 51] <nil>
	// [49 50 51] [84 65 71] <nil>
}

func ExampleSkipUntil() {
	char_parser := repeat.SkipUntil(char.Digit[[]rune]())

	fmt.Println(char_parser([]rune("abc123")))

	byte_parser := repeat.SkipUntil(bytes.Bytes[[]byte]([]byte("TAG")))

	fmt.Println(byte_parser([]byte("123TAG")))

	// Output:
	// <nil> [49 50 51] <nil>
	// <nil> [84 65 71] <nil>
}
