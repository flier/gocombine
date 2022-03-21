package tuple

// Tuple3 is a fixed-size collection of heterogeneous values.
type Tuple3[T1, T2, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

// New3 creates a new Tuple3[T1, T2, T3].
func New3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Tuple3[T1, T2, T3] {
	return Tuple3[T1, T2, T3]{v1, v2, v3}
}

// Tuple4 is a fixed-size collection of heterogeneous values.
type Tuple4[T1, T2, T3, T4 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
}

// New4 creates a new Tuple4[T1, T2, T3, T4].
func New4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Tuple4[T1, T2, T3, T4] {
	return Tuple4[T1, T2, T3, T4]{v1, v2, v3, v4}
}

// Tuple5 is a fixed-size collection of heterogeneous values.
type Tuple5[T1, T2, T3, T4, T5 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
}

// New5 creates a new Tuple5[T1, T2, T3, T4, T5].
func New5[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Tuple5[T1, T2, T3, T4, T5] {
	return Tuple5[T1, T2, T3, T4, T5]{v1, v2, v3, v4, v5}
}
