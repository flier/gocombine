package repeat_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
	"github.com/flier/gocombine/pkg/parser/to"
)

func ExampleUntil() {
	p := repeat.Until(
		sequence.Skip(to.String(repeat.Many1(char.Letter())), char.Spaces()),
		char.Char('!'))

	fmt.Println(p([]rune("Hello World!")))

	// Output:
	// [Hello World] [33] <nil>
}
