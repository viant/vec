package bits

import (
	"unsafe"
)

//go:noescape

func _populate_positions_avx2(input uint64, out unsafe.Pointer, count unsafe.Pointer)
func (p *Positions) Populate(input uint64) uint8 {
	var count uint64
	_populate_positions_avx2(input, unsafe.Pointer(&(*p)[0]), unsafe.Pointer(&count))
	return uint8(count)
}

func _populate_positions_ctz(input uint64, out unsafe.Pointer, count unsafe.Pointer)
func (p *Positions) populate_ctz(input uint64) uint8 {
	var count uint64
	_populate_positions_ctz(input, unsafe.Pointer(&(*p)[0]), unsafe.Pointer(&count))
	return uint8(count)
}
