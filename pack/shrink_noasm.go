//go:build (!amd64 && !arm64) || noasm
// +build !amd64,!arm64 noasm

package pack

func (u *Uint64) Shrink(value byte, vec []uint64) {
	o.shrink(input)
}

func (u *Uint64) ShrinkValue(value byte, vec []uint64) {
	u.shrinkValue(value, input)
}
