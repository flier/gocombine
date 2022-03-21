package main

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/to"
)

func main() {
	parser := combinator.Map(
		repeat.SepBy(to.String(repeat.Many1(char.Letter())), char.Space()),
		func(words []string) string {
			return words[len(words)-1]
		})

	result, _, err := parser.Parse([]rune("Pick up that word!"))
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
