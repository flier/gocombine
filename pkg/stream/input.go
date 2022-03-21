package stream

// Input represents a stream-like input.
type Input interface {
	~[]byte | ~[]uint16 | ~[]rune | ~string
}

// New returns a stream from a slice of tokens.
func New[S Stream[T], T Token, I Input](input I) (s S) {
	switch v := interface{}(input).(type) {
	case []byte:
		s = interface{}(v).([]T)
	case []uint16:
		s = interface{}(v).([]T)
	case []rune:
		s = interface{}(v).([]T)
	case string:
		s = interface{}([]rune(v)).([]T)
	}
	return
}
