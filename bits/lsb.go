package bits

import "math/bits"

func lsb(data []uint64) int {
	for i, d := range data {
		if d != 0 {
			return i*64 + bits.TrailingZeros64(d)
		}
	}
	return -1
}
