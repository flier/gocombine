package parser

import (
	"errors"
	"fmt"

	"github.com/flier/gocombine/pkg/stream"
)

var ErrExpected = errors.New("expected")

type UnexpectedErr[T stream.Token] struct {
	Expected []T
	Actual   []T
}

func (e *UnexpectedErr[T]) Error() string {
	if e.Expected != nil {
		return fmt.Sprintf("expected `%v`, got `%v`", e.Expected, e.Actual)
	}

	return fmt.Sprintf("unexpected `%v`", e.Actual)
}

func Unexpected[T stream.Token](expected, actual []T) error {
	return &UnexpectedErr[T]{expected, actual}
}
