package char_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/to"
)

func ExampleFloat() {
	p := to.Float(char.Float())

	fmt.Println(p([]rune("0")))
	fmt.Println(p([]rune("123")))
	fmt.Println(p([]rune("3.1415926")))
	fmt.Println(p([]rune("31415926e-7")))
	fmt.Println(p([]rune("nan")))
	fmt.Println(p([]rune("-inf")))
	fmt.Println(p([]rune("foobar")))

	// Output:
	// 0 [] <nil>
	// 123 [] <nil>
	// 3.1415926 [] <nil>
	// 3.1415926 [] <nil>
	// NaN [] <nil>
	// -Inf [] <nil>
	// 0 [111 111 98 97 114] float, and then, map, recognize, pair, or, 3 errors occurred:
	// 	* char fold, cmp, expected "nan", actual "f", unexpected
	// 	* char fold, cmp, expected "inf", actual "f", unexpected
	// 	* recognize, tuple3, many1, digit, satisfy, actual 'f', unexpected
}
