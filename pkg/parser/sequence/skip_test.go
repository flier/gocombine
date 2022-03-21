package sequence_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
)

func ExampleSkip() {
	p := sequence.Skip(repeat.Many(char.Digit()), char.Char('i'))

	fmt.Println(p([]rune("123i456")))

	// Output:
	// [49 50 51] [52 53 54] <nil>
}
