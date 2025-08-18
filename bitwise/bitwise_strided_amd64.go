package bitwise

func (s Strides) SetActiveStrides(set Uint64s) {
	s.setActiveStrides(set)
}

// =======================
// AND
// =======================

// AND: 2 inputs (exported wrapper)
func (o Uint64s) AndStrided(v1, v2 Uint64s, strides Strides) {
	o.andStrided(v1, v2, strides)
}

// AND: 3–6 inputs (exported wrappers)
func (o Uint64s) And3Strided(v1, v2, v3 Uint64s, strides Strides) {
	o.and3Strided(v1, v2, v3, strides)
}

func (o Uint64s) And4Strided(v1, v2, v3, v4 Uint64s, strides Strides) {
	o.and4Strided(v1, v2, v3, v4, strides)
}

func (o Uint64s) And5Strided(v1, v2, v3, v4, v5 Uint64s, strides Strides) {
	o.and5Strided(v1, v2, v3, v4, v5, strides)
}

func (o Uint64s) And6Strided(v1, v2, v3, v4, v5, v6 Uint64s, strides Strides) {
	o.and6Strided(v1, v2, v3, v4, v5, v6, strides)
}

// =======================
// OR
// =======================

// OR: 2 inputs (exported wrapper)
func (o Uint64s) OrStrided(v1, v2 Uint64s, strides Strides) {
	o.orStrided(v1, v2, strides)
}

// OR: 3–6 inputs (exported wrappers)
func (o Uint64s) Or3Strided(v1, v2, v3 Uint64s, strides Strides) {
	o.or3Strided(v1, v2, v3, strides)
}

func (o Uint64s) Or4Strided(v1, v2, v3, v4 Uint64s, strides Strides) {
	o.or4Strided(v1, v2, v3, v4, strides)
}

func (o Uint64s) Or5Strided(v1, v2, v3, v4, v5 Uint64s, strides Strides) {
	o.or5Strided(v1, v2, v3, v4, v5, strides)
}

func (o Uint64s) Or6Strided(v1, v2, v3, v4, v5, v6 Uint64s, strides Strides) {
	o.or6Strided(v1, v2, v3, v4, v5, v6, strides)
}
