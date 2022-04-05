package combinator_test

import (
	"fmt"
	"strconv"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
)

func ExampleAndThen() {
	p := combinator.AndThen(repeat.Many1(char.Digit()), func(s []rune) (int, error) {
		return strconv.Atoi(string(s))
	})

	fmt.Println(p([]rune("123")))
	fmt.Println(p([]rune("9999999999999999")))
	fmt.Println(p([]rune("foobar")))

	// Output:
	// 123 [] <nil>
	// 9999999999999999 [] <nil>
	// 0 [111 111 98 97 114] and then, many1, digit, satisfy, actual 'f', unexpected
}
