package choice

import (
	"github.com/hashicorp/go-multierror"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/stream"
)

// Or returns a parser which attempts to parse using `parsers`.
func Or[
	S stream.Stream[T],
	T stream.Token,
	O any,
](
	parsers ...parser.Func[S, T, O],
) parser.Func[S, T, O] {
	return combinator.Attempt(func(input S) (out O, remaining S, err error) {
		var errs error

		for _, p := range parsers {
			if out, remaining, err = p(input); err != nil {
				errs = multierror.Append(errs, err)
			} else {
				return
			}
		}

		err = errs

		return
	})
}
