//go:build (!amd64 && !arm64) || noasm
// +build !amd64,!arm64 noasm

package bits

func (p *Positions) Populate(input uint64) uint8 {
	return p.populate(input)
}
