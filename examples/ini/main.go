package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/flier/gocombine/pkg/pair"
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
	"github.com/flier/gocombine/pkg/tuple"
)

type Properties map[string]string
type Sections map[string]Properties

type Ini struct {
	Global   Properties
	Sections Sections
}

func property[S stream.Stream[rune]]() parser.Func[S, rune, []string] {
	key := char.AsString(repeat.Many1(token.Satisfy[S](func(c rune) bool {
		return c != '=' && c != '[' && c != ';'
	}))).Expected("key")

	assign := token.Token[S]('=')

	value := char.AsString(repeat.Many1(token.Satisfy[S](func(c rune) bool {
		return c != '\n' && c != ';'
	}))).Expected("value")

	return choice.And(sequence.Skip(key, assign), value).Message("while parsing property")
}

func whitespace[S stream.Stream[rune]]() parser.Func[S, rune, any] {
	comment := sequence.With(token.Token[S](';'),
		repeat.SkipMany(token.Satisfy[S](func(c rune) bool { return c != '\n' }))).Message("while parsing comment")

	return repeat.SkipMany(choice.Or(repeat.SkipMany1(char.Space[S]()), comment)).Message("while parsing whitespace")
}

func properties[S stream.Stream[rune]]() parser.Func[S, rune, Properties] {
	// After each property we skip any whitespace that followed it
	return combinator.Fold(
		repeat.Many(sequence.Skip(property[S](), whitespace[S]())),
		func() Properties { return make(Properties) },
		func(m Properties, s []string) { m[s[0]] = s[1] },
	).Message("while parsing properties")
}

func sections[S stream.Stream[rune]]() parser.Func[S, rune, Sections] {
	name := sequence.Between(
		token.Token[S]('['),
		token.Token[S](']'),
		char.AsString(repeat.Many(token.Satisfy[S](func(c rune) bool { return c != ']' }))),
	).Message("while parsing name")

	return combinator.Fold(
		repeat.Many(combinator.Pair(sequence.Skip(name, whitespace[S]()), properties[S]())),
		func() Sections { return make(Sections) },
		func(s Sections, p pair.Pair[string, Properties]) { s[p.First] = p.Second },
	).Message("while parsing sections")
}

func ini[S stream.Stream[rune]]() parser.Func[S, rune, *Ini] {
	return combinator.Map(
		combinator.Tuple3(whitespace[S](), properties[S](), sections[S]()),
		func(s tuple.Tuple3[any, Properties, Sections]) *Ini {
			return &Ini{
				Global:   s.I2,
				Sections: s.I3,
			}
		},
	)
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println(filepath.Base(os.Args[0]), "[ini files]...")
		return
	}

	p := ini[[]rune]()

	for _, filename := range flag.Args() {
		b, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalln("open file", filename, err)
		}

		s := string(b)

		r, remaining, err := p([]rune(s))
		if err != nil {
			log.Fatalln("parse file", filename, err)
		}
		if len(remaining) > 0 {
			log.Println("unexpected", string(remaining))
		}
		log.Printf("parsed: %#+v", r)
	}
}
