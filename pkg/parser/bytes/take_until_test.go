package bytes_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/bytes"
)

func ExampleTakeUntil() {
	p := bytes.TakeUntil('\r')

	fmt.Println(p([]byte("To: user@example.com\r\n")))
	fmt.Println(p([]byte("foobar")))

	// Output:
	// [84 111 58 32 117 115 101 114 64 101 120 97 109 112 108 101 46 99 111 109] [13 10] <nil>
	// [] [] take until, expected [0x0d], actual "foobar", unexpected
}
