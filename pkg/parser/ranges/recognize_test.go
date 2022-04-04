package ranges_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/ranges"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
	"github.com/flier/gocombine/pkg/parser/to"
)

func ExampleRecognize() {
	p := to.String(ranges.Recognize(combinator.Pair(
		repeat.SkipMany1(char.Digit()),
		choice.Optional(
			sequence.With(char.Char('.'),
				repeat.SkipMany1(char.Digit()))),
	)))

	fmt.Println(p([]rune("1234!")))
	fmt.Println(p([]rune("1234.0001!")))
	fmt.Println(p([]rune("!")))
	fmt.Println(p([]rune("1234.")))

	// Output:
	// 1234 [33] <nil>
	// 1234.0001 [33] <nil>
	//  [33] map, recognize, pair, skip, ignore, many1, digit, satisfy, actual '!', unexpected
	// 1234 [46] <nil>
}
