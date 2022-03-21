package repeat

import (
	"github.com/flier/gocombine/pkg/pair"
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// ChainL1 parses `parser` one or more times separated by `op`.
// The value returned is the one produced
// by the left associative application of the function returned by the parser `op`.
func ChainL1[

	T stream.Token,
	O any,
](
	parser parser.Func[T, O],
	op parser.Func[T, func(l, h O) O],
) parser.Func[T, O] {
	return combinator.Map(
		combinator.Pair(parser, Many1(combinator.Pair(op, parser))),
		func(p pair.Pair[O, []pair.Pair[func(l, h O) O, O]]) (acc O) {
			acc = p.First
			for _, v := range p.Second {
				acc = v.First(acc, v.Second)
			}

			return
		},
	)
}

// ChainR1 parses `p` one or more times separated by `op`.
// The value returned is the one produced
// by the right associative application of the function returned by `op`.
func ChainR1[

	T stream.Token,
	O any,
](
	parser parser.Func[T, O],
	op parser.Func[T, func(l, h O) O],
) parser.Func[T, O] {
	return combinator.Map(
		combinator.Pair(Many1(combinator.Pair(parser, op)), parser),
		func(p pair.Pair[[]pair.Pair[O, func(l, h O) O], O]) (acc O) {
			acc = p.Second

			n := len(p.First)
			for i := 0; i < n; i++ {
				acc = p.First[n-1-i].Second(p.First[n-1-i].First, acc)
			}

			return
		},
	)
}
