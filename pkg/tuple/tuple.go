package tuple

type Tuple3[T1, T2, T3 any] struct {
	I1 T1
	I2 T2
	I3 T3
}

func New3[T1, T2, T3 any](i1 T1, i2 T2, i3 T3) Tuple3[T1, T2, T3] {
	return Tuple3[T1, T2, T3]{i1, i2, i3}
}

type Tuple4[T1, T2, T3, T4 any] struct {
	I1 T1
	I2 T2
	I3 T3
	I4 T4
}

func New4[T1, T2, T3, T4 any](i1 T1, i2 T2, i3 T3, i4 T4) Tuple4[T1, T2, T3, T4] {
	return Tuple4[T1, T2, T3, T4]{i1, i2, i3, i4}
}

type Tuple5[T1, T2, T3, T4, T5 any] struct {
	I1 T1
	I2 T2
	I3 T3
	I4 T4
	I5 T5
}

func New5[T1, T2, T3, T4, T5 any](i1 T1, i2 T2, i3 T3, i4 T4, i5 T5) Tuple5[T1, T2, T3, T4, T5] {
	return Tuple5[T1, T2, T3, T4, T5]{i1, i2, i3, i4, i5}
}
