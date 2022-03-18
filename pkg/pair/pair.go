package pair

type Pair[F, S any] struct {
	First  F
	Second S
}

func New[F, S any](first F, second S) Pair[F, S] {
	return Pair[F, S]{first, second}
}
