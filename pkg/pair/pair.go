package pair

// Pair provides a way to store two heterogeneous types as a single unit.
type Pair[F, S any] struct {
	First  F
	Second S
}

// New creates a new Pair[F, S].
func New[F, S any](first F, second S) Pair[F, S] {
	return Pair[F, S]{first, second}
}
