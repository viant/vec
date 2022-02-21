package bits

import "math/bits"

type Positions [64]uint32

func (p *Positions) populate(input uint64) uint8 {
	var count = uint8(0)
	for input != 0 {
		(*p)[count] = uint32(bits.TrailingZeros64(input))
		count++
		input ^= input & -input
	}
	return count
}
