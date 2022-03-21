package char_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/to"
)

func TestFloat(t *testing.T) {
	Convey("Given a number parser", t, func() {
		p := to.Float(char.Float())

		Convey("When parse a zero", func() {
			r, s, err := p([]rune("0"))
			So(err, ShouldBeNil)
			So(s, ShouldBeEmpty)
			So(r, ShouldEqual, 0.0)
		})

		Convey("When parse a integer", func() {
			r, s, err := p([]rune("123"))
			So(err, ShouldBeNil)
			So(s, ShouldBeEmpty)
			So(r, ShouldEqual, 123)
		})

		Convey("When parse a number", func() {
			r, s, err := p([]rune("3.1415926"))
			So(err, ShouldBeNil)
			So(s, ShouldBeEmpty)
			So(r, ShouldEqual, 3.1415926)
		})

		Convey("When parse a number with exponent", func() {
			r, s, err := p([]rune("31415926e-7"))
			So(err, ShouldBeNil)
			So(s, ShouldBeEmpty)
			So(r, ShouldEqual, 3.1415926)
		})
	})
}
