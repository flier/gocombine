// Parser example for INI files.

package ini

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println(filepath.Base(os.Args[0]), "<ini files>...")
		return
	}

	p := Parser[[]rune]()

	for _, filename := range flag.Args() {
		b, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalln("open file", filename, err)
		}

		s := string(b)

		r, remaining, err := p([]rune(s))
		if err != nil {
			log.Fatalln("parse file", filename, err)
		}
		if len(remaining) > 0 {
			log.Println("unexpected", string(remaining))
		}
		log.Printf("parsed: %#+v", r)
	}
}
