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
