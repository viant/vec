package lut

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape
func _lookup_subrange_avx(input, output, range_upper, table unsafe.Pointer, subranges, info uint64)

func LookupSubrange(input, range_upper, table []byte, output *[]byte) {
	_lookup_subrange_avx(unsafe.Pointer(&input[0]), unsafe.Pointer(&(*output)[0]), unsafe.Pointer(&range_upper[0]),
		unsafe.Pointer(&table[0]), uint64(len(range_upper)), uint64(len(input)|cpu.Info))
}
