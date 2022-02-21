//go:build (!amd64 && !arm64) || noasm
// +build !amd64,!arm64 noasm

package bitwise

//And applies v1 & v2 -> o
func (o Uint64s) And(v1, v2 Uint64s) {
	o.and(v1, v2)
}

//And applies v1 & v2 & v3 -> o
func (o Uint64s) AndV3(v1, v2, v3 Uint64s) {
	o.andV3(v1, v2, v3)
}

//And applies v1 & v2 & v3 & v4 -> o
func (o Uint64s) AndV4(v1, v2, v3, v4 Uint64s) {
	o.andV3(v1, v2, v3, v4)
}

//And applies v1 & v2 & v3 & v4 & v5 -> o
func (o Uint64s) AndV5(v1, v2, v3, v4, v5 Uint64s) {
	o.andV3(v1, v2, v3, v4, v5)
}

//Or applies v1 | v2 -> o
func (o Uint64s) Or(v1, v2 Uint64s) {
	o.or(v1, v2)
}

//OrV3 applies v1 | v2 | v3 -> o
func (o Uint64s) OrV3(v1, v2, v3 Uint64s) {
	o.orV3(v1, v2, v3)
}

//OrV4 applies v1 | v2 | v3 | v4-> o
func (o Uint64s) OrV4(v1, v2, v3, v4 Uint64s) {
	o.orV4(v1, v2, v3, v4)
}

//OrV5 applies v1 | v2 | v3 | v4 | v5 -> o
func (o Uint64s) OrV5(v1, v2, v3, v4, v5 Uint64s) {
	o.orV5(v1, v2, v3, v4, v5)
}

//OrV6 applies v1 | v2 | v3 | v4 | v5 | v6 -> o
func (o Uint64s) OrV6(v1, v2, v3, v4, v5, v6 Uint64s) {
	o.orV6(v1, v2, v3, v4, v5, v6)
}

//Xor applies v1 ^ v2 -> o
func (o Uint64s) Xor(v1, v2 Uint64s) {
	return o.xor(v1, v2)
}

//XorV3 applies v1 ^ v2 ^ v3 -> o
func (o Uint64s) XorV3(v1, v2, v3 Uint64s) {
	return o.xorV3(v1, v2, v3)
}
