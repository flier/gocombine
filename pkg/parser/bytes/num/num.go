package num

import (
	"encoding/binary"
	"math"
	"unsafe"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/ranges"
	"github.com/flier/gocombine/pkg/stream"
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
func Uint16[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, uint16] {
	return combinator.Map(ranges.Take[S](sizeofUint16), endian.Uint16).Expected("uint16")
}

// Uint32 reads a uint32 out of the byte stream with the specified endianess.
func Uint32[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, uint32] {
	return combinator.Map(ranges.Take[S](sizeofUint32), endian.Uint32).Expected("uint32")
}

// Uint64 reads a uint64 out of the byte stream with the specified endianess.
func Uint64[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, uint64] {
	return combinator.Map(ranges.Take[S](sizeofUint64), endian.Uint64).Expected("uint64")
}

// Int16 reads a int16 out of the byte stream with the specified endianess.
func Int16[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, int16] {
	return combinator.Map(ranges.Take[S](sizeofInt16), func(b []byte) int16 {
		return int16(endian.Uint16(b))
	}).Expected("int16")
}

// Int32 reads a int32 out of the byte stream with the specified endianess.
func Int32[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, int32] {
	return combinator.Map(ranges.Take[S](sizeofInt32), func(b []byte) int32 {
		return int32(endian.Uint32(b))
	}).Expected("int32")
}

// Int64 reads a int64 out of the byte stream with the specified endianess.
func Int64[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, int64] {
	return combinator.Map(ranges.Take[S](sizeofInt64), func(b []byte) int64 {
		return int64(endian.Uint64(b))
	}).Expected("int64")
}

// Float32 reads a float32 out of the byte stream with the specified endianess.
func Float32[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, float32] {
	return combinator.Map(ranges.Take[S](sizeofFloat32), func(b []byte) float32 {
		return math.Float32frombits(endian.Uint32(b))
	}).Expected("float32")
}

// Float64 reads a float64 out of the byte stream with the specified endianess.
func Float64[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, float64] {
	return combinator.Map(ranges.Take[S](sizeofFloat64), func(b []byte) float64 {
		return math.Float64frombits(endian.Uint64(b))
	}).Expected("float64")
}
