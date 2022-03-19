package main

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
)

func main() {
	word := char.AsString(repeat.Many1(char.Letter[[]rune]()))
	parser := combinator.Map(repeat.SepBy(word, char.Space[[]rune]()), func(words []string) string {
		return words[len(words)-1]
	})

	result, remaining, err := parser.Parse([]rune("Pick up that word!"))
	if err != nil {
		panic(err)
	}
	if len(remaining) > 0 {
		panic(remaining)
	}

	fmt.Println(result)
}
