package repeat_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/token"
)

func ExampleCount() {
	p := repeat.Count(2, token.Token('a'))

	fmt.Println(p([]rune("aaab")))
	fmt.Println(p([]rune("foobar")))

	// Output:
	// [97 97] [97 98] <nil>
	// [] [102 111 111 98 97 114] <nil>
}

func ExampleCountMinMax() {
	p := repeat.CountMinMax(1, 2, token.Token('a'))

	fmt.Println(p([]rune("aaab")))
	fmt.Println(p([]rune("foobar")))

	// Output:
	// [97 97] [97 98] <nil>
	// [] [102 111 111 98 97 114] 1 more elements, expected
}

func ExampleSkipCount() {
	p := repeat.SkipCount(2, token.Token('a'))

	fmt.Println(p([]rune("aaab")))
	fmt.Println(p([]rune("foobar")))

	// Output:
	// <nil> [97 98] <nil>
	// <nil> [102 111 111 98 97 114] <nil>
}

func ExampleSkipCountMinMax() {
	p := repeat.SkipCountMinMax(1, 2, token.Token('a'))

	fmt.Println(p([]rune("aaab")))
	fmt.Println(p([]rune("foobar")))

	// Output:
	// <nil> [97 98] <nil>
	// <nil> [102 111 111 98 97 114] 1 more elements, expected
}
