package ranges_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/bytes/be"
	"github.com/flier/gocombine/pkg/parser/ranges"
)

func ExampleLengthPrefix() {
	p := ranges.LengthPrefix(be.Uint16())

	fmt.Println(p([]byte{0x00, 0x03, 0x01, 0x02, 0x03, 0x04}))
	fmt.Println(p([]byte{0x00, 0x03, 0x01, 0x02}))

	// Output:
	// [1 2 3] [4] <nil>
	// [] [0 3 1 2] length prefix, unexpected EOF
}
