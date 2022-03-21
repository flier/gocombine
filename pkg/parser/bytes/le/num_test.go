package le_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/bytes/le"
)

func ExampleUint16() {
	p := le.Uint16()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 13330 [86 120 144 171 205 239] <nil>
}

func ExampleUint32() {
	p := le.Uint32()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 2018915346 [144 171 205 239] <nil>
}

func ExampleUint64() {
	p := le.Uint64()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 17279655982273016850 [] <nil>
}

func ExampleInt16() {
	p := le.Int16()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 13330 [86 120 144 171 205 239] <nil>
}

func ExampleInt32() {
	p := le.Int32()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 2018915346 [144 171 205 239] <nil>
}

func ExampleInt64() {
	p := le.Int64()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// -1167088091436534766 [] <nil>
}

func ExampleFloat32() {
	p := le.Float32()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 1.7378244e+34 [144 171 205 239] <nil>
}

func ExampleFloat64() {
	p := le.Float64()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// -3.5987094278483163e+230 [] <nil>
}
