TEXT ·_maxbits(SB), $0-24

	    MOVD in+0(FP),   R0
	    MOVD size+8(FP), R1

      CBZW R1,L1 // cbz w1, .+0x94
      WORD $0x71001c3f	 // cmp w1, #0x7
      WORD $0x2a0103e8	 // mov w8, w1
      WORD $0x54000088	 // b.hi .+0x10
      WORD $0xaa1f03e9	 // mov x9, xzr
      WORD $0x2a1f03ea	 // mov w10, wzr
      WORD $0x14000014	 // b .+0x50
      WORD $0x927d7109	 // and x9, x8, #0xfffffff8
      WORD $0x9100400a	 // add x10, x0, #0x10
      WORD $0x6f00e400	 // movi v0.2d, #0x0
      WORD $0xaa0903eb	 // mov x11, x9
      WORD $0x6f00e401	 // movi v1.2d, #0x0
      WORD $0xad7f8d42	 // ldp q2, q3, [x10,#-16]
      WORD $0xf100216b	 // subs x11, x11, #0x8
      WORD $0x9100814a	 // add x10, x10, #0x20
      WORD $0x4ea01c40	 // orr v0.16b, v2.16b, v0.16b
      WORD $0x4ea11c61	 // orr v1.16b, v3.16b, v1.16b
      WORD $0x54ffff61	 // b.ne .+0xffffffffffffffec
      WORD $0x4ea01c20	 // orr v0.16b, v1.16b, v0.16b
      WORD $0x6e004001	 // ext v1.16b, v0.16b, v0.16b, #8
      WORD $0x4ea11c00	 // orr v0.16b, v0.16b, v1.16b
      WORD $0x4e0c0401	 // dup v1.4s, v0.s[1]
      WORD $0x4ea11c00	 // orr v0.16b, v0.16b, v1.16b
      WORD $0xeb08013f	 // cmp x9, x8
      WORD $0x1e26000a	 // fmov w10, s0
      WORD $0x540000e0	 // b.eq .+0x1c
      WORD $0x8b09080b	 // add x11, x0, x9, lsl #2
      WORD $0xcb090108	 // sub x8, x8, x9
      WORD $0xb8404569	 // ldr w9, [x11],#4
      WORD $0xf1000508	 // subs x8, x8, #0x1
      WORD $0x2a0a012a	 // orr w10, w9, w10
      WORD $0x54ffffa1	 // b.ne .+0xfffffffffffffff4
      CBZW R10,L1// cbz w10, .+0x14
      WORD $0x5ac01148	 // clz w8, w10
      WORD $0x321b03e9	 // orr w9, wzr, #0x20
      WORD $0x4b080120	 // sub w0, w9, w8
      MOVD R0, ret+16(FP)
      RET
L1:   WORD $0x2a1f03e0	 // mov w0, wzr
      MOVD R0, ret+16(FP)
      RET
