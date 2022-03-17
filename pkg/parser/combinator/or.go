package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
	"github.com/hashicorp/go-multierror"
)

// Or returns a parser which attempts to parse using `parsers`.
func Or[
	S stream.Stream[T],
	T stream.Token,
	O any,
	P parser.Parser[S, T, O],
](parsers ...P) parser.Func[S, T, O] {
	return func(input S) (out O, remaining S, err error) {
		var errs error

		for _, p := range parsers {
			out, remaining, err = p.Parse(input)
			if err != nil {
				errs = multierror.Append(errs, err)
			} else {
				return
			}
		}

		remaining, err = input, errs
		return
	}
}
