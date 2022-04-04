package bytes_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/bytes"
)

func ExampleTake() {
	p := bytes.Take(2)

	fmt.Println(p([]byte("let")))
	fmt.Println(p([]byte("1")))

	// Output:
	// [108 101] [116] <nil>
	// [] [49] take, unexpected EOF
}

func ExampleTakeOf() {
	p := bytes.TakeOf[uint16]()

	fmt.Println(p([]byte{0x01, 0x02, 0x03, 0x04}))
	fmt.Println(p([]byte{0x01}))

	// Output:
	// [1 2] [3 4] <nil>
	// [] [1] take of, take, unexpected EOF
}
