package token_test

import (
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/flier/gocombine/pkg/parser/token"
)

func TestToken(t *testing.T) {
	Convey("Given a token parser", t, func() {
		p := token.Token[[]rune]('h')

		Convey("When I parse string", func() {
			c, s, err := p([]rune("hello"))

			Convey("Then I should get the first character", func() {
				So(err, ShouldBeNil)
				So(c, ShouldEqual, 'h')
				So(string(s), ShouldEqual, "ello")

				Convey("Then parse again should get a error", func() {
					c, s, err = p([]rune(s))

					So(err, ShouldBeError, &token.PeekErr[rune]{'h', 'e'})
					So(string(s), ShouldEqual, "ello")
				})
			})

			Convey("When I parse a empty string", func() {
				_, _, err = p(nil)

				So(err, ShouldBeError, io.ErrUnexpectedEOF)
			})
		})
	})
}
