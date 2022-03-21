package be

import (
	"encoding/binary"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/bytes/num"
)

// Uint16 reads a uint16 out of the byte stream with the big endianess.
func Uint16() parser.Func[byte, uint16] {
	return num.Uint16(binary.BigEndian)
}

// Uint32 reads a uint32 out of the byte stream with the big endianess.
func Uint32() parser.Func[byte, uint32] {
	return num.Uint32(binary.BigEndian)
}

// Uint64 reads a uint64 out of the byte stream with the big endianess.
func Uint64() parser.Func[byte, uint64] {
	return num.Uint64(binary.BigEndian)
}

// Int16 reads a int16 out of the byte stream with the big endianess.
func Int16() parser.Func[byte, int16] {
	return num.Int16(binary.BigEndian)
}

// Int32 reads a int32 out of the byte stream with the big endianess.
func Int32() parser.Func[byte, int32] {
	return num.Int32(binary.BigEndian)
}

// Int64 reads a int64 out of the byte stream with the big endianess.
func Int64() parser.Func[byte, int64] {
	return num.Int64(binary.BigEndian)
}

// Float32 reads a float32 out of the byte stream with the big endianess.
func Float32() parser.Func[byte, float32] {
	return num.Float32(binary.BigEndian)
}

// Float64 reads a float64 out of the byte stream with the big endianess.
func Float64() parser.Func[byte, float64] {
	return num.Float64(binary.BigEndian)
}
