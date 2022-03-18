package ranges_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/ranges"
)

func ExampleTakeUntil() {
	p := ranges.TakeUntil[[]byte]('\r')

	fmt.Println(p([]byte("To: user@example.com\r\n")))

	// Output:
	// [84 111 58 32 117 115 101 114 64 101 120 97 109 112 108 101 46 99 111 109] [13 10] <nil>
}
