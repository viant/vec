//go:build (!amd64 && !arm64) || noasm
// +build !amd64,!arm64 noasm

package lut

func Lookup(input, table []byte, output *[]byte) {
	lookup(input, table, output)
}

func LookupSubrange(input, range_upper, table []byte, output *[]byte) {
	lookupSubrange(input, range_upper, table, output)
}
