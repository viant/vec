package pack

import "unsafe"

//go:noescape

func _turbopack32(in unsafe.Pointer, number, bit int, out unsafe.Pointer)
func (o Uint32s) PackBits(in []uint32, nbits int) {
	_turbopack32(unsafe.Pointer(&in[0]), len(in), nbits, unsafe.Pointer(&o[0]))
}

func _turbounpack32(in unsafe.Pointer, number, bit int, out unsafe.Pointer)
func (o Uint32s) UnpackBits(in []uint32, nbits int) {
	_turbounpack32(unsafe.Pointer(&in[0]), len(o), nbits, unsafe.Pointer(&o[0]))
}
