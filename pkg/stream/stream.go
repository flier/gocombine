package stream

import (
	"io"
)

type Token interface {
	~byte | ~uint16 | ~rune
}

type Stream[T Token] interface {
	~[]T
}

func Empty[S Stream[T], T Token](s S) bool {
	return len(s) == 0
}

func Len[S Stream[T], T Token](s S) int {
	return len(s)
}

// Takes a stream `input` and removes its first token, yielding the `tok` and the `remaining` of the elements.
// Returns `err` if no element could be retrieved.
func Uncons[S Stream[T], T Token](input S) (tok T, remaining S, err error) {
	if Empty(input) {
		remaining, err = input, io.ErrUnexpectedEOF
	} else {
		tok, remaining = input[0], input[1:]
	}
	return
}
