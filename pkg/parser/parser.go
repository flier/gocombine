package parser

import "github.com/flier/gocombine/pkg/stream"

type Parser[S stream.Stream[T], T stream.Token, O any] interface {
	Parse(input S) (out O, remaining S, err error)
}

type ParseFunc[S stream.Stream[T], T stream.Token, O any] func(input S) (out O, remaining S, err error)

func (f ParseFunc[S, T, O]) Parse(input S) (out O, remaining S, err error) { return f(input) }
