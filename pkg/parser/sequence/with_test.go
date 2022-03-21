package sequence_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
)

func ExampleWith() {
	p := sequence.With(repeat.Many(char.Digit()), char.Char('i'))

	fmt.Println(p.Parse([]rune("123i456")))

	// Output:
	// 105 [52 53 54] <nil>
}
