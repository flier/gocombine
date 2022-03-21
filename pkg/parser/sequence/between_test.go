package sequence_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/sequence"
	"github.com/flier/gocombine/pkg/parser/to"
)

func ExampleBetween() {
	p := sequence.Between(
		char.Char('['),
		char.Char(']'),
		to.String(char.String("Golang")))

	fmt.Println(p([]rune("[Golang]")))

	// Output:
	// Golang [] <nil>
}
