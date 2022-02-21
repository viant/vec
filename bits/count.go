package bits

import "math/bits"

func count(data []uint64) int {
	result := 0
	for i := 0; i < len(data); i++ {
		result += bits.OnesCount64(data[i])
	}
	return result
}
