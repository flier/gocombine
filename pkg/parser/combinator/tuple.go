package combinator

import (
	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/stream"
	"github.com/flier/gocombine/pkg/tuple"
)

func Tuple3[
	S stream.Stream[T],
	T stream.Token,
	O1, O2, O3 any,
](
	p1 parser.Func[S, T, O1],
	p2 parser.Func[S, T, O2],
	p3 parser.Func[S, T, O3],
) parser.Func[S, T, tuple.Tuple3[O1, O2, O3]] {
	return func(input S) (out tuple.Tuple3[O1, O2, O3], remaining S, err error) {
		var o1 O1
		if o1, remaining, err = p1.Parse(input); err != nil {
			remaining = input
			return
		}

		var o2 O2
		if o2, remaining, err = p2.Parse(remaining); err != nil {
			remaining = input
			return
		}

		var o3 O3
		if o3, remaining, err = p3.Parse(remaining); err != nil {
			remaining = input
			return
		}

		out = tuple.New3(o1, o2, o3)
		return
	}
}

func Tuple4[
	S stream.Stream[T],
	T stream.Token,
	O1, O2, O3, O4 any,
](
	p1 parser.Func[S, T, O1],
	p2 parser.Func[S, T, O2],
	p3 parser.Func[S, T, O3],
	p4 parser.Func[S, T, O4],
) parser.Func[S, T, tuple.Tuple4[O1, O2, O3, O4]] {
	return func(input S) (out tuple.Tuple4[O1, O2, O3, O4], remaining S, err error) {
		var o1 O1
		if o1, remaining, err = p1.Parse(input); err != nil {
			remaining = input
			return
		}

		var o2 O2
		if o2, remaining, err = p2.Parse(remaining); err != nil {
			remaining = input
			return
		}

		var o3 O3
		if o3, remaining, err = p3.Parse(remaining); err != nil {
			remaining = input
			return
		}

		var o4 O4
		if o4, remaining, err = p4.Parse(remaining); err != nil {
			remaining = input
			return
		}

		out = tuple.New4(o1, o2, o3, o4)
		return
	}
}

func Tuple5[
	S stream.Stream[T],
	T stream.Token,
	O1, O2, O3, O4, O5 any,
](
	p1 parser.Func[S, T, O1],
	p2 parser.Func[S, T, O2],
	p3 parser.Func[S, T, O3],
	p4 parser.Func[S, T, O4],
	p5 parser.Func[S, T, O5],
) parser.Func[S, T, tuple.Tuple5[O1, O2, O3, O4, O5]] {
	return func(input S) (out tuple.Tuple5[O1, O2, O3, O4, O5], remaining S, err error) {
		var o1 O1
		if o1, remaining, err = p1.Parse(input); err != nil {
			remaining = input
			return
		}

		var o2 O2
		if o2, remaining, err = p2.Parse(remaining); err != nil {
			remaining = input
			return
		}

		var o3 O3
		if o3, remaining, err = p3.Parse(remaining); err != nil {
			remaining = input
			return
		}

		var o4 O4
		if o4, remaining, err = p4.Parse(remaining); err != nil {
			remaining = input
			return
		}

		var o5 O5
		if o5, remaining, err = p5.Parse(remaining); err != nil {
			remaining = input
			return
		}

		out = tuple.New5(o1, o2, o3, o4, o5)
		return
	}
}
