package be_test

import (
	"fmt"

	"github.com/flier/gocombine/pkg/parser/bytes/be"
)

func ExampleUint16() {
	p := be.Uint16[[]byte]()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 4660 [86 120 144 171 205 239] <nil>
}

func ExampleUint32() {
	p := be.Uint32[[]byte]()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 305419896 [144 171 205 239] <nil>
}

func ExampleUint64() {
	p := be.Uint64[[]byte]()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 1311768467294899695 [] <nil>
}

func ExampleInt16() {
	p := be.Int16[[]byte]()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 4660 [86 120 144 171 205 239] <nil>
}

func ExampleInt32() {
	p := be.Int32[[]byte]()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 305419896 [144 171 205 239] <nil>
}

func ExampleInt64() {
	p := be.Int64[[]byte]()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 1311768467294899695 [] <nil>
}

func ExampleFloat32() {
	p := be.Float32[[]byte]()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 5.6904566e-28 [144 171 205 239] <nil>
}

func ExampleFloat64() {
	p := be.Float64[[]byte]()

	fmt.Println(p([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}))

	// Output:
	// 5.626349108908516e-221 [] <nil>
}
