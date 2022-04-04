package token_test

import (
	"fmt"
	"io"
	"testing"
	"unicode"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/token"
)

func TestToken(t *testing.T) {
	Convey("Given a token parser", t, func() {
		p := token.Token('h')

		Convey("When parse string", func() {
			c, s, err := p([]rune("hello"))

			Convey("Then should get the first character", func() {
				So(err, ShouldBeNil)
				So(c, ShouldEqual, 'h')
				So(string(s), ShouldEqual, "ello")

				Convey("Then parse again should get an error", func() {
					c, s, err = p([]rune(s))

					So(err, ShouldBeError, parser.UnexpectedToken('e', 'h'))
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
		p := token.Tokens(func(lhs, rhs byte) bool { return lhs == rhs }, []byte("foo"), []byte("foo"))

		Convey("When parse string", func() {
			a, s, err := p([]byte("foobar"))

			Convey("Then should get the prefix", func() {
				So(err, ShouldBeNil)
				So(string(a), ShouldEqual, "foo")
				So(string(s), ShouldEqual, "bar")
			})

			Convey("Then parse again should get an error", func() {
				_, _, err = p(s)

				So(err, ShouldBeError, parser.UnexpectedRange([]byte("foo"), []byte("b")))
			})

			Convey("When I parse a empty string", func() {
				_, _, err = p(nil)

				So(err, ShouldBeError, io.ErrUnexpectedEOF)
			})
		})

		p2 := token.Tokens(func(lhs, rhs rune) bool {
			return unicode.ToLower(lhs) == unicode.ToLower(rhs)
		}, []rune("[fF][oO][oO]"), []rune("foo"))

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

func ExampleAny() {
	p := token.Any[rune]()

	fmt.Println(p([]rune("apple")))
	fmt.Println(p(nil))

	// Output:
	// 97 [112 112 108 101] <nil>
	// 0 [] unexpected EOF
}

func ExampleToken() {
	p := token.Token('a')

	fmt.Println(p([]rune("apple")))
	fmt.Println(p([]rune("foobar")))

	// Output:
	// 97 [112 112 108 101] <nil>
	// 102 [102 111 111 98 97 114] expected 'a', actual 'f', unexpected
}

func ExampleTokens() {
	p := token.Tokens(func(l, r rune) bool {
		return unicode.ToLower(l) == unicode.ToLower(r)
	}, []rune("foo"), []rune("foo"))

	fmt.Println(p([]rune("apple")))
	fmt.Println(p([]rune("FooBar")))

	// Output:
	// [97] [97 112 112 108 101] expected "foo", actual "a", unexpected
	// [70 111 111] [66 97 114] <nil>
}

func ExampleOneOf() {
	p := token.OneOf([]rune("abc"))

	fmt.Println(p([]rune("apple")))
	fmt.Println(p([]rune("foobar")))

	// Output:
	// 97 [112 112 108 101] <nil>
	// 102 [102 111 111 98 97 114] one of "abc", satisfy, actual 'f', unexpected
}

func ExampleNoneOf() {
	p := token.NoneOf([]rune("abc"))

	fmt.Println(p([]rune("apple")))
	fmt.Println(p([]rune("foobar")))

	// Output:
	// 97 [97 112 112 108 101] none of "abc", satisfy, actual 'a', unexpected
	// 102 [111 111 98 97 114] <nil>
}

func ExampleEOF() {
	p := token.EOF[rune]()

	fmt.Println(p(nil))
	fmt.Println(p([]rune("foobar")))

	// Output:
	// true [] <nil>
	// false [102 111 111 98 97 114] eof, expected end of input, actual "foobar", unexpected
}
