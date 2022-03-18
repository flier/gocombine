package choice

import (
	"github.com/flier/gocombine/pkg/option"
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Optional parses `parser` and outputs `Some(value)` if it succeeds, `None` if it fails without consuming any input.
func Optional[
	S stream.Stream[T],
	T stream.Token,
	O any,
](parser parser.Func[S, T, O]) parser.Func[S, T, option.Option[O]] {
	return func(input S) (parsed option.Option[O], remaining S, err error) {
		var o O
		o, remaining, err = parser(input)
		if err != nil {
			return option.None[O](), input, nil
		}
		return option.Some(o), remaining, nil
	}
}
