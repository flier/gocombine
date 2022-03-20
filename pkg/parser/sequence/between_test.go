package sequence_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/sequence"
	"github.com/flier/gocombine/pkg/parser/to"
)

func ExampleBetween() {
	p := sequence.Between(
		char.Char[[]rune]('['),
		char.Char[[]rune](']'),
		to.String(char.String[[]rune]("Golang")))

	fmt.Println(p.Parse([]rune("[Golang]")))

	// Output:
	// Golang [] <nil>
}
