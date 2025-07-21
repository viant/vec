package bitwise

// Uint64s defines Uint64s
type Uint64s []uint64

// And applies v1 & v2 -> o
func (o Uint64s) and(v1, v2 Uint64s) {
	for i, v := range v1 {
		o[i] = v & v2[i]
	}
}

// andV3 applies v1 & v2 & v3 -> o
func (o Uint64s) andV3(v1, v2, v3 Uint64s) {
	for i, v := range v1 {
		o[i] = v & v2[i] & v3[i]
	}
}

// andV4 applies v1 & v2 & v3 -> o
func (o Uint64s) andV4(v1, v2, v3, v4 Uint64s) {
	for i, v := range v1 {
		o[i] = v & v2[i] & v3[i]
	}
}

// andV5 applies v1 & v2 -> o
func (o Uint64s) andV5(v1, v2, v3, v4, v5 Uint64s) {
	for i, v := range v1 {
		o[i] = v & v2[i] & v3[i] & v4[i] & v5[i]
	}
}

// or applies v1 | v2 -> o
func (o Uint64s) or(v1, v2 Uint64s) {
	for i, v := range v1 {
		o[i] = v | v2[i]
	}
}

// orV3 applies v1 | v2 | v3 -> o
func (o Uint64s) orV3(v1, v2, v3 Uint64s) {
	for i, v := range v1 {
		o[i] = v | v2[i] | v3[i]
	}
}

// orV4 applies v1 | v2 | v3 | v4-> o
func (o Uint64s) orV4(v1, v2, v3, v4 Uint64s) {
	for i, v := range v1 {
		o[i] = v | v2[i] | v3[i] | v4[i]
	}
}

// orV5 applies v1 | v2 | v3 | v4 | v5 -> o
func (o Uint64s) orV5(v1, v2, v3, v4, v5 Uint64s) {
	for i, v := range v1 {
		o[i] = v | v2[i] | v3[i] | v4[i] | v5[i]
	}
}

// orV6 applies v1 | v2 | v3 | v4 | v5 -> o
func (o Uint64s) orV6(v1, v2, v3, v4, v5, v6 Uint64s) {
	for i, v := range v1 {
		o[i] = v | v2[i] | v3[i] | v4[i] | v5[i] | v6[i]
	}
}

// And applies v1 ^ v2 -> o
func (o Uint64s) xor(v1, v2 Uint64s) {
	for i, v := range v1 {
		o[i] = v ^ v2[i]
	}
}
