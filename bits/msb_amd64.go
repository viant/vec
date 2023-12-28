package bits

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape
func _msb_avx2(data unsafe.Pointer, len uint64, result unsafe.Pointer)
func Msb(data []uint64) int {
	var res uint64
	_msb_avx2(unsafe.Pointer(&data[0]), uint64(len(data)|cpu.Info), unsafe.Pointer(&res))
	return int(res)
}
