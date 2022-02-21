package lut

import (
	"unsafe"
)

//go:noescape
func _lookup_neon(input, output, table unsafe.Pointer, size int)
func Lookup(input, table []byte, output *[]byte) {
	_lookup_neon(unsafe.Pointer(&input[0]), unsafe.Pointer(&(*output)[0]), unsafe.Pointer(&table[0]), len(input))
}

//go:noescape
func _lookup_sve_vla(input, output, table unsafe.Pointer, size int)
func lookup_sve_vla(input, table []byte, output *[]byte) {
	_lookup_sve_vla(unsafe.Pointer(&input[0]), unsafe.Pointer(&(*output)[0]), unsafe.Pointer(&table[0]), len(input))
}

//go:noescape
func _lookup_sve_2048(input, output, table unsafe.Pointer, size int)
func lookup_sve_2048(input, table []byte, output *[]byte) {
	_lookup_sve_2048(unsafe.Pointer(&input[0]), unsafe.Pointer(&(*output)[0]), unsafe.Pointer(&table[0]), len(input))
}
