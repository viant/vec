package pack

import "unsafe"

//go:noescape
func _maxbits(in unsafe.Pointer, size int) (bits uint32)

func MaxBits(in []uint32) uint32 {
	return _maxbits(unsafe.Pointer(&in[0]), len(in))
}
