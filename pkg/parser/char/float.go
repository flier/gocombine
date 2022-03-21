package char

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/ranges"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
)

// Float parses a floating-point numbers.
func Float[S stream.Stream[rune]]() parser.Func[S, rune, []rune] {
	sign := token.OneOf[S]([]rune("+-"))
	integer := repeat.Many1(Digit[S]())
	frac := combinator.Pair(Char[S]('.'), repeat.Many(Digit[S]()))
	e := token.Satisfy[S](func(r rune) bool { return r == 'e' || r == 'E' })
	exp := combinator.Tuple3(e, choice.Optional(sign), integer)

	return ranges.Recognize(combinator.Pair(
		choice.Optional(sign),
		choice.Or(
			StringFold[S]("nan"),
			StringFold[S]("inf"),
			ranges.Recognize(combinator.Tuple3(
				integer,
				choice.Optional(frac),
				choice.Optional(exp),
			)),
		),
	))
}
