package char

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/ranges"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/token"
)

// Float parses a floating-point numbers.
func Float() parser.Func[rune, []rune] {
	sign := OneOf("+-")
	integer := repeat.Many1(Digit())
	frac := combinator.Pair(Char('.'), repeat.Many(Digit()))
	e := token.Satisfy(func(r rune) bool { return r == 'e' || r == 'E' })
	exp := combinator.Tuple3(e, choice.Optional(sign), integer)

	return ranges.Recognize(combinator.Pair(
		choice.Optional(sign),
		choice.Or(
			Fold("nan"),
			Fold("inf"),
			ranges.Recognize(combinator.Tuple3(
				integer,
				choice.Optional(frac),
				choice.Optional(exp),
			)),
		),
	))
}
