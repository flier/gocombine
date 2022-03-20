package json_test

import (
	_ "embed"
	stdjson "encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/flier/gocombine/benches/json"
)

//go:embed testdata/data.json
var data string

func BenchmarkJSON(b *testing.B) {
	p := json.Parser[[]rune]()

	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, remaining, err := p([]rune(data)); err != nil {
			b.Fatal(err, string(remaining))
		}
	}
}

func BenchmarkStdJSON(b *testing.B) {
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var a []interface{}

		if err := stdjson.Unmarshal([]byte(data), &a); err != nil {
			b.Fatal(err)
		}
	}
}

func TestNumber(t *testing.T) {
	Convey("Given a number parser", t, func() {
		p := json.JsonNumber[[]rune]()

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

func TestString(t *testing.T) {
	Convey("Given a string parser", t, func() {
		p := json.JsonString[[]rune]()

		Convey("When parse a string", func() {
			r, s, err := p([]rune(`"hello"`))
			So(err, ShouldBeNil)
			So(s, ShouldBeEmpty)
			So(r, ShouldEqual, "hello")
		})

		Convey("When parse a escaped string", func() {
			r, s, err := p([]rune(`"hello\t\"world\"\n"`))
			So(err, ShouldBeNil)
			So(s, ShouldBeEmpty)
			So(r, ShouldEqual, "hello\t\"world\"\n")
		})
	})
}

func TestObject(t *testing.T) {
	Convey("Given a object parser", t, func() {
		p := json.JsonObject[[]rune]()

		Convey("When parsing a empty object", func() {
			r, s, err := p([]rune(`{}`))
			So(err, ShouldBeNil)
			So(s, ShouldBeEmpty)
			So(r, ShouldResemble, map[string]*json.Value{})
		})

		Convey("When parse a object", func() {
			r, s, err := p([]rune(`{
    "array": [1, ""],
    "object": {},
    "number": 3.14,
    "small_number": 0.59,
    "int": -100,
    "exp": -1e2,
    "exp_neg": 23e-2,
    "true": true,
    "false"  : false,
    "null" : null
}`))
			So(err, ShouldBeNil)
			So(s, ShouldBeEmpty)
			So(r, ShouldResemble, map[string]*json.Value{
				"array":        json.Array(json.Number(1), json.String("")),
				"object":       json.Object(map[string]*json.Value{}),
				"number":       json.Number(3.14),
				"small_number": json.Number(0.59),
				"int":          json.Number(-100),
				"exp":          json.Number(-1e2),
				"exp_neg":      json.Number(23e-2),
				"true":         json.True,
				"false":        json.False,
				"null":         json.Null,
			})
		})
	})
}
