package http

import (
	"net/http"
	"net/url"

	"github.com/flier/gocombine/pkg/pair"
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/bytes"
	"github.com/flier/gocombine/pkg/parser/bytes/num"
	"github.com/flier/gocombine/pkg/parser/char"
	"github.com/flier/gocombine/pkg/parser/choice"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/ranges"
	"github.com/flier/gocombine/pkg/parser/repeat"
	"github.com/flier/gocombine/pkg/parser/sequence"
	"github.com/flier/gocombine/pkg/stream"
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

func EoL[S stream.Stream[byte]]() parser.Func[S, byte, byte] {
	return choice.Or(
		sequence.Skip(bytes.Byte[S]('\r'), bytes.Byte[S]('\n')),
		bytes.Byte[S]('\n'))
}

func MessageHeader[S stream.Stream[byte]]() parser.Func[S, byte, http.Header] {
	header := sequence.Skip(sequence.With(
		ranges.TakeWhile1[S](IsHozizontalSpace),
		char.AsString(ranges.TakeWhile1[S](func(b byte) bool { return b != '\r' && b != '\n' }))),
		EoL[S]())
	name := char.AsString(ranges.TakeWhile1[S](IsToken))
	sep := bytes.Byte[S](':')
	value := repeat.Many1(header)

	return combinator.Fold(
		repeat.Many(combinator.Tuple3(name, sep, value)),
		func() http.Header { return make(http.Header) },
		func(header http.Header, t tuple.Tuple3[string, byte, []string]) {
			header[t.V1] = t.V3
		},
	).Message("while parsing header")
}

func Parser[S stream.Stream[byte]]() parser.Func[S, byte, *http.Request] {
	method := char.AsString(ranges.TakeWhile1[S](IsToken))
	uri := char.AsString(ranges.TakeWhile1[S](IsNotSpace))
	ver := num.Atoi(repeat.Many1(bytes.Digit[S]()))
	version := sequence.With(bytes.Bytes[S]([]byte("HTTP/")),
		combinator.Tuple3(ver, bytes.Byte[S]('.'), ver))

	requestLine := combinator.AndThen(
		combinator.Tuple3(
			sequence.Skip(method, bytes.Spaces[S]()),
			sequence.Skip(uri, bytes.Spaces[S]()),
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
			sequence.Skip(requestLine, EoL[S]()),
			sequence.Skip(MessageHeader[S](), EoL[S]()),
		),
		func(p pair.Pair[*http.Request, http.Header]) *http.Request {
			req := p.First
			req.Header = p.Second
			return req
		},
	)
}
