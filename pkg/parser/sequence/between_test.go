package sequence_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/sequence"
)

func ExampleBetween() {
	p := sequence.Between(char.Char[[]rune]('['), char.Char[[]rune](']'), char.String[[]rune]("Golang"))

	fmt.Println(p.Parse([]rune("[Golang]")))

	// Output:
	// Golang [] <nil>
}
