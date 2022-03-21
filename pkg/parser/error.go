package parser

import (
	"errors"
	"fmt"

	"github.com/flier/gocombine/pkg/stream"
)

var (
	// ErrExpected indicates an expected value.
	ErrExpected = errors.New("expected")

	// ErrUnexpected indicates a unexpected value.
	ErrUnexpected = errors.New("unexpected")
)

// UnexpectedErr is an error that indicates an unexpected value.
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

// Unexpected returns an error that indicates an unexpected value.
func Unexpected[T stream.Token](expected, actual []T) *UnexpectedErr[T] {
	return &UnexpectedErr[T]{expected, actual}
}
