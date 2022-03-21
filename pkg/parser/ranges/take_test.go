package ranges_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/bytes"
	"github.com/flier/gocombine/pkg/parser/ranges"
)

func ExampleTake() {
	p := bytes.Take(1)

	fmt.Println(p([]byte("!")))
	fmt.Println(p(nil))

	p = bytes.Take(4)

	fmt.Println(p([]byte("1234abc")))
	fmt.Println(p([]byte("123")))

	// Output:
	// [33] [] <nil>
	// [] [] unexpected EOF
	// [49 50 51 52] [97 98 99] <nil>
	// [] [49 50 51] unexpected EOF
}

func ExampleTakeWhile() {
	p := ranges.TakeWhile[byte](func(b byte) bool { return '0' <= b && b <= '9' })

	fmt.Println(p([]byte("123abc")))
	fmt.Println(p([]byte("abc")))

	// Output:
	// [49 50 51] [97 98 99] <nil>
	// [] [97 98 99] <nil>
}

func ExampleTakeWhile1() {
	p := ranges.TakeWhile1[byte](func(b byte) bool { return '0' <= b && b <= '9' })

	fmt.Println(p([]byte("123abc")))
	fmt.Println(p([]byte("abc")))

	// Output:
	// [49 50 51] [97 98 99] <nil>
	// [] [97 98 99] one or more elements, expected
}
