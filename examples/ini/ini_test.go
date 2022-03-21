package ini_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/flier/gocombine/examples/ini"
)

func TestIni(t *testing.T) {
	Convey("Given a INI parser", t, func() {
		p := ini.Parser()

		Convey("When parse string", func() {
			r, remaining, err := p([]rune(`
language=rust

[section]
name=combine; Comment
type=LL(1)

`))

			So(err, ShouldBeNil)
			So(remaining, ShouldBeEmpty)
			So(r, ShouldResemble, &ini.Ini{
				Global: ini.Properties{"language": "rust"},
				Sections: ini.Sections{
					"section": ini.Properties{
						"name": "combine",
						"type": "LL(1)",
					},
				},
			})
		})

		Convey("When parse an incomplete text", func() {
			text := "[error"

			r, remaining, err := p([]rune(text))

			So(r, ShouldBeNil)
			So(string(remaining), ShouldResemble, text)
			So(err, ShouldBeNil)
		})
	})
}
