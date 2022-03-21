package le

import (
	"encoding/binary"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/bytes/num"
)

// Uint16 reads a uint16 out of the byte stream with the little endianess.
func Uint16() parser.Func[byte, uint16] {
	return num.Uint16(binary.LittleEndian)
}

// Uint32 reads a uint32 out of the byte stream with the little endianess.
func Uint32() parser.Func[byte, uint32] {
	return num.Uint32(binary.LittleEndian)
}

// Uint64 reads a uint64 out of the byte stream with the little endianess.
func Uint64() parser.Func[byte, uint64] {
	return num.Uint64(binary.LittleEndian)
}

// Int16 reads a int16 out of the byte stream with the little endianess.
func Int16() parser.Func[byte, int16] {
	return num.Int16(binary.LittleEndian)
}

// Int32 reads a int32 out of the byte stream with the little endianess.
func Int32() parser.Func[byte, int32] {
	return num.Int32(binary.LittleEndian)
}

// Int64 reads a int64 out of the byte stream with the little endianess.
func Int64() parser.Func[byte, int64] {
	return num.Int64(binary.LittleEndian)
}

// Float32 reads a float32 out of the byte stream with the little endianess.
func Float32() parser.Func[byte, float32] {
	return num.Float32(binary.LittleEndian)
}

// Float64 reads a float64 out of the byte stream with the little endianess.
func Float64() parser.Func[byte, float64] {
	return num.Float64(binary.LittleEndian)
}
