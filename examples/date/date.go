package date

import (
	"github.com/flier/gocombine/pkg/option"
	"github.com/flier/gocombine/pkg/pair"
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
	"github.com/flier/gocombine/pkg/parser/to"
	"github.com/flier/gocombine/pkg/tuple"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

type Time struct {
	Hour     int
	Minute   int
	Second   int
	TimeZone int
}

type DateTime struct {
	Date
	Time
}

func two_digit() parser.Func[rune, int] {
	digit := to.Int(char.Digit())

	return combinator.Map(
		combinator.Pair(digit, digit),
		func(p pair.Pair[int, int]) int {
			return p.First*10 + p.Second
		},
	)
}

// Parses a time zone
// +0012
// -06:30
// -01
// Z
func time_zone() parser.Func[rune, int] {
	utc := combinator.Map(char.Char('Z'), func(r rune) int { return 0 })
	offset := combinator.Map(
		combinator.Tuple3(
			choice.Or(char.Char('+'), char.Char('-')),
			two_digit(),
			choice.Optional(sequence.With(choice.Optional(char.Char(':')), two_digit())),
		),
		func(t tuple.Tuple3[rune, int, option.Option[int]]) (offset int) {
			offset = t.V2*60 + t.V3.UnwrapOrDefault()
			if t.V1 == '-' {
				offset = -offset
			}
			return
		},
	)

	return choice.Or(utc, offset)
}

// Parses a date
// 2010-01-30
func date() parser.Func[rune, Date] {
	year := to.Int(repeat.Many1(char.Digit()))
	month := two_digit()
	day := two_digit()
	sep := char.Char('-')

	return combinator.Map(
		combinator.Tuple3(
			sequence.Skip(year, sep),
			sequence.Skip(month, sep),
			day,
		),
		func(t tuple.Tuple3[int, int, int]) Date {
			return Date{t.V1, t.V2, t.V3}
		},
	)
}

// Parses a time
// 12:30:02
func time() parser.Func[rune, Time] {
	hour := two_digit()
	minute := two_digit()
	second := two_digit()
	sep := char.Char(':')

	return combinator.Map(
		combinator.Tuple4(
			sequence.Skip(hour, sep),
			sequence.Skip(minute, sep),
			second,
			time_zone(),
		),
		func(t tuple.Tuple4[int, int, int, int]) Time {
			return Time{t.V1, t.V2, t.V3, t.V4}
		},
	)
}

// Parses a date time according to ISO8601
// 2015-08-02T18:54:42+02
func Parser() parser.Func[rune, DateTime] {
	return combinator.Map(
		combinator.Tuple3(
			date(), char.Char('T'), time(),
		),
		func(t tuple.Tuple3[Date, rune, Time]) DateTime {
			return DateTime{t.V1, t.V3}
		},
	)
}
