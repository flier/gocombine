package parser

import (
	"fmt"

	"github.com/flier/gocombine/pkg/stream"
)

type UnexpectedErr[T stream.Token] struct {
	Expected []T
	Actual   []T
}

func (e *UnexpectedErr[T]) Error() string {
	if e.Expected != nil {
		return fmt.Sprintf("expected `%c`, got `%c`", e.Expected, e.Actual)
	}

	return fmt.Sprintf("unexpected `%c`", e.Actual)
}

func Unexpected[T stream.Token](expected, actual []T) error {
	return &UnexpectedErr[T]{expected, actual}
}
