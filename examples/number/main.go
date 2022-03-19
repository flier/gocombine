package main

import (
	"fmt"
	"strconv"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/ranges"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
)

func main() {
	num := repeat.SkipMany1(char.Digit[[]rune]())
	frac := choice.Optional(sequence.With(char.Char[[]rune]('.'), num))
	parser := combinator.AndThen(
		char.AsString(ranges.Recognize(combinator.Pair(num, frac))),
		func(s string) (float64, error) { return strconv.ParseFloat(s, 64) },
	)

	result, remaining, err := parser.Parse([]rune("1234.45"))
	if err != nil {
		panic(err)
	}
	if len(remaining) > 0 {
		panic(remaining)
	}

	fmt.Println(result)
}
