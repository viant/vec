package bits

import "unsafe"

//go:noescape
func _msb_neon(in unsafe.Pointer, size int, out unsafe.Pointer)
func Msb(data []uint64) int {
	var res uint64
	_msb_neon(unsafe.Pointer(&data[0]), len(data), unsafe.Pointer(&res))
	return int(res)
}
