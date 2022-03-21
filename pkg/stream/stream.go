package stream

import (
	"io"

	"golang.org/x/exp/slices"
)

// Token represents a token.
type Token interface {
	~byte | ~uint16 | ~rune
}

// Stream represents a stream of tokens which can be parsed.
type Stream[T Token] interface {
	~[]T
}

// Empty returns true if the stream is empty.
func Empty[S Stream[T], T Token](s S) bool {
	return len(s) == 0
}

// Len returns the number of tokens in the stream.
func Len[S Stream[T], T Token](s S) int {
	return len(s)
}

// Uncons takes a stream `input` and removes its first token, yielding the `tok` and the `remaining` of the elements.
// Returns `err` if no element could be retrieved.
func Uncons[S Stream[T], T Token](input S) (tok T, remaining S, err error) {
	if Empty(input) {
		remaining, err = input, io.ErrUnexpectedEOF
	} else {
		tok, remaining = input[0], input[1:]
	}

	return
}

// UnconsRange takes `size` elements from the stream.
// Fails if the length of the stream is less than `size`.
func UnconsRange[S Stream[T], T Token](input S, size int) (tokens []T, remaining S, err error) {
	if Len(input) < size {
		remaining, err = input, io.ErrUnexpectedEOF
	} else {
		tokens, remaining = input[:size], input[size:]
	}

	return
}

// UnconsWhile takes items from stream, testing each one with `predicate`.
// returns the range of items which passed `predicate`.
func UnconsWhile[S Stream[T], T Token](input S, predicate func(T) bool) (tokens []T, remaining S, err error) {
	if i := slices.IndexFunc(input, func(tok T) bool { return !predicate(tok) }); i >= 0 {
		tokens, remaining = input[:i], input[i:]
	} else {
		tokens = input
	}

	return
}

// UnconsUntil takes items from stream, testing each one with `predicate`.
// returns the range of items which not passed `predicate`.
func UnconsUntil[S Stream[T], T Token](input S, predicate func(T) bool) (tokens []T, remaining S, err error) {
	i := slices.IndexFunc(input, func(tok T) bool { return predicate(tok) })
	switch i {
	case -1:
		tokens = input
	case 0:
		remaining = input
	default:
		tokens, remaining = input[:i-1], input[i-1:]
	}

	return
}

// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
func Index[S Stream[T], T Token](s S, sep []T) int {
	n := len(sep)

	switch {
	case n == 0:
		return 0

	case n == 1:
		return slices.Index(s, sep[0])

	case n == Len(s):
		if slices.Equal(s, sep) {
			return 0
		}

		return -1

	default: // slow path
		for i := 0; i < Len(s)-n+1; i++ {
			if slices.Equal(s[i:i+n], sep) {
				return i
			}
		}

		return -1
	}
}
