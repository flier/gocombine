package stream_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/stream"
)

func ExampleDistance() {
	input := []rune{'a', 'b'}

	before := stream.Checkpoint(input)
	fmt.Println(stream.Distance(input, before))

	b, input, err := stream.Uncons(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(stream.Distance(input, before), b)

	b, input, err = stream.Uncons(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(stream.Distance(input, before), b)

	_, _, err = stream.Uncons(input)
	fmt.Println(err)

	// Output:
	// 0
	// 1 97
	// 2 98
	// unexpected EOF
}
