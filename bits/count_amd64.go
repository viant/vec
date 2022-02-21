package bits

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape
func _count_optimized(data unsafe.Pointer, len uint64, result unsafe.Pointer)
func Count(data []uint64) int {
	var res uint64
	_count_optimized(unsafe.Pointer(&data[0]), uint64(len(data)|cpu.Info), unsafe.Pointer(&res))
	return int(res)
}
