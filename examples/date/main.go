// Parser example for ISO8601 dates. This does not handle the entire specification but it should
// show the gist of it and be easy to extend to parse additional forms.

package date

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println(filepath.Base(os.Args[0]), "<date time>...")
		return
	}

	p := Parser[[]rune]()

	for _, arg := range flag.Args() {
		dt, _, err := p([]rune(arg))
		if err != nil {
			panic(err)
		}

		fmt.Printf("%#+v\n", dt)
	}
}
