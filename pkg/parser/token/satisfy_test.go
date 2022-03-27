package token_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/token"
)

func ExampleSatisfy() {
	p := token.Satisfy(func(c rune) bool { return c == '!' || c == '?' })

	fmt.Println(p([]rune("!")))
	fmt.Println(p([]rune("?")))
	fmt.Println(p([]rune("#")))

	// Output:
	// 33 [] <nil>
	// 63 [] <nil>
	// 35 [35] satisfy, actual '#', unexpected
}

func ExampleSatisfyMap() {
	p := token.SatisfyMap(func(c rune) (bool, error) {
		switch c {
		case 'y', 'Y':
			return true, nil
		case 'n', 'N':
			return false, nil
		default:
			return false, parser.UnexpectedRange([]rune{'y', 'Y', 'n', 'N'}, []rune{c})
		}
	})

	fmt.Println(p([]rune("y")))
	fmt.Println(p([]rune("N")))
	fmt.Println(p([]rune("#")))

	// Output:
	// true [] <nil>
	// false [] <nil>
	// false [35] satisfy and map, expected
}
