package num

import (
	"encoding/binary"
	"math"
	"unsafe"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/ranges"
)

const (
	sizeofUint16  = int(unsafe.Sizeof(uint16(0)))
	sizeofUint32  = int(unsafe.Sizeof(uint32(0)))
	sizeofUint64  = int(unsafe.Sizeof(uint64(0)))
	sizeofInt16   = int(unsafe.Sizeof(int16(0)))
	sizeofInt32   = int(unsafe.Sizeof(int32(0)))
	sizeofInt64   = int(unsafe.Sizeof(int64(0)))
	sizeofFloat32 = int(unsafe.Sizeof(float32(0)))
	sizeofFloat64 = int(unsafe.Sizeof(float64(0)))
)

// Uint16 reads a uint16 out of the byte stream with the specified endianess.
func Uint16(endian binary.ByteOrder) parser.Func[byte, uint16] {
	return combinator.Map(ranges.Take[byte](sizeofUint16), endian.Uint16).Expected("uint16")
}

// Uint32 reads a uint32 out of the byte stream with the specified endianess.
func Uint32(endian binary.ByteOrder) parser.Func[byte, uint32] {
	return combinator.Map(ranges.Take[byte](sizeofUint32), endian.Uint32).Expected("uint32")
}

// Uint64 reads a uint64 out of the byte stream with the specified endianess.
func Uint64(endian binary.ByteOrder) parser.Func[byte, uint64] {
	return combinator.Map(ranges.Take[byte](sizeofUint64), endian.Uint64).Expected("uint64")
}

// Int16 reads a int16 out of the byte stream with the specified endianess.
func Int16(endian binary.ByteOrder) parser.Func[byte, int16] {
	return combinator.Map(ranges.Take[byte](sizeofInt16), func(b []byte) int16 {
		return int16(endian.Uint16(b))
	}).Expected("int16")
}

// Int32 reads a int32 out of the byte stream with the specified endianess.
func Int32(endian binary.ByteOrder) parser.Func[byte, int32] {
	return combinator.Map(ranges.Take[byte](sizeofInt32), func(b []byte) int32 {
		return int32(endian.Uint32(b))
	}).Expected("int32")
}

// Int64 reads a int64 out of the byte stream with the specified endianess.
func Int64(endian binary.ByteOrder) parser.Func[byte, int64] {
	return combinator.Map(ranges.Take[byte](sizeofInt64), func(b []byte) int64 {
		return int64(endian.Uint64(b))
	}).Expected("int64")
}

// Float32 reads a float32 out of the byte stream with the specified endianess.
func Float32(endian binary.ByteOrder) parser.Func[byte, float32] {
	return combinator.Map(ranges.Take[byte](sizeofFloat32), func(b []byte) float32 {
		return math.Float32frombits(endian.Uint32(b))
	}).Expected("float32")
}

// Float64 reads a float64 out of the byte stream with the specified endianess.
func Float64(endian binary.ByteOrder) parser.Func[byte, float64] {
	return combinator.Map(ranges.Take[byte](sizeofFloat64), func(b []byte) float64 {
		return math.Float64frombits(endian.Uint64(b))
	}).Expected("float64")
}
