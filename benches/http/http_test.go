package http_test

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"io"
	"testing"

	h1 "net/http"

	"github.com/flier/gocombine/benches/http"
)

//go:embed testdata/http-requests.txt
var requests string

func BenchmarkRequest(b *testing.B) {
	p := http.Parser()

	b.SetBytes(int64(len(requests)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		remaining := []byte(requests)

		for len(remaining) > 0 {
			// var r *http.Request
			var err error

			if _, remaining, err = p(remaining); err != nil {
				b.Fatal(err)
			} else {
				// b.Logf("%#+v", r)
			}
		}
	}
}

func BenchmarkStdHttpRequest(b *testing.B) {
	b.SetBytes(int64(len(requests)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r := bufio.NewReader(bytes.NewReader([]byte(requests)))

		for r.Size() > 0 {
			_, err := h1.ReadRequest(r)
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}
