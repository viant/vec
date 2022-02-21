//go:build (!amd64 && !arm64) || noasm
// +build !amd64,!arm64 noasm

package pack

func (o Uint64s) Expand(input uint64) {
	o.expand(input)
}
