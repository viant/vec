package pack

import (
	"unsafe"
)

func (o Uint64s) expand(input uint64) {
	for i := 0; i < 64; i++ {
		if input&(1<<i) != 0 {
			*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&o[0])) + uintptr(i))) = 0xff
		}
	}
}
