package bitwise

import "github.com/viant/vec/cpu"

var allOnes = ^uint64(0)

// Strides semantics used by *andStrided* and *orStrided*:
//   - strides[0] — absolute index at which a scan starts.
//   - strides[j] (j>0) — increment applied after processing word j-1.
//
// Increment growth rules
//
//	AND : strides[j]++ when word j collapses to zero (all bits = 0).
//	OR  : strides[j]++ when word j becomes fully saturated (all bits = 1).
//
// Saturation indicates that further applications of the same bitwise
// operation cannot change the value, therefore the algorithm can safely
// skip farther on subsequent passes.
//
// We keep it as []uint32 because increments seldom exceed 32‑bit range.
type Strides []uint32

var initStrides = make([]uint32, 1024)

func init() {

	for i := 0; i < len(initStrides); i++ {
		initStrides[i] = 1
	}
}

func (s *Strides) ensureStrides(v Uint64s) {
	n := len(v) + 3
	if len(*s) == n {
		return
	}
	tmp := make(Strides, n)
	// the first two elements contain the vector length and the CPU type, neon/SVE
	tmp[0] = uint32(len(v))
	tmp[1] = uint32(cpu.Info >> 32)

	// Efficient copy from initStrides buffer
	if n <= len(initStrides) {
		copy(tmp[2:], initStrides[:n])
	} else {
		copy(tmp[2:], initStrides)
		for i := len(initStrides) + 2; i < n; i++ {
			tmp[i] = 1
		}
	}
	tmp[2] = 0
	*s = tmp
}

// And applies v1 & v2 -> o
// andStrided applies bitwise AND with dynamic stride adjustment.
//
// A stride is widened (i.e. strides[j]++) when the AND result at word j
// equals allOnes (0xFFFF_FFFF_FFFF_FFFF). The rationale is that once all
// bits are set at position j, any further AND with another vector will keep
// this word saturated, so subsequent passes can safely skip farther.
//
// NOTE: Earlier implementation widened the stride when the result collapsed
// to zero. The semantics changed so that *all bits set* triggers the
// widening instead.
func (o Uint64s) andStrided(v1, v2 Uint64s, strides Strides, count *int) {
	size := uint32(len(v1))
	i := strides[2]
	j := 2
	for i < size {
		*count++
		v := v1[i] & v2[i]
		if v == 0 {
			strides[j]++ // widen increment for NEXT visit to word j
			strides[j+1] = 1
		}
		o[i] = v
		j++
		i += strides[j] // jump by increment stored at the NEXT slot
	}
}

// Or applies v1 | v2 -> o with strided evaluation and dynamic stride adjustment.
func (o Uint64s) orStrided(v1, v2 Uint64s, strides Strides) {
	size := uint32(len(v1))
	i := strides[2]
	j := 2

	for i < size {
		v := v1[i] | v2[i]
		if v == allOnes {
			strides[j]++ // widen next stride if fully saturated
			strides[j+1] = 1
		}
		o[i] = v
		j++
		i += strides[j] // use updated stride for next jump
	}
}

func (s *Strides) fromMask(mask []uint32) {
	prev := uint32(0)
	for i := uint32(2); i < mask[2]; i++ {
		(*s)[i] = mask[i+1] - prev
		prev = mask[i+1]
	}
}
