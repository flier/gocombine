package combinator_test

import (
	"fmt"
	"strconv"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
)

func ExampleAndThen() {
	p := combinator.AndThen(repeat.Many1(char.Digit[[]rune]()), func(s []rune) (int, error) {
		return strconv.Atoi(string(s))
	})

	fmt.Println(p([]rune("123")))
	fmt.Println(p([]rune("9999999999999999")))

	// Output:
	// 123 [] <nil>
	// 9223372036854775807 [57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57] strconv.Atoi: parsing "999999999999999999999999": value out of range
}
