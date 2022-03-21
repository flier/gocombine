package regex_test

import (
	"fmt"
	"regexp"

	"github.com/flier/gocombine/pkg/parser/regex"
	"github.com/flier/gocombine/pkg/parser/to"
)

func ExampleFind() {
	digits := to.String(regex.Find[rune](regexp.MustCompile("^[0-9]+")))

	fmt.Println(digits([]rune("123 456 ")))

	digits2 := to.String(regex.Find[rune](regexp.MustCompile("[0-9]+")))

	fmt.Println(digits2([]rune("123 456 ")))
	fmt.Println(digits2([]rune("abcd 123 456 ")))

	// Output:
	// 123 [32 52 53 54 32] <nil>
	// 123 [32 52 53 54 32] <nil>
	// 123 [32 52 53 54 32] <nil>
}

func ExampleFindMany() {
	digits2 := to.StringSlice(regex.FindMany[rune](regexp.MustCompile("[0-9]+")))

	fmt.Println(digits2([]rune("123 456 ")))
	fmt.Println(digits2([]rune("abcd 123 456 ")))
	fmt.Println(digits2([]rune("abcd")))

	// Output:
	// [123 456] [32] <nil>
	// [123 456] [32] <nil>
	// [] [97 98 99 100] <nil>
}
