package http

import (
	"net/http"
	"net/url"

	"github.com/flier/gocombine/pkg/pair"
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/bytes"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/ranges"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
	"github.com/flier/gocombine/pkg/parser/to"
	"github.com/flier/gocombine/pkg/tuple"
)

type (
	Request = http.Request
	Header  = http.Header
)

func IsToken(b byte) bool {
	if ((b & 0x80) == 0x80) || ((b & 0xE0) == 0) {
		return false
	}

	switch b {
	case '(', ')', '<', '>', '@', ',', ';', ':', '\\', '"', '/', '[', ']', '?', '=', '{', '}', ' ':
		return false
	default:
		return true
	}
}

func IsHozizontalSpace(b byte) bool { return b == ' ' || b == '\t' }
func IsNotSpace(b byte) bool        { return b != ' ' }

func EOL() parser.Func[byte, byte] {
	return choice.Or(
		sequence.Skip(bytes.Byte('\r'), bytes.Byte('\n')),
		bytes.Byte('\n'))
}

func MessageHeader() parser.Func[byte, http.Header] {
	header := sequence.Skip(sequence.With(
		ranges.TakeWhile1(IsHozizontalSpace),
		to.String(ranges.TakeWhile1(func(b byte) bool { return b != '\r' && b != '\n' }))),
		EOL())
	name := to.String(ranges.TakeWhile1(IsToken))
	sep := bytes.Byte(':')
	value := repeat.Many1(header)

	return combinator.Fold(
		repeat.Many(combinator.Tuple3(name, sep, value)),
		func() http.Header { return make(http.Header) },
		func(header http.Header, t tuple.Tuple3[string, byte, []string]) http.Header {
			header[t.V1] = t.V3
			return header
		},
	).Message("while parsing header")
}

func Parser() parser.Func[byte, *http.Request] {
	method := to.String(ranges.TakeWhile1(IsToken))
	uri := to.String(ranges.TakeWhile1(IsNotSpace))
	ver := to.Int(repeat.Many1(bytes.Digit()))
	version := sequence.With(bytes.Bytes([]byte("HTTP/")),
		combinator.Tuple3(ver, bytes.Byte('.'), ver))

	requestLine := combinator.AndThen(
		combinator.Tuple3(
			sequence.Skip(method, bytes.Spaces()),
			sequence.Skip(uri, bytes.Spaces()),
			version,
		),
		func(t tuple.Tuple3[string, string, tuple.Tuple3[int, byte, int]]) (req *http.Request, err error) {
			var uri *url.URL
			uri, err = url.ParseRequestURI(t.V2)
			if err != nil {
				return
			}

			req = &http.Request{
				Method:     t.V1,
				URL:        uri,
				ProtoMajor: t.V3.V1,
				ProtoMinor: t.V3.V3,
			}

			return
		},
	).Message("while parsing request line")

	return combinator.Map(
		combinator.Pair(
			sequence.Skip(requestLine, EOL()),
			sequence.Skip(MessageHeader(), EOL()),
		),
		func(p pair.Pair[*http.Request, http.Header]) *http.Request {
			req := p.First
			req.Header = p.Second

			return req
		},
	)
}
