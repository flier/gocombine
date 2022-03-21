package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
	"github.com/flier/gocombine/pkg/tuple"
)

// Tuple3 parses a tuple of heterogeneous values.
func Tuple3[

	T stream.Token,
	O1, O2, O3 any,
](
	p1 parser.Func[T, O1],
	p2 parser.Func[T, O2],
	p3 parser.Func[T, O3],
) parser.Func[T, tuple.Tuple3[O1, O2, O3]] {
	return Attempt(func(input []T) (out tuple.Tuple3[O1, O2, O3], remaining []T, err error) {
		var o1 O1
		if o1, remaining, err = p1.Parse(input); err != nil {
			return
		}

		var o2 O2
		if o2, remaining, err = p2.Parse(remaining); err != nil {
			return
		}

		var o3 O3
		if o3, remaining, err = p3.Parse(remaining); err != nil {
			return
		}

		out = tuple.New3(o1, o2, o3)

		return
	})
}

// Tuple4 parses a tuple of heterogeneous values.
func Tuple4[

	T stream.Token,
	O1, O2, O3, O4 any,
](
	p1 parser.Func[T, O1],
	p2 parser.Func[T, O2],
	p3 parser.Func[T, O3],
	p4 parser.Func[T, O4],
) parser.Func[T, tuple.Tuple4[O1, O2, O3, O4]] {
	return Attempt(func(input []T) (out tuple.Tuple4[O1, O2, O3, O4], remaining []T, err error) {
		var o1 O1
		if o1, remaining, err = p1.Parse(input); err != nil {
			return
		}

		var o2 O2
		if o2, remaining, err = p2.Parse(remaining); err != nil {
			return
		}

		var o3 O3
		if o3, remaining, err = p3.Parse(remaining); err != nil {
			return
		}

		var o4 O4
		if o4, remaining, err = p4.Parse(remaining); err != nil {
			return
		}

		out = tuple.New4(o1, o2, o3, o4)

		return
	})
}

// Tuple5 parses a tuple of heterogeneous values.
func Tuple5[

	T stream.Token,
	O1, O2, O3, O4, O5 any,
](
	p1 parser.Func[T, O1],
	p2 parser.Func[T, O2],
	p3 parser.Func[T, O3],
	p4 parser.Func[T, O4],
	p5 parser.Func[T, O5],
) parser.Func[T, tuple.Tuple5[O1, O2, O3, O4, O5]] {
	return Attempt(func(input []T) (out tuple.Tuple5[O1, O2, O3, O4, O5], remaining []T, err error) {
		var o1 O1
		if o1, remaining, err = p1.Parse(input); err != nil {
			return
		}

		var o2 O2
		if o2, remaining, err = p2.Parse(remaining); err != nil {
			return
		}

		var o3 O3
		if o3, remaining, err = p3.Parse(remaining); err != nil {
			return
		}

		var o4 O4
		if o4, remaining, err = p4.Parse(remaining); err != nil {
			return
		}

		var o5 O5
		if o5, remaining, err = p5.Parse(remaining); err != nil {
			return
		}

		out = tuple.New5(o1, o2, o3, o4, o5)

		return
	})
}
