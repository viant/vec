package bitwise

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

// ensureStrides makes sure s has at least len(v)+1 entries,
// initializing every entry to 1 except the first.
func (s *Strides) ensureStrides(v Uint64s) {
	n := len(v) + 1
	if len(*s) == n {
		return
	}
	tmp := make(Strides, n)

	// Efficient copy from initStrides buffer
	if n <= len(initStrides) {
		copy(tmp, initStrides[:n])
	} else {
		copy(tmp, initStrides)
		for i := len(initStrides); i < n; i++ {
			tmp[i] = 1
		}
	}
	tmp[0] = 0
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
func (o Uint64s) andStrided(v1, v2 Uint64s, strides Strides) {
	size := uint32(len(v1))
	i := strides[0]
	j := uint32(0)
	for i < size {
		v := v1[i] & v2[i]
		if v == 0 {
			strides[j]++ // widen increment for NEXT visit to word j
		}
		o[i] = v
		j++
		i += strides[j] // jump by increment stored at the NEXT slot
	}
}

// Or applies v1 | v2 -> o with strided evaluation and dynamic stride adjustment.
func (o Uint64s) orStrided(v1, v2 Uint64s, strides Strides) {
	size := uint32(len(v1))
	i := strides[0]
	j := uint32(0)

	for i < size {
		v := v1[i] | v2[i]
		if v == allOnes {
			strides[j]++ // widen next stride if fully saturated
		}
		o[i] = v
		j++
		i += strides[j] // use updated stride for next jump
	}
}
