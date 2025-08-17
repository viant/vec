package bitwise

import "github.com/viant/vec/cpu"

type Strides []uint32

func (s Strides) Init(n int) {
	s[0] = uint32(n)
	s[1] = uint32(cpu.Info >> 32)
	s[2] = 1
	s[3] = 0
	s[4] = uint32(n)
}

func (s Strides) setActiveStrides(set Uint64s) {
	length := len(set)
	s[0] = uint32(length)
	s[1] = uint32(cpu.Info >> 32)

	out := s[3:]
	numSpans := 0

	i := 0
	for i < length {
		// Skip zeros
		for i < length && set[i] == 0 {
			i++
		}
		if i == length {
			break
		}

		// Start of a non-zero span
		start := i
		for i < length && set[i] != 0 {
			i++
		}

		// Write [start, length] as uint32 pairs
		// Ensure we have space; C++ assumes sufficient capacity.
		if 2*numSpans+1 < len(out) {
			out[2*numSpans+0] = uint32(start)     // start index in 64-bit words
			out[2*numSpans+1] = uint32(i - start) // length in 64-bit words
		}
		numSpans++
	}

	s[2] = uint32(numSpans)
}

// AND: 2 inputs (original)
func (o Uint64s) andStrided(v1, v2 Uint64s, strides Strides) {
	numSpans := strides[2]
	spans := strides[3:]

	for s := uint32(0); s < numSpans; s++ {
		start := int(spans[2*s+0])
		length := int(spans[2*s+1])
		end := start + length

		for i := start; i < end; i++ {
			o[i] = v1[i] & v2[i]
		}
	}
}

// AND: 3 inputs
func (o Uint64s) and3Strided(v1, v2, v3 Uint64s, strides Strides) {
	numSpans := strides[2]
	spans := strides[3:]

	for s := uint32(0); s < numSpans; s++ {
		start := int(spans[2*s+0])
		length := int(spans[2*s+1])
		end := start + length

		for i := start; i < end; i++ {
			o[i] = v1[i] & v2[i] & v3[i]
		}
	}
}

// AND: 4 inputs
func (o Uint64s) and4Strided(v1, v2, v3, v4 Uint64s, strides Strides) {
	numSpans := strides[2]
	spans := strides[3:]

	for s := uint32(0); s < numSpans; s++ {
		start := int(spans[2*s+0])
		length := int(spans[2*s+1])
		end := start + length

		for i := start; i < end; i++ {
			o[i] = v1[i] & v2[i] & v3[i] & v4[i]
		}
	}
}

// AND: 5 inputs
func (o Uint64s) and5Strided(v1, v2, v3, v4, v5 Uint64s, strides Strides) {
	numSpans := strides[2]
	spans := strides[3:]

	for s := uint32(0); s < numSpans; s++ {
		start := int(spans[2*s+0])
		length := int(spans[2*s+1])
		end := start + length

		for i := start; i < end; i++ {
			o[i] = v1[i] & v2[i] & v3[i] & v4[i] & v5[i]
		}
	}
}

// AND: 6 inputs
func (o Uint64s) and6Strided(v1, v2, v3, v4, v5, v6 Uint64s, strides Strides) {
	numSpans := strides[2]
	spans := strides[3:]

	for s := uint32(0); s < numSpans; s++ {
		start := int(spans[2*s+0])
		length := int(spans[2*s+1])
		end := start + length

		for i := start; i < end; i++ {
			o[i] = v1[i] & v2[i] & v3[i] & v4[i] & v5[i] & v6[i]
		}
	}
}

// OR: 2 inputs
func (o Uint64s) orStrided(v1, v2 Uint64s, strides Strides) {
	numSpans := strides[2]
	spans := strides[3:]

	for s := uint32(0); s < numSpans; s++ {
		start := int(spans[2*s+0])
		length := int(spans[2*s+1])
		end := start + length

		for i := start; i < end; i++ {
			o[i] = v1[i] | v2[i]
		}
	}
}

// OR: 3 inputs
func (o Uint64s) or3Strided(v1, v2, v3 Uint64s, strides Strides) {
	numSpans := strides[2]
	spans := strides[3:]

	for s := uint32(0); s < numSpans; s++ {
		start := int(spans[2*s+0])
		length := int(spans[2*s+1])
		end := start + length

		for i := start; i < end; i++ {
			o[i] = v1[i] | v2[i] | v3[i]
		}
	}
}

// OR: 4 inputs
func (o Uint64s) or4Strided(v1, v2, v3, v4 Uint64s, strides Strides) {
	numSpans := strides[2]
	spans := strides[3:]

	for s := uint32(0); s < numSpans; s++ {
		start := int(spans[2*s+0])
		length := int(spans[2*s+1])
		end := start + length

		for i := start; i < end; i++ {
			o[i] = v1[i] | v2[i] | v3[i] | v4[i]
		}
	}
}

// OR: 5 inputs
func (o Uint64s) or5Strided(v1, v2, v3, v4, v5 Uint64s, strides Strides) {
	numSpans := strides[2]
	spans := strides[3:]

	for s := uint32(0); s < numSpans; s++ {
		start := int(spans[2*s+0])
		length := int(spans[2*s+1])
		end := start + length

		for i := start; i < end; i++ {
			o[i] = v1[i] | v2[i] | v3[i] | v4[i] | v5[i]
		}
	}
}

// OR: 6 inputs
func (o Uint64s) or6Strided(v1, v2, v3, v4, v5, v6 Uint64s, strides Strides) {
	numSpans := strides[2]
	spans := strides[3:]

	for s := uint32(0); s < numSpans; s++ {
		start := int(spans[2*s+0])
		length := int(spans[2*s+1])
		end := start + length

		for i := start; i < end; i++ {
			o[i] = v1[i] | v2[i] | v3[i] | v4[i] | v5[i] | v6[i]
		}
	}
}
