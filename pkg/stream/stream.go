package stream

import (
	"io"

	"golang.org/x/exp/slices"
)

// Token represents a token.
type Token interface {
	~byte | ~uint16 | ~rune
}

// Empty returns true if the stream is empty.
func Empty[T Token](s []T) bool {
	return len(s) == 0
}

// Len returns the number of tokens in the stream.
func Len[T Token](s []T) int {
	return len(s)
}

// Uncons takes a stream `input` and removes its first token, yielding the `tok` and the `remaining` of the elements.
// Returns `err` if no element could be retrieved.
func Uncons[T Token](input []T) (tok T, remaining []T, err error) {
	if Empty(input) {
		remaining, err = input, io.ErrUnexpectedEOF
	} else {
		tok, remaining = input[0], input[1:]
	}

	return
}

// UnconsRange takes `size` elements from the stream.
// Fails if the length of the stream is less than `size`.
func UnconsRange[T Token](input []T, size int) (tokens []T, remaining []T, err error) {
	if Len(input) < size {
		remaining, err = input, io.ErrUnexpectedEOF
	} else {
		tokens, remaining = input[:size], input[size:]
	}

	return
}

// UnconsWhile takes items from stream, testing each one with `predicate`.
// returns the range of items which passed `predicate`.
func UnconsWhile[T Token](input []T, predicate func(T) bool) (tokens []T, remaining []T, err error) {
	i := slices.IndexFunc(input, func(tok T) bool { return !predicate(tok) })

	if i >= 0 {
		tokens, remaining = input[:i], input[i:]
	} else {
		tokens = input
	}

	return
}

// UnconsUntil takes items from stream, testing each one with `predicate`.
// returns the range of items which not passed `predicate`.
func UnconsUntil[T Token](input []T, predicate func(T) bool) (tokens []T, remaining []T, err error) {
	i := slices.IndexFunc(input, func(tok T) bool { return predicate(tok) })

	switch i {
	case -1:
		tokens, err = input, io.ErrUnexpectedEOF
	case 0:
		remaining = input
	default:
		tokens, remaining = input[:i], input[i:]
	}

	return
}
