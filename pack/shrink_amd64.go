package pack

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape
func _packed_mask_avx(in, out unsafe.Pointer, size uint64)

func (u *Uint64) Shrink(vec []uint64) {
	_packed_mask_avx(unsafe.Pointer(&vec[0]), unsafe.Pointer(u), uint64(len(vec)|cpu.Info))
}

//go:noescape
func _packed_mask2_avx(in, out unsafe.Pointer, value, size uint64)

func (u *Uint64) ShrinkValue(value byte, vec []uint64) {
	_packed_mask2_avx(unsafe.Pointer(&vec[0]), unsafe.Pointer(u), uint64(value), uint64(len(vec)|cpu.Info))
}
