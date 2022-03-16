package stream

type Input interface {
	~[]byte | ~[]uint16 | ~[]rune | ~string
}

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
