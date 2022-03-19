package regex_test

import (
	"fmt"
	"regexp"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/regex"
)

func ExampleCaptures() {
	digits := char.AsStringSlice(regex.Captures[[]rune](regexp.MustCompile("([a-z]+):([0-9]+)")))

	fmt.Println(digits([]rune("test:123 field:456 ")))
	fmt.Println(digits([]rune("test:123 :456 ")))

	// Output:
	// [test:123 test 123] [32 102 105 101 108 100 58 52 53 54 32] <nil>
	// [test:123 test 123] [32 58 52 53 54 32] <nil>
}

func ExampleCapturesMany() {
	digits := combinator.Map(
		regex.CapturesMany[[]rune](regexp.MustCompile("([a-z]+)?:([0-9]+)")),
		func(captured [][][]rune) (r [][]string) {
			r = make([][]string, len(captured))

			for i, sub := range captured {
				r[i] = make([]string, len(sub))

				for j, s := range sub {
					r[i][j] = string(s)
				}
			}

			return
		},
	)

	fmt.Println(digits([]rune("test:123 field:456 ")))
	fmt.Println(digits([]rune("test:123 :456 ")))

	// Output:
	// [[test:123 test 123] [field:456 field 456]] [32] <nil>
	// [[test:123 test 123] [:456  456]] [32] <nil>
}
