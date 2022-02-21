package bits

import (
	"unsafe"
)

//go:noescape
func popcnt(src unsafe.Pointer, len uint64) (ret int)
func Count(data []uint64) int {
	return popcnt(unsafe.Pointer(&data[0]), uint64(len(data)*8))
}
