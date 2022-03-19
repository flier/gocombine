package regex_test

import (
	"fmt"
	"regexp"

	"github.com/flier/gocombine/pkg/parser/regex"
)

func ExampleMatch() {
	p := regex.Match[[]rune](regexp.MustCompile("[:alpha:]+"))

	fmt.Println(p([]rune("abc123")))
	fmt.Println(p([]rune("Bingo!")))
	fmt.Println(p(nil))

	// Output:
	// true [97 98 99 49 50 51] <nil>
	// false [66 105 110 103 111 33] <nil>
	// false [] <nil>
}
