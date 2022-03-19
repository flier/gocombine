package repeat_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
)

func ExampleUntil() {
	p := repeat.Until(
		sequence.Skip(char.AsString(repeat.Many1(char.Letter[[]rune]())), char.Spaces[[]rune]()),
		char.Char[[]rune]('!'))

	fmt.Println(p([]rune("Hello World!")))

	// Output:
	// [Hello World] [33] <nil>
}
