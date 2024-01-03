package bits

import (
	"unsafe"
)

//go:noescape
func _lsb_neon(in unsafe.Pointer, size int, out unsafe.Pointer)
func Lsb(data []uint64) int {
	var res uint64
	_lsb_neon(unsafe.Pointer(&data[0]), len(data), unsafe.Pointer(&res))
	return int(res)
}
