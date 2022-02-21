package pack

import (
	"unsafe"
)

//go:noescape
func _packed_mask_neon(in, out unsafe.Pointer, size uint64)

func (u *Uint64) Shrink(vec []uint64) {
	size := uint64(len(vec))
	_packed_mask_neon(unsafe.Pointer(&vec[0]), unsafe.Pointer(u), size)
}

//go:noescape
func _packed_mask2_neon(in, out unsafe.Pointer, value, size uint64)

func (u *Uint64) ShrinkValue(value byte, vec []uint64) {
	size := uint64(len(vec))
	_packed_mask2_neon(unsafe.Pointer(&vec[0]), unsafe.Pointer(u), uint64(value), size)
}
