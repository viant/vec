TEXT ·_hsum_float32(SB), $0-32

	MOVD input+0(FP),   R0
	MOVD result+8(FP),  R1
	MOVD size+16(FP),   R2

WORD $0xb900003f	 // str wzr, [x1]
WORD $0xd342fc43	 // lsr x3, x2, #2
WORD $0x7100007f	 // cmp w3, #0x0
WORD $0x5400036d	 // b.le .+0x6c
WORD $0x51000465	 // sub w5, w3, #0x1
WORD $0x91004007	 // add x7, x0, #0x10
WORD $0x2a0503e6	 // mov w6, w5
WORD $0x0f000401	 // movi v1.2s, #0x0
WORD $0xaa0003e4	 // mov x4, x0
WORD $0x8b2650e5	 // add x5, x7, w6, uxtw #4
WORD $0x3cc10480	 // ldr q0, [x4],#16
WORD $0x6e20d400	 // faddp v0.4s, v0.4s, v0.4s
WORD $0xeb0400bf	 // cmp x5, x4
WORD $0x6e20d400	 // faddp v0.4s, v0.4s, v0.4s
WORD $0x1e202821	 // fadd s1, s1, s0
WORD $0xbd000021	 // str s1, [x1]
WORD $0x54ffff41	 // b.ne .+0xffffffffffffffe8
WORD $0x531e7463	 // lsl w3, w3, #2
WORD $0x93407c63	 // sxtw x3, w3
WORD $0xeb03005f	 // cmp x2, x3
WORD $0x54000129	 // b.ls .+0x24
WORD $0x8b030803	 // add x3, x0, x3, lsl #2
WORD $0x8b020800	 // add x0, x0, x2, lsl #2
WORD $0xd503201f	 // nop
WORD $0xbc404460	 // ldr s0, [x3],#4
WORD $0x1e202821	 // fadd s1, s1, s0
WORD $0xeb03001f	 // cmp x0, x3
WORD $0xbd000021	 // str s1, [x1]
WORD $0x54ffff81	 // b.ne .+0xfffffffffffffff0
WORD $0xd65f03c0	 // ret
WORD $0x0f000401	 // movi v1.2s, #0x0
WORD $0x17fffff2	 // b .+0xffffffffffffffc8
