package ranges_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/ranges"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
)

func ExampleRecognize() {
	p := char.AsString(ranges.Recognize(combinator.Pair(
		repeat.SkipMany1(char.Digit[[]rune]()),
		choice.Optional(
			sequence.With(char.Char[[]rune]('.'),
				repeat.SkipMany1(char.Digit[[]rune]()))),
	)))

	fmt.Println(p([]rune("1234!")))
	fmt.Println(p([]rune("1234.0001!")))
	fmt.Println(p([]rune("!")))
	fmt.Println(p([]rune("1234.")))

	// Output:
	// 1234 [33] <nil>
	// 1234.0001 [33] <nil>
	//  [33] digit, satisfy, expected
	// 1234 [46] <nil>
}
