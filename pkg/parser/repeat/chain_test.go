package repeat_test

import (
	"fmt"
	"math"
	"strconv"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
)

func ExampleChainL1() {
	number := combinator.AndThen(char.AsString(char.Digit[[]rune]()), strconv.Atoi)
	sub := combinator.Map(char.Char[[]rune]('-'), func(rune) func(l, r int) int {
		return func(l, r int) int { return l - r }
	})
	p := repeat.ChainL1(number, sub)

	fmt.Println(p([]rune("9-3-5")))

	// Output:
	// 1 [] <nil>
}

func ExampleChainR1() {
	number := combinator.AndThen(char.AsString(char.Digit[[]rune]()), strconv.Atoi)
	sub := combinator.Map(char.Char[[]rune]('^'), func(rune) func(l, r int) int {
		return func(l, r int) int { return int(math.Pow(float64(l), float64(r))) }
	})
	p := repeat.ChainR1(number, sub)

	fmt.Println(p([]rune("2^3^2")))

	// Output:
	// 512 [] <nil>
}
