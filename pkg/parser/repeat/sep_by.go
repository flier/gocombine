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
	S stream.Stream[T],
	T stream.Token,
	O, P any,
](
	parser parser.Func[S, T, O],
	separator parser.Func[S, T, P],
) parser.Func[S, T, []O] {
	return combinator.Attempt(choice.Or(SepBy1(parser, separator), token.Value[S]([]O{})))
}

// SepBy1 parses `parser` one or more time separated by `separator`,
// returning a collection with the values from `parser`.
func SepBy1[
	S stream.Stream[T],
	T stream.Token,
	O, P any,
](
	parser parser.Func[S, T, O],
	separator parser.Func[S, T, P],
) parser.Func[S, T, []O] {
	return combinator.Attempt(combinator.Map(
		combinator.Pair(parser, Many(sequence.With(separator, parser))),
		func(p pair.Pair[O, []O]) []O { return append([]O{p.First}, p.Second...) },
	))
}

// SepEndBy parses `parser` zero or more times separated and ended by `separator`,
// returning a collection with the values from `parser`.
func SepEndBy[
	S stream.Stream[T],
	T stream.Token,
	O, P any,
](
	parser parser.Func[S, T, O],
	separator parser.Func[S, T, P],
) parser.Func[S, T, []O] {
	return Many(sequence.Skip(parser, separator))
}

// SepEndBy1 parses `parser` one or more times separated and ended by `separator`,
// returning a collection with the values from `p`.
func SepEndBy1[
	S stream.Stream[T],
	T stream.Token,
	O, P any,
](
	parser parser.Func[S, T, O],
	separator parser.Func[S, T, P],
) parser.Func[S, T, []O] {
	return Many1(sequence.Skip(parser, separator))
}
