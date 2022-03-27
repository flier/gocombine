package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf16"

	"github.com/flier/gocombine/pkg/stream"
)

var (
	// ErrExpected indicates an expected value.
	ErrExpected = errors.New("expected")

	// ErrUnexpected indicates a unexpected value.
	ErrUnexpected = errors.New("unexpected")
)

type ErrorInfo interface {
	ErrorInfo() string
}

type Error struct {
	*Span
	Err error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

type TokenInfo[T stream.Token] struct {
	Expected *T
	Actual   T
}

func FormatToken[T stream.Token](t T) string {
	switch t := interface{}(t).(type) {
	case byte:
		return fmt.Sprintf("'0x%02x'", t)

	case rune:
		return strconv.QuoteRune(t)

	case uint16:
		return strconv.Quote(string(utf16.Decode([]uint16{t})))

	default:
		return fmt.Sprintf("'%v'", t)
	}
}

func (i *TokenInfo[T]) ErrorInfo() string {
	var b strings.Builder

	if i.Expected != nil {
		fmt.Fprintf(&b, "expected %s, ", FormatToken(*i.Expected))
	}

	fmt.Fprintf(&b, "actual %s", FormatToken(i.Actual))

	return b.String()
}

type RangeInfo[T stream.Token] struct {
	Expected []T
	Actual   []T
}

func FormatRange[T stream.Token](r []T) string {
	switch r := interface{}(r).(type) {
	case []byte:
		var o strings.Builder

		o.WriteRune('[')

		for i, b := range r {
			if i > 0 {
				o.WriteString(", ")
			}

			fmt.Fprintf(&o, "0x%02x", b)
		}

		o.WriteRune(']')

		return o.String()

	case []rune:
		return strconv.Quote(string(r))

	case []uint16:
		return strconv.Quote(string(utf16.Decode(r)))

	default:
		return fmt.Sprintf("'%v'", r)
	}
}

func (i *RangeInfo[T]) ErrorInfo() string {
	var b strings.Builder

	if i.Expected != nil {
		fmt.Fprintf(&b, "expected %s, ", FormatRange(i.Expected))
	} else {
		fmt.Fprintf(&b, "expected end of input, ")
	}

	fmt.Fprintf(&b, "actual %s", FormatRange(i.Actual))

	return b.String()
}

type MessageInfo string

func (i MessageInfo) ErrorInfo() string {
	return string(i)
}

type EOI struct{}

var eoi = EOI{}

func (i EOI) ErrorInfo() string {
	return "end of input"
}

// UnexpectedErr is an error that indicates an unexpected value.
type UnexpectedErr[E ErrorInfo] struct {
	Info E
}

func (e *UnexpectedErr[E]) Err() error {
	return fmt.Errorf("%s, %w", e.Info.ErrorInfo(), ErrUnexpected)
}

func (e *UnexpectedErr[E]) Error() string {
	return e.Err().Error()
}

// Unexpected returns an error that indicates an unexpected value.
func UnexpectedToken[T stream.Token](actual T, x ...T) *UnexpectedErr[*TokenInfo[T]] {
	var expected *T
	if len(x) > 0 {
		expected = &x[0]
	}

	return &UnexpectedErr[*TokenInfo[T]]{&TokenInfo[T]{expected, actual}}
}

// Unexpected returns an error that indicates an unexpected range.
func UnexpectedRange[T stream.Token](expected, actual []T) *UnexpectedErr[*RangeInfo[T]] {
	return &UnexpectedErr[*RangeInfo[T]]{&RangeInfo[T]{expected, actual}}
}

// Unexpected returns an error that indicates an unexpected message.
func UnexpectedMessage(msg string) *UnexpectedErr[MessageInfo] {
	return &UnexpectedErr[MessageInfo]{MessageInfo(msg)}
}

// Unexpected returns an error that indicates an unexpected EOI.
func UnexpectedEOI() *UnexpectedErr[EOI] {
	return &UnexpectedErr[EOI]{eoi}
}

type Span struct {
	Start, End stream.Offset
}

// Spanned marks errors produced inside the `parser` parser with the span from the start of the parse to the end of it.
func Spanned[T stream.Token, O any](parser Func[T, O]) Func[T, O] {
	return func(input []T) (parsed O, remaining []T, err error) {
		parsed, remaining, err = parser(input)

		if err != nil {
			var e *Error

			if errors.As(err, &e) {
				if e.Span == nil || e.Span.Start == e.Span.End {
					e.Span = &Span{stream.Pos(input), stream.Pos(remaining)}
				}
			}
		}

		return
	}
}
