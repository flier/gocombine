package json

import (
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

type Value struct {
	Number *float64
	String *string
	Bool   *bool
	Object map[string]*Value
	Array  []*Value
}

var (
	True  = Bool(true)
	False = Bool(false)
	Null  = (*Value)(nil)
)

func Number(f float64) *Value           { return &Value{Number: &f} }
func String(s string) *Value            { return &Value{String: &s} }
func Bool(b bool) *Value                { return &Value{Bool: &b} }
func Object(m map[string]*Value) *Value { return &Value{Object: m} }
func Array(a []*Value) *Value           { return &Value{Array: a} }
func ArrayOf(a ...*Value) *Value        { return &Value{Array: a} }

func lex[O any](p parser.Func[rune, O]) parser.Func[rune, O] {
	return sequence.Skip(p, char.Spaces())
}

func JsonNumber() parser.Func[rune, float64] {
	// integer := lex(to.Int(repeat.Many1(char.Digit())))

	// i := choice.Or(
	// 	combinator.Map(char.Char('0'), func(r rune) float64 { return 0.0 }),
	// 	combinator.Map(integer, func(i int) float64 { return float64(i) }))

	// frac := combinator.Map(
	// 	sequence.With(choice.Optional(char.Char('.')), repeat.Many(char.Digit())),
	// 	func(s []rune) (acc float64) {
	// 		magnitude := 1.0

	// 		for _, c := range s {
	// 			magnitude /= 10.0
	// 			acc += float64(c-'0') * magnitude
	// 		}

	// 		return
	// 	},
	// )

	// exp := combinator.Map(
	// 	combinator.Pair(
	// 		sequence.With(
	// 			token.Satisfy(func(r rune) bool { return r == 'e' || r == 'E' }),
	// 			choice.Optional(char.Char('-'))),
	// 		integer),
	// 	func(p pair.Pair[option.Option[rune], int]) int {
	// 		if p.First.HasSome() {
	// 			return -p.Second
	// 		}

	// 		return p.Second
	// 	},
	// )

	// i = combinator.Map(
	// 	combinator.Pair(choice.Optional(char.Char('-')), i),
	// 	func(p pair.Pair[option.Option[rune], float64]) float64 {
	// 		if p.First.HasSome() {
	// 			return -p.Second
	// 		}

	// 		return p.Second
	// 	})

	// float := combinator.Map(
	// 	combinator.Pair(i, sequence.With(choice.Optional(char.Char('.')), frac)),
	// 	func(p pair.Pair[float64, float64]) float64 {
	// 		if p.First >= 0.0 {
	// 			return p.First + p.Second
	// 		}

	// 		return p.First - p.Second
	// 	})

	// return lex(combinator.Map(
	// 	combinator.Pair(float, choice.Optional(exp)),
	// 	func(p pair.Pair[float64, option.Option[int]]) float64 {
	// 		if p.Second.HasSome() {
	// 			return p.First * math.Pow10(p.Second.Unwrap())
	// 		}

	// 		return p.First
	// 	},
	// )).Expected("number")

	return lex(to.Float(char.Float()))
}

func JsonString() parser.Func[rune, string] {
	c := choice.Or(
		sequence.With(
			token.Token('\\'),
			token.OneOf([]rune(`"\/bfnrt`)).Map(
				func(c rune) rune {
					switch c {
					case 'b':
						return '\b'
					case 'f':
						return '\f'
					case 'n':
						return '\n'
					case 'r':
						return '\r'
					case 't':
						return '\t'
					default:
						return c
					}
				}),
		),
		token.NoneOf([]rune(`"`)),
	)

	return to.String(sequence.Between(
		char.Char('"'),
		lex(char.Char('"')),
		repeat.Many(c),
	)).Expected("string")
}

func JsonObject() parser.Func[rune, map[string]*Value] {
	field := combinator.Tuple3(
		to.String(JsonString()),
		lex(char.Char(':')),
		combinator.Lazy(value),
	).Expected("field")

	fields := combinator.Map(
		repeat.SepBy(field, lex(char.Char(','))),
		func(t []tuple.Tuple3[string, rune, *Value]) (m map[string]*Value) {
			m = make(map[string]*Value)

			for _, f := range t {
				m[f.V1] = f.V3
			}

			return
		},
	)

	return sequence.Between(
		lex(char.Char('{')),
		lex(char.Char('}')),
		fields,
	).Expected("object")
}

func JsonArray() parser.Func[rune, []*Value] {
	return sequence.Between(
		lex(char.Char('[')),
		lex(char.Char(']')),
		repeat.SepBy(
			combinator.Lazy(value),
			lex(char.Char(','))),
	).Expected("array")
}

func Parser() parser.Func[rune, *Value] {
	return value()
}

func value() parser.Func[rune, *Value] {
	return choice.Or(
		combinator.Map(JsonString(), String),
		combinator.Map(JsonObject(), Object),
		combinator.Map(JsonArray(), Array),
		combinator.Map(JsonNumber(), Number),
		combinator.Map(lex(char.String("false")), func(s []rune) *Value { return False }),
		combinator.Map(lex(char.String("true")), func(s []rune) *Value { return True }),
		combinator.Map(lex(char.String("null")), func(s []rune) *Value { return nil }),
	)
}
