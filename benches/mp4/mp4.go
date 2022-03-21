package mp4

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/bytes"
	"github.com/flier/gocombine/pkg/parser/bytes/be"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/ranges"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
	"github.com/flier/gocombine/pkg/parser/to"
	"github.com/flier/gocombine/pkg/parser/token"
	"github.com/flier/gocombine/pkg/stream"
	"github.com/flier/gocombine/pkg/tuple"
)

type FileType struct {
	MajorBrand       string
	MajorVersion     []byte
	CompatibleBrands []string
}

type Type uint32

const (
	Ftyp Type = iota // Contains the file type, description, and the common data structures used.
	Moov             // Container for all the movie metadata.
	Mdat             // Data container for media.
	Free
	Skip
	Wide
	Unknown
)

type Box struct {
	Type
	*FileType
}

func Mp4[S stream.Stream[byte]]() parser.Func[S, byte, []*Box] {
	name := to.String(ranges.Take[S](4))
	filetype := combinator.Map(
		combinator.Tuple4(
			bytes.Bytes[S]([]byte("ftyp")),
			name,
			ranges.Take[S](4),
			repeat.Many(name),
		),
		func(t tuple.Tuple4[[]byte, string, []byte, []string]) *Box {
			return &Box{
				Type: Ftyp,
				FileType: &FileType{
					MajorBrand:       t.V2,
					MajorVersion:     t.V3,
					CompatibleBrands: t.V4,
				},
			}
		},
	)
	box := sequence.Then(be.Uint32[S](), func(offset uint32) parser.Func[S, byte, []byte] {
		return ranges.Take[S](int(offset) - 4)
	})

	parser := choice.Or(
		filetype,
		combinator.Map(bytes.Bytes[S]([]byte("moov")), func([]byte) *Box { return &Box{Type: Moov} }),
		combinator.Map(bytes.Bytes[S]([]byte("mdat")), func([]byte) *Box { return &Box{Type: Mdat} }),
		combinator.Map(bytes.Bytes[S]([]byte("free")), func([]byte) *Box { return &Box{Type: Free} }),
		combinator.Map(bytes.Bytes[S]([]byte("skip")), func([]byte) *Box { return &Box{Type: Skip} }),
		combinator.Map(bytes.Bytes[S]([]byte("wide")), func([]byte) *Box { return &Box{Type: Wide} }),
		token.Value[S](&Box{Type: Unknown}),
	)

	interpreter := combinator.FlatMap(box, func(b []byte) (box *Box, err error) {
		box, _, err = parser(b)
		return
	})

	return repeat.Many(interpreter)
}