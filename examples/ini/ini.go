package ini

import (
	"github.com/flier/gocombine/pkg/pair"
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
	"github.com/flier/gocombine/pkg/parser/to"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/tuple"
)

type (
	Properties map[string]string
	Sections   map[string]Properties
)

type Ini struct {
	Global   Properties
	Sections Sections
}

func property() parser.Func[rune, []string] {
	key := to.String(repeat.Many1(token.Satisfy(func(c rune) bool {
		return c != '=' && c != '[' && c != ';'
	}))).Expected("key")

	assign := token.Token('=')

	value := to.String(repeat.Many1(token.Satisfy(func(c rune) bool {
		return c != '\n' && c != ';'
	}))).Expected("value")

	return choice.And(sequence.Skip(key, assign), value).Message("while parsing property")
}

func whitespace() parser.Func[rune, any] {
	comment := sequence.With(token.Token(';'),
		repeat.SkipMany(token.Satisfy(func(c rune) bool {
			return c != '\n'
		}))).Message("while parsing comment")

	return repeat.SkipMany(choice.Or(
		repeat.SkipMany1(char.Space()),
		comment,
	)).Message("while parsing whitespace")
}

func properties() parser.Func[rune, Properties] {
	// After each property we skip any whitespace that followed it
	return combinator.Fold(
		repeat.Many(sequence.Skip(property(), whitespace())),
		func() Properties { return make(Properties) },
		func(m Properties, s []string) { m[s[0]] = s[1] },
	).Message("while parsing properties")
}

func sections() parser.Func[rune, Sections] {
	name := sequence.Between(
		token.Token('['),
		token.Token(']'),
		to.String(repeat.Many(token.Satisfy(func(c rune) bool { return c != ']' }))),
	).Message("while parsing name")

	return combinator.Fold(
		repeat.Many(combinator.Pair(sequence.Skip(name, whitespace()), properties())),
		func() Sections { return make(Sections) },
		func(s Sections, p pair.Pair[string, Properties]) { s[p.First] = p.Second },
	).Message("while parsing sections")
}

func Parser() parser.Func[rune, *Ini] {
	return combinator.Map(
		combinator.Tuple3(whitespace(), properties(), sections()),
		func(s tuple.Tuple3[any, Properties, Sections]) *Ini {
			if len(s.V2) == 0 && len(s.V3) == 0 {
				return nil
			}

			return &Ini{
				Global:   s.V2,
				Sections: s.V3,
			}
		},
	)
}
