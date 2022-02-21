package lut

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape
func _lookup_512vbmi(input, output, table unsafe.Pointer, info uint64)

func Lookup(input, table []byte, output *[]byte) {
	_lookup_512vbmi(unsafe.Pointer(&input[0]), unsafe.Pointer(&(*output)[0]), unsafe.Pointer(&table[0]), uint64(len(input)|cpu.Info))
}
