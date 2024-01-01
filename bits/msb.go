package bits

import "math/bits"

func msb(data []uint64) int {
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] != 0 {
			return i*64 + 63 - bits.LeadingZeros64(data[i])
		}
	}
	return -1
}
