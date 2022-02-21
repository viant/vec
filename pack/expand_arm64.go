package pack

import (
	"unsafe"
)

//go:noescape
func _expand_mask_arm(in, out unsafe.Pointer)

func (o Uint64s) Expand(input uint64) {
	_expand_mask_arm(unsafe.Pointer(&input), unsafe.Pointer(&o[0]))
}
