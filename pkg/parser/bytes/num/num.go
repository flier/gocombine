package num

import (
	"encoding/binary"
	"math"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/combinator"
	"github.com/flier/gocombine/pkg/parser/ranges"
	"github.com/flier/gocombine/pkg/stream"
)

// Uint16 reads a uint16 out of the byte stream with the specified endianess
func Uint16[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, uint16] {
	return combinator.Map(ranges.Take[S](2), endian.Uint16).Expected("uint16")
}

// Uint32 reads a uint32 out of the byte stream with the specified endianess
func Uint32[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, uint32] {
	return combinator.Map(ranges.Take[S](4), endian.Uint32).Expected("uint32")
}

// Uint64 reads a uint64 out of the byte stream with the specified endianess
func Uint64[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, uint64] {
	return combinator.Map(ranges.Take[S](8), endian.Uint64).Expected("uint64")
}

// Int16 reads a int16 out of the byte stream with the specified endianess
func Int16[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, int16] {
	return combinator.Map(ranges.Take[S](2), func(b []byte) int16 {
		return int16(endian.Uint16(b))
	}).Expected("int16")
}

// Int32 reads a int32 out of the byte stream with the specified endianess
func Int32[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, int32] {
	return combinator.Map(ranges.Take[S](4), func(b []byte) int32 {
		return int32(endian.Uint32(b))
	}).Expected("int32")
}

// Int64 reads a int64 out of the byte stream with the specified endianess
func Int64[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, int64] {
	return combinator.Map(ranges.Take[S](8), func(b []byte) int64 {
		return int64(endian.Uint64(b))
	}).Expected("int64")
}

// Float32 reads a float32 out of the byte stream with the specified endianess
func Float32[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, float32] {
	return combinator.Map(ranges.Take[S](4), func(b []byte) float32 {
		return math.Float32frombits(endian.Uint32(b))
	}).Expected("float32")
}

// Float64 reads a float64 out of the byte stream with the specified endianess
func Float64[S stream.Stream[byte]](endian binary.ByteOrder) parser.Func[S, byte, float64] {
	return combinator.Map(ranges.Take[S](8), func(b []byte) float64 {
		return math.Float64frombits(endian.Uint64(b))
	}).Expected("float64")
}
