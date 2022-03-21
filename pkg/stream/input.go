package stream

// Input represents a stream-like input.
type Input interface {
	byte | rune | ~[]byte | ~[]uint16 | ~[]rune | ~string
}

// New returns a stream from a slice of tokens.
func New[S Stream[T], T Token, I Input](input I) (s S) {
	switch v := interface{}(input).(type) {
	case byte:
		s = New[S]([]byte{v})

	case rune:
		s = New[S]([]rune{v})

	case []byte:
		s, _ = interface{}(v).([]T)

	case []uint16:
		s, _ = interface{}(v).([]T)

	case []rune:
		s, _ = interface{}(v).([]T)

	case string:
		s = New[S]([]rune(v))
	}

	return
}
