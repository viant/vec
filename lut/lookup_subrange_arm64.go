package lut

import (
	"unsafe"
)

//go:noescape
func _lookup_subrange_neon(input, output, range_upper, table unsafe.Pointer, subranges, chunks int)

func LookupSubrange(input, range_upper, table []byte, output *[]byte) {
	_lookup_subrange_neon(unsafe.Pointer(&input[0]), unsafe.Pointer(&(*output)[0]), unsafe.Pointer(&range_upper[0]),
		unsafe.Pointer(&table[0]), len(range_upper), len(input))
}
