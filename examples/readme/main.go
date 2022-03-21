package main

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/to"
	"github.com/flier/gocombine/pkg/stream"
)

func lastWord[S stream.Stream[rune]]() parser.Func[S, rune, string] {
	word := to.String(repeat.Many1(char.Letter[S]()))
	return combinator.Map(
		repeat.SepBy(word, char.Space[S]()),
		func(words []string) string {
			return words[len(words)-1]
		})
}

func main() {
	parser := lastWord[[]rune]()

	result, _, err := parser.Parse([]rune("Pick up that word!"))
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
