package to_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/to"
)

func ExampleFloat() {
	p := to.Float(char.Float[[]rune]())

	fmt.Println(p([]rune("3.1415926535")))
	fmt.Println(p([]rune("NaN")))
	fmt.Println(p([]rune("nan")))
	fmt.Println(p([]rune("inf")))
	fmt.Println(p([]rune("+Inf")))
	fmt.Println(p([]rune("-Inf")))
	fmt.Println(p([]rune("-0")))
	fmt.Println(p([]rune("+0")))

	// Output:
	// 3.1415926535 [] <nil>
	// NaN [] <nil>
	// NaN [] <nil>
	// +Inf [] <nil>
	// +Inf [] <nil>
	// -Inf [] <nil>
	// -0 [] <nil>
	// 0 [] <nil>
}
