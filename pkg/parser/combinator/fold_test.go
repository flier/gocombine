package combinator_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
)

func ExampleFold() {
	p := combinator.Fold(repeat.Many1(char.Digit()),
		func() int { return 0 },
		func(acc int, c rune) int { return acc*10 + int(c-'0') })

	fmt.Println(p([]rune("123")))
	fmt.Println(p([]rune("9999999999999999")))
	fmt.Println(p([]rune("foobar")))

	// Output:
	// 123 [] <nil>
	// 9999999999999999 [] <nil>
	// 0 [102 111 111 98 97 114] fold, many1, digit, satisfy, actual 'f', unexpected
}
