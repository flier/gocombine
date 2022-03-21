package repeat

import (
	"github.com/flier/gocombine/pkg/pair"
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/sequence"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
)

// SepBy parses `parser` zero or more time separated by `separator`, returning a collection with the values from `p`.
func SepBy[

	T stream.Token,
	O, P any,
](
	parser parser.Func[T, O],
	separator parser.Func[T, P],
) parser.Func[T, []O] {
	return combinator.Attempt(choice.Or(SepBy1(parser, separator), token.Value[T]([]O{})))
}

// SepBy1 parses `parser` one or more time separated by `separator`,
// returning a collection with the values from `parser`.
func SepBy1[

	T stream.Token,
	O, P any,
](
	parser parser.Func[T, O],
	separator parser.Func[T, P],
) parser.Func[T, []O] {
	return combinator.Attempt(combinator.Map(
		combinator.Pair(parser, Many(sequence.With(separator, parser))),
		func(p pair.Pair[O, []O]) []O { return append([]O{p.First}, p.Second...) },
	))
}

// SepEndBy parses `parser` zero or more times separated and ended by `separator`,
// returning a collection with the values from `parser`.
func SepEndBy[

	T stream.Token,
	O, P any,
](
	parser parser.Func[T, O],
	separator parser.Func[T, P],
) parser.Func[T, []O] {
	return Many(sequence.Skip(parser, separator))
}

// SepEndBy1 parses `parser` one or more times separated and ended by `separator`,
// returning a collection with the values from `p`.
func SepEndBy1[

	T stream.Token,
	O, P any,
](
	parser parser.Func[T, O],
	separator parser.Func[T, P],
) parser.Func[T, []O] {
	return Many1(sequence.Skip(parser, separator))
}
