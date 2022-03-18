package le

import (
	"encoding/binary"

	"github.com/flier/gocombine/pkg/parser"
	"github.com/flier/gocombine/pkg/parser/bytes/num"
	"github.com/flier/gocombine/pkg/stream"
)

// Uint16 reads a uint16 out of the byte stream with the little endianess
func Uint16[S stream.Stream[byte]]() parser.Func[S, byte, uint16] {
	return num.Uint16[S](binary.LittleEndian)
}

// Uint32 reads a uint32 out of the byte stream with the little endianess
func Uint32[S stream.Stream[byte]]() parser.Func[S, byte, uint32] {
	return num.Uint32[S](binary.LittleEndian)
}

// Uint64 reads a uint64 out of the byte stream with the little endianess
func Uint64[S stream.Stream[byte]]() parser.Func[S, byte, uint64] {
	return num.Uint64[S](binary.LittleEndian)
}

// Int16 reads a int16 out of the byte stream with the little endianess
func Int16[S stream.Stream[byte]]() parser.Func[S, byte, int16] {
	return num.Int16[S](binary.LittleEndian)
}

// Int32 reads a int32 out of the byte stream with the little endianess
func Int32[S stream.Stream[byte]]() parser.Func[S, byte, int32] {
	return num.Int32[S](binary.LittleEndian)
}

// Int64 reads a int64 out of the byte stream with the little endianess
func Int64[S stream.Stream[byte]]() parser.Func[S, byte, int64] {
	return num.Int64[S](binary.LittleEndian)
}

// Float32 reads a float32 out of the byte stream with the little endianess
func Float32[S stream.Stream[byte]]() parser.Func[S, byte, float32] {
	return num.Float32[S](binary.LittleEndian)
}

// Float64 reads a float64 out of the byte stream with the little endianess
func Float64[S stream.Stream[byte]]() parser.Func[S, byte, float64] {
	return num.Float64[S](binary.LittleEndian)
}
