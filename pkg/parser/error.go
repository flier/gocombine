package parser

import (
	"errors"
	"fmt"

	"github.com/flier/gocombine/pkg/stream"
)

var (
	ErrExpected   = errors.New("expected")
	ErrUnexpected = errors.New("unexpected")
)

type UnexpectedErr[T stream.Token] struct {
	Expected []T
	Actual   []T
}

func (e *UnexpectedErr[T]) Err() error {
	if e.Expected != nil {
		return fmt.Errorf("expected `%v`, actual `%v`, %w", e.Expected, e.Actual, ErrUnexpected)
	}

	return fmt.Errorf("actual `%v`, %w", e.Actual, ErrUnexpected)
}

func (e *UnexpectedErr[T]) Error() string {
	return e.Err().Error()
}

func Unexpected[T stream.Token](expected, actual []T) error {
	return &UnexpectedErr[T]{expected, actual}
}
