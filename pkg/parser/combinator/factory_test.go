package combinator_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
)

func ExampleFactory() {
	p := combinator.Factory(func(s []rune) parser.Func[rune, []rune] {
		if s[0] == 'a' {
			return char.String("apple")
		}

		return char.String("banana")
	})

	fmt.Println(p([]rune("apple")))
	fmt.Println(p([]rune("banana")))

	// Output:
	// [97 112 112 108 101] [] <nil>
	// [] [98 97 110 97 110 97] string, char cmp, expected "apple", actual "b", unexpected
}
