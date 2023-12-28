package bits

import "math/bits"

func msb(data []uint64) int {
	for i, d := range data {
		if d != 0 {
			return i*64 + 63 - bits.LeadingZeros64(d)
		}
	}
	return -1
}
