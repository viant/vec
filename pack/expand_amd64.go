package pack

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape
func _expand_mask_avx(in, out unsafe.Pointer, info uint64)

func (o Uint64s) Expand(input uint64) {
	_expand_mask_avx(unsafe.Pointer(&input), unsafe.Pointer(&o[0]), uint64(cpu.Info))
}
