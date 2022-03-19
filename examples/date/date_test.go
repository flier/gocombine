package date_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/flier/gocombine/examples/date"
)

var cases = map[string]date.DateTime{
	"2015-08-02T18:54:42+02": {
		Date: date.Date{Year: 2015, Month: 8, Day: 2},
		Time: date.Time{Hour: 18, Minute: 54, Second: 42, TimeZone: 120},
	},
	"50015-12-30T08:54:42Z": {
		Date: date.Date{Year: 50015, Month: 12, Day: 30},
		Time: date.Time{Hour: 8, Minute: 54, Second: 42, TimeZone: 0},
	},
}

func TestDate(t *testing.T) {
	Convey("Given a date parser", t, func() {
		p := date.Parser[[]rune]()

		for s, dt := range cases {
			Convey("When parse a date: "+s, func() {
				r, remaining, err := p([]rune(s))

				So(err, ShouldBeNil)
				So(remaining, ShouldBeEmpty)
				So(r, ShouldResemble, dt)
			})
		}
	})
}
