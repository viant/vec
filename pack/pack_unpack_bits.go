package pack

func (o Uint32s) packBits(in []uint32, nbits int) {
	idx := 0
	bitoffset := 0
	for _, v := range in {
		o[idx] |= v << bitoffset
		bitoffset += nbits
		if bitoffset >= 32 {
			idx += 1
			bitoffset &= 0x1F
			o[idx] = v >> (nbits - bitoffset)
		}
	}
}

func (o Uint32s) unpackBits(in []uint32, nbits int) {
	index := 0
	offset := 0
	var mask uint32 = (1 << nbits) - 1
	for i := 0; i < len(o); i++ {
		o[i] = (in[index] >> offset) & mask
		offset += nbits
		index += offset >> 5
		if offset > 32 {
			extra := offset & 0x1F
			o[i] |= (in[index] & ((1 << extra) - 1)) << (nbits - extra)
		}
		offset &= 0x1F
	}
}

func PackedUint32Count(size, bits uint32) uint32 {
	return (size*bits + 31) >> 5
}
