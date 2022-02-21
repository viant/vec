package cpu

import (
	"golang.org/x/sys/cpu"
)

var tryAVX512 = false

//TryAVX512 would try use AVX512 if cpu support it, AVX512 control is introduced since early CPU with AVX512 slowed down CPU clock
//resulting in substantially worse overall performance comparing just to AVX2
func TryAVX512(flag bool) {
	tryAVX512 = flag
}

var Info = 0

const (
	useAVX2   = 1 << 32
	useAVX512 = 2 << 32
	useSVE    = 3 << 32
)

func init() {

	if tryAVX512 && cpu.X86.HasAVX512 {
		Info = useAVX512
	} else if cpu.X86.HasAVX2 {
		Info = useAVX2
	} else if cpu.ARM64.HasSVE {
		Info = useSVE
	}
}
