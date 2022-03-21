package choice

import (
	"github.com/flier/gocombine/pkg/option"
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
)

// Optional parses `parser` and outputs `Some(value)` if it succeeds, `None` if it fails without consuming any input.
func Optional[T stream.Token, O any](parser parser.Func[T, O]) parser.Func[T, option.Option[O]] {
	return func(input []T) (parsed option.Option[O], remaining []T, err error) {
		var o O

		if o, remaining, err = parser(input); err != nil {
			return option.None[O](), input, nil
		}

		return option.Some(o), remaining, nil
	}
}
