package stream

import "io"

type Token interface {
	~byte | ~rune
}

type Stream[T Token] interface {
	~[]T
}

func Uncons[S Stream[T], T Token](s S) (t T, err error) {
	if len(s) == 0 {
		return 0, io.EOF
	}
	return s[0], nil
}
