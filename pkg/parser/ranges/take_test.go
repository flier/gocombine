package ranges_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/ranges"
)

func ExampleTake() {
	p := ranges.Take[byte](1)

	fmt.Println(p([]byte("!")))
	fmt.Println(p(nil))

	p = ranges.Take[byte](4)

	fmt.Println(p([]byte("1234abc")))
	fmt.Println(p([]byte("123")))

	// Output:
	// [33] [] <nil>
	// [] [] take, unexpected EOF
	// [49 50 51 52] [97 98 99] <nil>
	// [] [49 50 51] take, unexpected EOF
}

func ExampleTakeWhile() {
	p := ranges.TakeWhile(func(b byte) bool { return '0' <= b && b <= '9' })

	fmt.Println(p([]byte("123abc")))
	fmt.Println(p([]byte("abc")))

	// Output:
	// [49 50 51] [97 98 99] <nil>
	// [] [97 98 99] <nil>
}

func ExampleTakeWhile1() {
	p := ranges.TakeWhile1(func(b byte) bool { return '0' <= b && b <= '9' })

	fmt.Println(p([]byte("123abc")))
	fmt.Println(p([]byte("abc")))

	// Output:
	// [49 50 51] [97 98 99] <nil>
	// [] [97 98 99] take while1, one or more elements, expected
}

func ExampleTakeUntil() {
	p := ranges.TakeUntil(func(b byte) bool { return b == '\n' })

	fmt.Println(p([]byte("123\nabc")))
	fmt.Println(p([]byte("abc")))

	// Output:
	// [49 50 51] [10 97 98 99] <nil>
	// [97 98 99] [] take until, unexpected EOF
}

func ExampleTakeUntil1() {
	p := ranges.TakeUntil1(func(b byte) bool { return b == '\n' })

	fmt.Println(p([]byte("123\nabc")))
	fmt.Println(p([]byte("\n")))
	fmt.Println(p([]byte("abc")))

	// Output:
	// [49 50 51] [10 97 98 99] <nil>
	// [] [10] take until, one or more elements, expected
	// [97 98 99] [] take until, unexpected EOF
}
