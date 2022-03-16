package token_test

import (
	"io"
	"testing"
	"unicode"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/flier/gocombine/pkg/parser/token"
)

func TestToken(t *testing.T) {
	Convey("Given a token parser", t, func() {
		p := token.Token[[]rune]('h')

		Convey("When parse string", func() {
			c, s, err := p([]rune("hello"))

			Convey("Then should get the first character", func() {
				So(err, ShouldBeNil)
				So(c, ShouldEqual, 'h')
				So(string(s), ShouldEqual, "ello")

				Convey("Then parse again should get an error", func() {
					c, s, err = p([]rune(s))

					So(err, ShouldBeError, token.Unexpected('h', 'e'))
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

func TestTokens(t *testing.T) {
	Convey("Given a tokens parser", t, func() {
		p := token.Tokens(func(lhs, rhs byte) bool { return lhs == rhs }, []byte("foo"))

		Convey("When parse string", func() {
			a, s, err := p([]byte("foobar"))

			Convey("Then should get the prefix", func() {
				So(err, ShouldBeNil)
				So(string(a), ShouldEqual, "foo")
				So(string(s), ShouldEqual, "bar")
			})

			Convey("Then parse again should get an error", func() {
				_, _, err = p(s)

				So(err, ShouldBeError, token.Unexpected('f', 'b'))
			})

			Convey("When I parse a empty string", func() {
				_, _, err = p(nil)

				So(err, ShouldBeError, io.ErrUnexpectedEOF)
			})
		})

		p2 := token.Tokens(func(lhs, rhs rune) bool { return unicode.ToLower(lhs) == unicode.ToLower(rhs) }, []rune("foo"))

		Convey("When parse case insensitive string ", func() {
			a, s, err := p2([]rune("foobar"))

			Convey("Then should get the prefix", func() {
				So(err, ShouldBeNil)
				So(string(a), ShouldEqual, "foo")
				So(string(s), ShouldEqual, "bar")
			})
		})
	})
}
