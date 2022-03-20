package main

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/to"
)

func main() {
	word := to.String(repeat.Many1(char.Letter[[]rune]()))
	parser := combinator.Map(repeat.SepBy(word, char.Space[[]rune]()), func(words []string) string {
		return words[len(words)-1]
	})

	result, _, err := parser.Parse([]rune("Pick up that word!"))
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
