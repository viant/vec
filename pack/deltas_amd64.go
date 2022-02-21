package pack

import "unsafe"

//go:noescape
func _compute_deltas(input unsafe.Pointer, length int, output unsafe.Pointer, starting_point int)
func (o Int32s) Deltas(in []int32) {
	_compute_deltas(unsafe.Pointer(&in[0]), len(in), unsafe.Pointer(&o[0]), 0)
}

func _compute_prefix_sum(input unsafe.Pointer, length int, output unsafe.Pointer, starting_point int)
func (o Int32s) RecoverDeltas(in []int32) {
	_compute_prefix_sum(unsafe.Pointer(&in[0]), len(in), unsafe.Pointer(&o[0]), 0)
}
