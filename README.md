# gocombine [![Continuous integration](https://github.com/flier/gocombine/actions/workflows/ci.yml/badge.svg)](https://github.com/flier/gocombine/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/flier/gocombine/gocombine.svg)](https://pkg.go.dev/github.com/flier/gocombine) [![Apache](https://img.shields.io/badge/license-Apache-blue.svg)](https://github.com/flier/gohs/blob/master/LICENSE-APACHE) [![MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/flier/gohs/blob/master/LICENSE-MIT)


An experimental implementation of parser combinators for `Golang[Generic]`, inspired by the Haskell library [Parsec][] and the Rust [combine][]. As in Parsec the parsers are [LL(1)][] by default but they can opt-in to arbitrary lookahead using the attempt combinator.

## Examples

```go
package main

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/to"
	"github.com/flier/gocombine/pkg/stream"
)

func lastWord[S stream.Stream[rune]]() parser.Func[S, rune, string] {
	word := to.String(repeat.Many1(char.Letter[S]()))
	return combinator.Map(
		repeat.SepBy(word, char.Space[S]()),
		func(words []string) string {
			return words[len(words)-1]
		})
}

func main() {
	parser := lastWord[[]rune]()

	result, _, err := parser.Parse([]rune("Pick up that word!"))
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

```

Larger examples can be found in the [examples](examples) and [benches](benches) folders.

## About

A parser combinator is, broadly speaking, a function which takes several parsers as arguments and returns a new parser, created by combining those parsers. For instance, the [Many][] parser takes one parser, `parser`, as input and returns a new parser which applies `parser` zero or more times. Thanks to the modularity that parser combinators gives it is possible to define parsers for a wide range of tasks without needing to implement the low level plumbing while still having the full power of Rust when you need it.

The library adheres to [semantic versioning][].

If you end up trying it I welcome any feedback from your experience with it. I am usually reachable within a day by opening an issue, sending an email or posting a message on Gitter.

## License

This project is licensed under either of Apache License ([LICENSE-APACHE](LICENSE-APACHE)) or MIT license ([LICENSE-MIT](LICENSE-MIT)) at your option.

[combine]:https://github.com/Marwes/combine
[LL(1)]:https://en.wikipedia.org/wiki/LL_parser
[Many]:https://pkg.go.dev/github.com/flier/gocombine/pkg/repeat#Many
[Parsec]:https://hackage.haskell.org/package/parsec
[semantic versioning]:https://semver.org/
