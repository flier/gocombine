package stream

import "io"

type Token interface {
	~byte | ~uint16 | ~rune
}

type Stream[T Token] interface {
	~[]T
}

func Uncons[S Stream[T], T Token](input S) (tok T, remaining S, err error) {
	if len(input) == 0 {
		err = io.ErrUnexpectedEOF
	} else {
		tok, remaining = input[0], input[1:]
	}
	return
}
