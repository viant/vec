package bitwise

// Strides semantics used by *addStrided* and *AndStrided*:
//   - strides[0] — absolute index at which a scan starts.
//   - strides[j] (j>0) — increment applied after processing word j‑1.
//     The increment grows whenever word j‑1 collapses to zero, causing
//     subsequent scans to skip further.
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
func (o Uint64s) addStrided(v1, v2 Uint64s, strides Strides) {
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
