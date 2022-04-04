package stream

import (
	"reflect"
	"unsafe"
)

type Offset uintptr

const EOI Offset = Offset(^uintptr(0))

func Pos[T Token](s []T) Offset {
	if len(s) == 0 {
		return EOI
	}

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))

	return Offset(hdr.Data)
}

func Translate[T Token](s []T, off Offset) Position {
	if off == EOI {
		return Position(len(s))
	}

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))

	return Position(uintptr(off)-hdr.Data) / Position(unsafe.Sizeof(T(0)))
}

func Distance[T Token](cur, end []T) int {
	if len(cur) == 0 {
		return len(end)
	}

	n := Pos(cur) - Pos(end)

	return int(n) / int(unsafe.Sizeof(T(0)))
}

func Checkpoint[T Token](s []T) []T {
	return s
}

type Position uint

type LineNumber struct {
	Line   int
	Column int
}

func NewLineNumber[T Token](input []T) *LineNumber {
	var line, column int

	for _, t := range input {
		if t == '\n' {
			line++

			column = 0
		} else {
			column++
		}
	}

	return &LineNumber{line, column}
}

type Span struct {
	Start, End Offset
}
