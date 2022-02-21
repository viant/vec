package pack

import "math/bits"

func maxBits(in []uint32) uint32 {
	acc := uint32(0)
	for _, v := range in {
		acc |= v
	}
	return uint32(bits.Len64(uint64(acc)))
}

func maxBits2(in []uint32) uint32 {
	min := in[0]
	max := in[0]
	for _, v := range in {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return uint32(bits.Len64(uint64(max - min)))
}
