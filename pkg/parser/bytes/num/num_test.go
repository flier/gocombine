package num_test

import (
	"encoding/binary"
	"fmt"

	"github.com/flier/gocombine/pkg/parser/bytes/num"
)

func ExampleUint16() {
	p := num.Uint16[[]byte](binary.LittleEndian)

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 13330 [86 120 144 171 205 239] <nil>
}

func ExampleUint32() {
	p := num.Uint32[[]byte](binary.LittleEndian)

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 2018915346 [144 171 205 239] <nil>
}

func ExampleUint64() {
	p := num.Uint64[[]byte](binary.LittleEndian)

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 17279655982273016850 [] <nil>
}

func ExampleInt16() {
	p := num.Int16[[]byte](binary.LittleEndian)

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 13330 [86 120 144 171 205 239] <nil>
}

func ExampleInt32() {
	p := num.Int32[[]byte](binary.LittleEndian)

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 2018915346 [144 171 205 239] <nil>
}

func ExampleInt64() {
	p := num.Int64[[]byte](binary.LittleEndian)

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// -1167088091436534766 [] <nil>
}

func ExampleFloat32() {
	p := num.Float32[[]byte](binary.LittleEndian)

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 1.7378244e+34 [144 171 205 239] <nil>
}

func ExampleFloat64() {
	p := num.Float64[[]byte](binary.LittleEndian)

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// -3.5987094278483163e+230 [] <nil>
}
