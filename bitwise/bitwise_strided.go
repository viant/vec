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

func (o Uint64s) andStrided(v1, v2 Uint64s, strides Strides) {
	if len(strides) < 3 {
		return
	}

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

func (o Uint64s) orStrided(v1, v2 Uint64s, strides Strides) {
	if len(strides) < 3 {
		return
	}

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
