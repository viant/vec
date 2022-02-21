package pack

import (
	"unsafe"
)

func (u *Uint64) shrink(vec []uint64) {
	for i := 0; i < 8*len(vec); i++ {
		if *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&vec[0])) + uintptr(i)))&0x80 != 0 {
			*u |= 1 << i
		}
	}
}

func (u *Uint64) shrinkValue(value byte, vec []uint64) {
	for i := 0; i < 8*len(vec); i++ {
		if *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&vec[0])) + uintptr(i))) == value {
			*u |= 1 << i
		}
	}
}
