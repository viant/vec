//go:build (!amd64 && !arm64) || noasm
// +build !amd64,!arm64 noasm

package pack

func Count(input uint64) {
	count(input)
}
