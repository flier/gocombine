package parser

import "github.com/flier/gocombine/pkg/stream"

type ParseFunc[S stream.Stream[T], T stream.Token, O any] func(s S) (t O, r S, err error)
