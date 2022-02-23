TEXT ·_add_float32(SB), $0-32

	MOVD input1+0(FP),  R0
	MOVD input2+8(FP),  R1
	MOVD output+16(FP), R2
	MOVD size+24(FP),   R3

WORD $0xd342fc66 // lsr x6, x3, #2
WORD $0x710000df // cmp w6, #0x0
WORD $0x540001ad // b.le .+0x34
WORD $0x510004c5 // sub w5, w6, #0x1
WORD $0xd2800004 // mov x4, #0x0
WORD $0x910004a5 // add x5, x5, #0x1
WORD $0xd37ceca5 // lsl x5, x5, #4
WORD $0xd503201f // nop
WORD $0x3ce46800 // ldr q0, [x0,x4]
WORD $0x3ce46821 // ldr q1, [x1,x4]
WORD $0x4e21d400 // fadd v0.4s, v0.4s, v1.4s
WORD $0x3ca46840 // str q0, [x2,x4]
WORD $0x91004084 // add x4, x4, #0x10
WORD $0xeb05009f // cmp x4, x5
WORD $0x54ffff41 // b.ne .+0xffffffffffffffe8
WORD $0x531e74c6 // lsl w6, w6, #2
WORD $0x93407ccb // sxtw x11, w6
WORD $0xeb0b007f // cmp x3, x11
WORD $0x54000ee9 // b.ls .+0x1dc
WORD $0x91001167 // add x7, x11, #0x4
WORD $0xcb0b006a // sub x10, x3, x11
WORD $0xd37ef4e7 // lsl x7, x7, #2
WORD $0xd10040e9 // sub x9, x7, #0x10
WORD $0x8b070048 // add x8, x2, x7
WORD $0x8b090025 // add x5, x1, x9
WORD $0x8b070024 // add x4, x1, x7
WORD $0xeb05011f // cmp x8, x5
WORD $0x8b09004c // add x12, x2, x9
WORD $0xfa448182 // ccmp x12, x4, #0x2, hi
WORD $0x8b090004 // add x4, x0, x9
WORD $0x1a9f37e5 // cset w5, cs
WORD $0x8b070007 // add x7, x0, x7
WORD $0xeb08009f // cmp x4, x8
WORD $0xaa0903e8 // mov x8, x9
WORD $0xfa473182 // ccmp x12, x7, #0x2, cc
WORD $0x1a9f37e7 // cset w7, cs
WORD $0xf1001d5f // cmp x10, #0x7
WORD $0x0a0700a5 // and w5, w5, w7
WORD $0x1a9f97e7 // cset w7, hi
WORD $0x6a0500ff // tst w7, w5
WORD $0x54000c40 // b.eq .+0x188
WORD $0xcb440be4 // neg x4, x4, lsr #2
WORD $0xd1000467 // sub x7, x3, #0x1
WORD $0x92400485 // and x5, x4, #0x3
WORD $0xcb0b00e4 // sub x4, x7, x11
WORD $0x91000ca7 // add x7, x5, #0x3
WORD $0xeb07009f // cmp x4, x7
WORD $0x540005c3 // b.cc .+0xb8
WORD $0x2a0603ec // mov w12, w6
WORD $0xb40002e5 // cbz x5, .+0x5c
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x110004cc // add w12, w6, #0x1
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0xf10004bf // cmp x5, #0x1
WORD $0x1e212800 // fadd s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x93407d8b // sxtw x11, w12
WORD $0x540001e0 // b.eq .+0x3c
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x110008cc // add w12, w6, #0x2
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0xf1000cbf // cmp x5, #0x3
WORD $0x1e212800 // fadd s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x93407d8b // sxtw x11, w12
WORD $0x540000e1 // b.ne .+0x1c
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x11000ccc // add w12, w6, #0x3
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0x1e212800 // fadd s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x93407d8b // sxtw x11, w12
WORD $0xcb050147 // sub x7, x10, x5
WORD $0x8b050925 // add x5, x9, x5, lsl #2
WORD $0x8b05000a // add x10, x0, x5
WORD $0x8b050029 // add x9, x1, x5
WORD $0xd342fce8 // lsr x8, x7, #2
WORD $0x8b050045 // add x5, x2, x5
WORD $0xd2800004 // mov x4, #0x0
WORD $0xd2800006 // mov x6, #0x0
WORD $0x3ce46920 // ldr q0, [x9,x4]
WORD $0x910004c6 // add x6, x6, #0x1
WORD $0x3ce46941 // ldr q1, [x10,x4]
WORD $0xeb06011f // cmp x8, x6
WORD $0x4e21d400 // fadd v0.4s, v0.4s, v1.4s
WORD $0x3ca468a0 // str q0, [x5,x4]
WORD $0x91004084 // add x4, x4, #0x10
WORD $0x54ffff28 // b.hi .+0xffffffffffffffe4
WORD $0x927ef4e4 // and x4, x7, #0xfffffffffffff
WORD $0x8b04016b // add x11, x11, x4
WORD $0x0b040186 // add w6, w12, w4
WORD $0xeb0400ff // cmp x7, x4
WORD $0x540005a0 // b.eq .+0xb4
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x110004c4 // add w4, w6, #0x1
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0x93407c84 // sxtw x4, w4
WORD $0xeb03009f // cmp x4, x3
WORD $0x1e212800 // fadd s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x540004a2 // b.cs .+0x94
WORD $0xbc647800 // ldr s0, [x0,x4,lsl #2]
WORD $0x110008c5 // add w5, w6, #0x2
WORD $0xbc647821 // ldr s1, [x1,x4,lsl #2]
WORD $0x93407ca5 // sxtw x5, w5
WORD $0xeb05007f // cmp x3, x5
WORD $0x1e212800 // fadd s0, s0, s1
WORD $0xbc247840 // str s0, [x2,x4,lsl #2]
WORD $0x540003a9 // b.ls .+0x74
WORD $0xbc657800 // ldr s0, [x0,x5,lsl #2]
WORD $0x11000cc4 // add w4, w6, #0x3
WORD $0xbc657821 // ldr s1, [x1,x5,lsl #2]
WORD $0x93407c84 // sxtw x4, w4
WORD $0xeb04007f // cmp x3, x4
WORD $0x1e212800 // fadd s0, s0, s1
WORD $0xbc257840 // str s0, [x2,x5,lsl #2]
WORD $0x540002a9 // b.ls .+0x54
WORD $0xbc647800 // ldr s0, [x0,x4,lsl #2]
WORD $0x110010c5 // add w5, w6, #0x4
WORD $0xbc647821 // ldr s1, [x1,x4,lsl #2]
WORD $0x93407ca5 // sxtw x5, w5
WORD $0xeb05007f // cmp x3, x5
WORD $0x1e212800 // fadd s0, s0, s1
WORD $0xbc247840 // str s0, [x2,x4,lsl #2]
WORD $0x540001a9 // b.ls .+0x34
WORD $0xbc657800 // ldr s0, [x0,x5,lsl #2]
WORD $0x110014c6 // add w6, w6, #0x5
WORD $0xbc657821 // ldr s1, [x1,x5,lsl #2]
WORD $0x93407cc6 // sxtw x6, w6
WORD $0xeb06007f // cmp x3, x6
WORD $0x1e212800 // fadd s0, s0, s1
WORD $0xbc257840 // str s0, [x2,x5,lsl #2]
WORD $0x540000a9 // b.ls .+0x14
WORD $0xbc667800 // ldr s0, [x0,x6,lsl #2]
WORD $0xbc667821 // ldr s1, [x1,x6,lsl #2]
WORD $0x1e212800 // fadd s0, s0, s1
WORD $0xbc267840 // str s0, [x2,x6,lsl #2]
WORD $0xd65f03c0 // ret
WORD $0xd37ef463 // lsl x3, x3, #2
WORD $0xd503201f // nop
WORD $0xbc686800 // ldr s0, [x0,x8]
WORD $0xbc686821 // ldr s1, [x1,x8]
WORD $0x1e212800 // fadd s0, s0, s1
WORD $0xbc286840 // str s0, [x2,x8]
WORD $0x91001108 // add x8, x8, #0x4
WORD $0xeb03011f // cmp x8, x3
WORD $0x54ffff41 // b.ne .+0xffffffffffffffe8
WORD $0xd65f03c0 // ret

TEXT ·_sub_float32(SB), $0-32

	MOVD input1+0(FP),  R0
	MOVD input2+8(FP),  R1
	MOVD output+16(FP), R2
	MOVD size+24(FP),   R3

WORD $0xd342fc66 // lsr x6, x3, #2
WORD $0x710000df // cmp w6, #0x0
WORD $0x540001ad // b.le .+0x34
WORD $0x510004c5 // sub w5, w6, #0x1
WORD $0xd2800004 // mov x4, #0x0
WORD $0x910004a5 // add x5, x5, #0x1
WORD $0xd37ceca5 // lsl x5, x5, #4
WORD $0xd503201f // nop
WORD $0x3ce46800 // ldr q0, [x0,x4]
WORD $0x3ce46821 // ldr q1, [x1,x4]
WORD $0x4ea1d400 // fsub v0.4s, v0.4s, v1.4s
WORD $0x3ca46840 // str q0, [x2,x4]
WORD $0x91004084 // add x4, x4, #0x10
WORD $0xeb05009f // cmp x4, x5
WORD $0x54ffff41 // b.ne .+0xffffffffffffffe8
WORD $0x531e74c6 // lsl w6, w6, #2
WORD $0x93407ccb // sxtw x11, w6
WORD $0xeb0b007f // cmp x3, x11
WORD $0x54000ee9 // b.ls .+0x1dc
WORD $0x91001167 // add x7, x11, #0x4
WORD $0xcb0b006a // sub x10, x3, x11
WORD $0xd37ef4e7 // lsl x7, x7, #2
WORD $0xd10040e9 // sub x9, x7, #0x10
WORD $0x8b070048 // add x8, x2, x7
WORD $0x8b090025 // add x5, x1, x9
WORD $0x8b070024 // add x4, x1, x7
WORD $0xeb05011f // cmp x8, x5
WORD $0x8b09004c // add x12, x2, x9
WORD $0xfa448182 // ccmp x12, x4, #0x2, hi
WORD $0x8b090004 // add x4, x0, x9
WORD $0x1a9f37e5 // cset w5, cs
WORD $0x8b070007 // add x7, x0, x7
WORD $0xeb08009f // cmp x4, x8
WORD $0xaa0903e8 // mov x8, x9
WORD $0xfa473182 // ccmp x12, x7, #0x2, cc
WORD $0x1a9f37e7 // cset w7, cs
WORD $0xf1001d5f // cmp x10, #0x7
WORD $0x0a0700a5 // and w5, w5, w7
WORD $0x1a9f97e7 // cset w7, hi
WORD $0x6a0500ff // tst w7, w5
WORD $0x54000c40 // b.eq .+0x188
WORD $0xcb440be4 // neg x4, x4, lsr #2
WORD $0xd1000467 // sub x7, x3, #0x1
WORD $0x92400485 // and x5, x4, #0x3
WORD $0xcb0b00e4 // sub x4, x7, x11
WORD $0x91000ca7 // add x7, x5, #0x3
WORD $0xeb07009f // cmp x4, x7
WORD $0x540005c3 // b.cc .+0xb8
WORD $0x2a0603ec // mov w12, w6
WORD $0xb40002e5 // cbz x5, .+0x5c
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x110004cc // add w12, w6, #0x1
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0xf10004bf // cmp x5, #0x1
WORD $0x1e213800 // fsub s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x93407d8b // sxtw x11, w12
WORD $0x540001e0 // b.eq .+0x3c
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x110008cc // add w12, w6, #0x2
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0xf1000cbf // cmp x5, #0x3
WORD $0x1e213800 // fsub s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x93407d8b // sxtw x11, w12
WORD $0x540000e1 // b.ne .+0x1c
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x11000ccc // add w12, w6, #0x3
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0x1e213800 // fsub s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x93407d8b // sxtw x11, w12
WORD $0xcb050147 // sub x7, x10, x5
WORD $0x8b050925 // add x5, x9, x5, lsl #2
WORD $0x8b05000a // add x10, x0, x5
WORD $0x8b050029 // add x9, x1, x5
WORD $0xd342fce8 // lsr x8, x7, #2
WORD $0x8b050045 // add x5, x2, x5
WORD $0xd2800004 // mov x4, #0x0
WORD $0xd2800006 // mov x6, #0x0
WORD $0x3ce46921 // ldr q1, [x9,x4]
WORD $0x910004c6 // add x6, x6, #0x1
WORD $0x3ce46940 // ldr q0, [x10,x4]
WORD $0xeb06011f // cmp x8, x6
WORD $0x4ea1d400 // fsub v0.4s, v0.4s, v1.4s
WORD $0x3ca468a0 // str q0, [x5,x4]
WORD $0x91004084 // add x4, x4, #0x10
WORD $0x54ffff28 // b.hi .+0xffffffffffffffe4
WORD $0x927ef4e4 // and x4, x7, #0xfffffffffffff
WORD $0x8b04016b // add x11, x11, x4
WORD $0x0b040186 // add w6, w12, w4
WORD $0xeb0400ff // cmp x7, x4
WORD $0x540005a0 // b.eq .+0xb4
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x110004c4 // add w4, w6, #0x1
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0x93407c84 // sxtw x4, w4
WORD $0xeb03009f // cmp x4, x3
WORD $0x1e213800 // fsub s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x540004a2 // b.cs .+0x94
WORD $0xbc647800 // ldr s0, [x0,x4,lsl #2]
WORD $0x110008c5 // add w5, w6, #0x2
WORD $0xbc647821 // ldr s1, [x1,x4,lsl #2]
WORD $0x93407ca5 // sxtw x5, w5
WORD $0xeb05007f // cmp x3, x5
WORD $0x1e213800 // fsub s0, s0, s1
WORD $0xbc247840 // str s0, [x2,x4,lsl #2]
WORD $0x540003a9 // b.ls .+0x74
WORD $0xbc657800 // ldr s0, [x0,x5,lsl #2]
WORD $0x11000cc4 // add w4, w6, #0x3
WORD $0xbc657821 // ldr s1, [x1,x5,lsl #2]
WORD $0x93407c84 // sxtw x4, w4
WORD $0xeb04007f // cmp x3, x4
WORD $0x1e213800 // fsub s0, s0, s1
WORD $0xbc257840 // str s0, [x2,x5,lsl #2]
WORD $0x540002a9 // b.ls .+0x54
WORD $0xbc647800 // ldr s0, [x0,x4,lsl #2]
WORD $0x110010c5 // add w5, w6, #0x4
WORD $0xbc647821 // ldr s1, [x1,x4,lsl #2]
WORD $0x93407ca5 // sxtw x5, w5
WORD $0xeb05007f // cmp x3, x5
WORD $0x1e213800 // fsub s0, s0, s1
WORD $0xbc247840 // str s0, [x2,x4,lsl #2]
WORD $0x540001a9 // b.ls .+0x34
WORD $0xbc657800 // ldr s0, [x0,x5,lsl #2]
WORD $0x110014c6 // add w6, w6, #0x5
WORD $0xbc657821 // ldr s1, [x1,x5,lsl #2]
WORD $0x93407cc6 // sxtw x6, w6
WORD $0xeb06007f // cmp x3, x6
WORD $0x1e213800 // fsub s0, s0, s1
WORD $0xbc257840 // str s0, [x2,x5,lsl #2]
WORD $0x540000a9 // b.ls .+0x14
WORD $0xbc667800 // ldr s0, [x0,x6,lsl #2]
WORD $0xbc667821 // ldr s1, [x1,x6,lsl #2]
WORD $0x1e213800 // fsub s0, s0, s1
WORD $0xbc267840 // str s0, [x2,x6,lsl #2]
WORD $0xd65f03c0 // ret
WORD $0xd37ef463 // lsl x3, x3, #2
WORD $0xd503201f // nop
WORD $0xbc686800 // ldr s0, [x0,x8]
WORD $0xbc686821 // ldr s1, [x1,x8]
WORD $0x1e213800 // fsub s0, s0, s1
WORD $0xbc286840 // str s0, [x2,x8]
WORD $0x91001108 // add x8, x8, #0x4
WORD $0xeb03011f // cmp x8, x3
WORD $0x54ffff41 // b.ne .+0xffffffffffffffe8
WORD $0xd65f03c0 // ret

TEXT ·_mul_float32(SB), $0-32

	MOVD input1+0(FP),  R0
	MOVD input2+8(FP),  R1
	MOVD output+16(FP), R2
	MOVD size+24(FP),   R3

WORD $0xd342fc66 // lsr x6, x3, #2
WORD $0x710000df // cmp w6, #0x0
WORD $0x540001ad // b.le .+0x34
WORD $0x510004c5 // sub w5, w6, #0x1
WORD $0xd2800004 // mov x4, #0x0
WORD $0x910004a5 // add x5, x5, #0x1
WORD $0xd37ceca5 // lsl x5, x5, #4
WORD $0xd503201f // nop
WORD $0x3ce46800 // ldr q0, [x0,x4]
WORD $0x3ce46821 // ldr q1, [x1,x4]
WORD $0x6e21dc00 // fmul v0.4s, v0.4s, v1.4s
WORD $0x3ca46840 // str q0, [x2,x4]
WORD $0x91004084 // add x4, x4, #0x10
WORD $0xeb05009f // cmp x4, x5
WORD $0x54ffff41 // b.ne .+0xffffffffffffffe8
WORD $0x531e74c6 // lsl w6, w6, #2
WORD $0x93407ccb // sxtw x11, w6
WORD $0xeb0b007f // cmp x3, x11
WORD $0x54000ee9 // b.ls .+0x1dc
WORD $0x91001167 // add x7, x11, #0x4
WORD $0xcb0b006a // sub x10, x3, x11
WORD $0xd37ef4e7 // lsl x7, x7, #2
WORD $0xd10040e9 // sub x9, x7, #0x10
WORD $0x8b070048 // add x8, x2, x7
WORD $0x8b090025 // add x5, x1, x9
WORD $0x8b070024 // add x4, x1, x7
WORD $0xeb05011f // cmp x8, x5
WORD $0x8b09004c // add x12, x2, x9
WORD $0xfa448182 // ccmp x12, x4, #0x2, hi
WORD $0x8b090004 // add x4, x0, x9
WORD $0x1a9f37e5 // cset w5, cs
WORD $0x8b070007 // add x7, x0, x7
WORD $0xeb08009f // cmp x4, x8
WORD $0xaa0903e8 // mov x8, x9
WORD $0xfa473182 // ccmp x12, x7, #0x2, cc
WORD $0x1a9f37e7 // cset w7, cs
WORD $0xf1001d5f // cmp x10, #0x7
WORD $0x0a0700a5 // and w5, w5, w7
WORD $0x1a9f97e7 // cset w7, hi
WORD $0x6a0500ff // tst w7, w5
WORD $0x54000c40 // b.eq .+0x188
WORD $0xcb440be4 // neg x4, x4, lsr #2
WORD $0xd1000467 // sub x7, x3, #0x1
WORD $0x92400485 // and x5, x4, #0x3
WORD $0xcb0b00e4 // sub x4, x7, x11
WORD $0x91000ca7 // add x7, x5, #0x3
WORD $0xeb07009f // cmp x4, x7
WORD $0x540005c3 // b.cc .+0xb8
WORD $0x2a0603ec // mov w12, w6
WORD $0xb40002e5 // cbz x5, .+0x5c
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x110004cc // add w12, w6, #0x1
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0xf10004bf // cmp x5, #0x1
WORD $0x1e210800 // fmul s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x93407d8b // sxtw x11, w12
WORD $0x540001e0 // b.eq .+0x3c
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x110008cc // add w12, w6, #0x2
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0xf1000cbf // cmp x5, #0x3
WORD $0x1e210800 // fmul s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x93407d8b // sxtw x11, w12
WORD $0x540000e1 // b.ne .+0x1c
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x11000ccc // add w12, w6, #0x3
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0x1e210800 // fmul s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x93407d8b // sxtw x11, w12
WORD $0xcb050147 // sub x7, x10, x5
WORD $0x8b050925 // add x5, x9, x5, lsl #2
WORD $0x8b05000a // add x10, x0, x5
WORD $0x8b050029 // add x9, x1, x5
WORD $0xd342fce8 // lsr x8, x7, #2
WORD $0x8b050045 // add x5, x2, x5
WORD $0xd2800004 // mov x4, #0x0
WORD $0xd2800006 // mov x6, #0x0
WORD $0x3ce46920 // ldr q0, [x9,x4]
WORD $0x910004c6 // add x6, x6, #0x1
WORD $0x3ce46941 // ldr q1, [x10,x4]
WORD $0xeb06011f // cmp x8, x6
WORD $0x6e21dc00 // fmul v0.4s, v0.4s, v1.4s
WORD $0x3ca468a0 // str q0, [x5,x4]
WORD $0x91004084 // add x4, x4, #0x10
WORD $0x54ffff28 // b.hi .+0xffffffffffffffe4
WORD $0x927ef4e4 // and x4, x7, #0xfffffffffffff
WORD $0x8b04016b // add x11, x11, x4
WORD $0x0b040186 // add w6, w12, w4
WORD $0xeb0400ff // cmp x7, x4
WORD $0x540005a0 // b.eq .+0xb4
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x110004c4 // add w4, w6, #0x1
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0x93407c84 // sxtw x4, w4
WORD $0xeb03009f // cmp x4, x3
WORD $0x1e210800 // fmul s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x540004a2 // b.cs .+0x94
WORD $0xbc647800 // ldr s0, [x0,x4,lsl #2]
WORD $0x110008c5 // add w5, w6, #0x2
WORD $0xbc647821 // ldr s1, [x1,x4,lsl #2]
WORD $0x93407ca5 // sxtw x5, w5
WORD $0xeb05007f // cmp x3, x5
WORD $0x1e210800 // fmul s0, s0, s1
WORD $0xbc247840 // str s0, [x2,x4,lsl #2]
WORD $0x540003a9 // b.ls .+0x74
WORD $0xbc657800 // ldr s0, [x0,x5,lsl #2]
WORD $0x11000cc4 // add w4, w6, #0x3
WORD $0xbc657821 // ldr s1, [x1,x5,lsl #2]
WORD $0x93407c84 // sxtw x4, w4
WORD $0xeb04007f // cmp x3, x4
WORD $0x1e210800 // fmul s0, s0, s1
WORD $0xbc257840 // str s0, [x2,x5,lsl #2]
WORD $0x540002a9 // b.ls .+0x54
WORD $0xbc647800 // ldr s0, [x0,x4,lsl #2]
WORD $0x110010c5 // add w5, w6, #0x4
WORD $0xbc647821 // ldr s1, [x1,x4,lsl #2]
WORD $0x93407ca5 // sxtw x5, w5
WORD $0xeb05007f // cmp x3, x5
WORD $0x1e210800 // fmul s0, s0, s1
WORD $0xbc247840 // str s0, [x2,x4,lsl #2]
WORD $0x540001a9 // b.ls .+0x34
WORD $0xbc657800 // ldr s0, [x0,x5,lsl #2]
WORD $0x110014c6 // add w6, w6, #0x5
WORD $0xbc657821 // ldr s1, [x1,x5,lsl #2]
WORD $0x93407cc6 // sxtw x6, w6
WORD $0xeb06007f // cmp x3, x6
WORD $0x1e210800 // fmul s0, s0, s1
WORD $0xbc257840 // str s0, [x2,x5,lsl #2]
WORD $0x540000a9 // b.ls .+0x14
WORD $0xbc667800 // ldr s0, [x0,x6,lsl #2]
WORD $0xbc667821 // ldr s1, [x1,x6,lsl #2]
WORD $0x1e210800 // fmul s0, s0, s1
WORD $0xbc267840 // str s0, [x2,x6,lsl #2]
WORD $0xd65f03c0 // ret
WORD $0xd37ef463 // lsl x3, x3, #2
WORD $0xd503201f // nop
WORD $0xbc686800 // ldr s0, [x0,x8]
WORD $0xbc686821 // ldr s1, [x1,x8]
WORD $0x1e210800 // fmul s0, s0, s1
WORD $0xbc286840 // str s0, [x2,x8]
WORD $0x91001108 // add x8, x8, #0x4
WORD $0xeb03011f // cmp x8, x3
WORD $0x54ffff41 // b.ne .+0xffffffffffffffe8
WORD $0xd65f03c0 // ret

TEXT ·_div_float32(SB), $0-32

	MOVD input1+0(FP),  R0
	MOVD input2+8(FP),  R1
	MOVD output+16(FP), R2
	MOVD size+24(FP),   R3

WORD $0xd342fc66 // lsr x6, x3, #2
WORD $0x710000df // cmp w6, #0x0
WORD $0x540001ad // b.le .+0x34
WORD $0x510004c5 // sub w5, w6, #0x1
WORD $0xd2800004 // mov x4, #0x0
WORD $0x910004a5 // add x5, x5, #0x1
WORD $0xd37ceca5 // lsl x5, x5, #4
WORD $0xd503201f // nop
WORD $0x3ce46800 // ldr q0, [x0,x4]
WORD $0x3ce46821 // ldr q1, [x1,x4]
WORD $0x6e21fc00 // fdiv v0.4s, v0.4s, v1.4s
WORD $0x3ca46840 // str q0, [x2,x4]
WORD $0x91004084 // add x4, x4, #0x10
WORD $0xeb05009f // cmp x4, x5
WORD $0x54ffff41 // b.ne .+0xffffffffffffffe8
WORD $0x531e74c6 // lsl w6, w6, #2
WORD $0x93407ccb // sxtw x11, w6
WORD $0xeb0b007f // cmp x3, x11
WORD $0x54000ee9 // b.ls .+0x1dc
WORD $0x91001167 // add x7, x11, #0x4
WORD $0xcb0b006a // sub x10, x3, x11
WORD $0xd37ef4e7 // lsl x7, x7, #2
WORD $0xd10040e9 // sub x9, x7, #0x10
WORD $0x8b070048 // add x8, x2, x7
WORD $0x8b090025 // add x5, x1, x9
WORD $0x8b070024 // add x4, x1, x7
WORD $0xeb05011f // cmp x8, x5
WORD $0x8b09004c // add x12, x2, x9
WORD $0xfa448182 // ccmp x12, x4, #0x2, hi
WORD $0x8b090004 // add x4, x0, x9
WORD $0x1a9f37e5 // cset w5, cs
WORD $0x8b070007 // add x7, x0, x7
WORD $0xeb08009f // cmp x4, x8
WORD $0xaa0903e8 // mov x8, x9
WORD $0xfa473182 // ccmp x12, x7, #0x2, cc
WORD $0x1a9f37e7 // cset w7, cs
WORD $0xf1001d5f // cmp x10, #0x7
WORD $0x0a0700a5 // and w5, w5, w7
WORD $0x1a9f97e7 // cset w7, hi
WORD $0x6a0500ff // tst w7, w5
WORD $0x54000c40 // b.eq .+0x188
WORD $0xcb440be4 // neg x4, x4, lsr #2
WORD $0xd1000467 // sub x7, x3, #0x1
WORD $0x92400485 // and x5, x4, #0x3
WORD $0xcb0b00e4 // sub x4, x7, x11
WORD $0x91000ca7 // add x7, x5, #0x3
WORD $0xeb07009f // cmp x4, x7
WORD $0x540005c3 // b.cc .+0xb8
WORD $0x2a0603ec // mov w12, w6
WORD $0xb40002e5 // cbz x5, .+0x5c
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0x110004cc // add w12, w6, #0x1
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0xf10004bf // cmp x5, #0x1
WORD $0x1e211800 // fdiv s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x93407d8b // sxtw x11, w12
WORD $0x540001e0 // b.eq .+0x3c
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0x110008cc // add w12, w6, #0x2
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0xf1000cbf // cmp x5, #0x3
WORD $0x1e211800 // fdiv s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x93407d8b // sxtw x11, w12
WORD $0x540000e1 // b.ne .+0x1c
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0x11000ccc // add w12, w6, #0x3
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x1e211800 // fdiv s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x93407d8b // sxtw x11, w12
WORD $0xcb050147 // sub x7, x10, x5
WORD $0x8b050925 // add x5, x9, x5, lsl #2
WORD $0x8b05000a // add x10, x0, x5
WORD $0x8b050029 // add x9, x1, x5
WORD $0xd342fce8 // lsr x8, x7, #2
WORD $0x8b050045 // add x5, x2, x5
WORD $0xd2800004 // mov x4, #0x0
WORD $0xd2800006 // mov x6, #0x0
WORD $0x3ce46921 // ldr q1, [x9,x4]
WORD $0x910004c6 // add x6, x6, #0x1
WORD $0x3ce46940 // ldr q0, [x10,x4]
WORD $0xeb06011f // cmp x8, x6
WORD $0x6e21fc00 // fdiv v0.4s, v0.4s, v1.4s
WORD $0x3ca468a0 // str q0, [x5,x4]
WORD $0x91004084 // add x4, x4, #0x10
WORD $0x54ffff28 // b.hi .+0xffffffffffffffe4
WORD $0x927ef4e4 // and x4, x7, #0xfffffffffffff
WORD $0x8b04016b // add x11, x11, x4
WORD $0x0b040186 // add w6, w12, w4
WORD $0xeb0400ff // cmp x7, x4
WORD $0x540005a0 // b.eq .+0xb4
WORD $0xbc6b7821 // ldr s1, [x1,x11,lsl #2]
WORD $0x110004c4 // add w4, w6, #0x1
WORD $0xbc6b7800 // ldr s0, [x0,x11,lsl #2]
WORD $0x93407c84 // sxtw x4, w4
WORD $0xeb03009f // cmp x4, x3
WORD $0x1e211800 // fdiv s0, s0, s1
WORD $0xbc2b7840 // str s0, [x2,x11,lsl #2]
WORD $0x540004a2 // b.cs .+0x94
WORD $0xbc647821 // ldr s1, [x1,x4,lsl #2]
WORD $0x110008c5 // add w5, w6, #0x2
WORD $0xbc647800 // ldr s0, [x0,x4,lsl #2]
WORD $0x93407ca5 // sxtw x5, w5
WORD $0xeb05007f // cmp x3, x5
WORD $0x1e211800 // fdiv s0, s0, s1
WORD $0xbc247840 // str s0, [x2,x4,lsl #2]
WORD $0x540003a9 // b.ls .+0x74
WORD $0xbc657821 // ldr s1, [x1,x5,lsl #2]
WORD $0x11000cc4 // add w4, w6, #0x3
WORD $0xbc657800 // ldr s0, [x0,x5,lsl #2]
WORD $0x93407c84 // sxtw x4, w4
WORD $0xeb04007f // cmp x3, x4
WORD $0x1e211800 // fdiv s0, s0, s1
WORD $0xbc257840 // str s0, [x2,x5,lsl #2]
WORD $0x540002a9 // b.ls .+0x54
WORD $0xbc647821 // ldr s1, [x1,x4,lsl #2]
WORD $0x110010c5 // add w5, w6, #0x4
WORD $0xbc647800 // ldr s0, [x0,x4,lsl #2]
WORD $0x93407ca5 // sxtw x5, w5
WORD $0xeb05007f // cmp x3, x5
WORD $0x1e211800 // fdiv s0, s0, s1
WORD $0xbc247840 // str s0, [x2,x4,lsl #2]
WORD $0x540001a9 // b.ls .+0x34
WORD $0xbc657821 // ldr s1, [x1,x5,lsl #2]
WORD $0x110014c6 // add w6, w6, #0x5
WORD $0xbc657800 // ldr s0, [x0,x5,lsl #2]
WORD $0x93407cc6 // sxtw x6, w6
WORD $0xeb06007f // cmp x3, x6
WORD $0x1e211800 // fdiv s0, s0, s1
WORD $0xbc257840 // str s0, [x2,x5,lsl #2]
WORD $0x540000a9 // b.ls .+0x14
WORD $0xbc667821 // ldr s1, [x1,x6,lsl #2]
WORD $0xbc667800 // ldr s0, [x0,x6,lsl #2]
WORD $0x1e211800 // fdiv s0, s0, s1
WORD $0xbc267840 // str s0, [x2,x6,lsl #2]
WORD $0xd65f03c0 // ret
WORD $0xd37ef463 // lsl x3, x3, #2
WORD $0xd503201f // nop
WORD $0xbc686821 // ldr s1, [x1,x8]
WORD $0xbc686800 // ldr s0, [x0,x8]
WORD $0x1e211800 // fdiv s0, s0, s1
WORD $0xbc286840 // str s0, [x2,x8]
WORD $0x91001108 // add x8, x8, #0x4
WORD $0xeb03011f // cmp x8, x3
WORD $0x54ffff41 // b.ne .+0xffffffffffffffe8
WORD $0xd65f03c0 // ret
