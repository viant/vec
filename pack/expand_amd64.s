//+build !noasm !appengine
// AUTO-GENERATED BY C2GOASM -- DO NOT EDIT

DATA LCDATA1<>+0x000(SB)/8, $0x0000000000000000
DATA LCDATA1<>+0x008(SB)/8, $0x0101010101010101
DATA LCDATA1<>+0x010(SB)/8, $0x0202020202020202
DATA LCDATA1<>+0x018(SB)/8, $0x0303030303030303
DATA LCDATA1<>+0x020(SB)/8, $0x7fbfdfeff7fbfdfe
GLOBL LCDATA1<>(SB), 8, $40

TEXT ·_expand_mask_avx(SB), $0-24

	MOVQ in+0(FP), DI
	MOVQ out+8(FP), SI
	MOVQ info+16(FP), DX
	LEAQ LCDATA1<>(SB), BP

	LONG $0x20eac148               // shr    rdx, 32
	WORD $0xfa83; BYTE $0x01       // cmp    edx, 1
	JNE  LBB0_1
	LONG $0x076ef9c5               // vmovd    xmm0, dword [rdi]
	LONG $0x00fde3c4; WORD $0x44c0 // vpermq    ymm0, ymm0, 68
	LONG $0x4d6ffdc5; BYTE $0x00   // vmovdqa    ymm1, yword 0[rbp] /* [rip + LCPI0_0] */
	LONG $0x007de2c4; BYTE $0xc1   // vpshufb    ymm0, ymm0, ymm1
	LONG $0x597de2c4; WORD $0x2055 // vpbroadcastq    ymm2, qword 32[rbp] /* [rip + LCPI0_1] */
	LONG $0xc2ebfdc5               // vpor    ymm0, ymm0, ymm2
	LONG $0xdb76e5c5               // vpcmpeqd    ymm3, ymm3, ymm3
	LONG $0xc374fdc5               // vpcmpeqb    ymm0, ymm0, ymm3
	LONG $0x067ffec5               // vmovdqu    yword [rsi], ymm0
	LONG $0x476ef9c5; BYTE $0x04   // vmovd    xmm0, dword [rdi + 4]
	LONG $0x00fde3c4; WORD $0x44c0 // vpermq    ymm0, ymm0, 68
	LONG $0x007de2c4; BYTE $0xc1   // vpshufb    ymm0, ymm0, ymm1
	LONG $0xc2ebfdc5               // vpor    ymm0, ymm0, ymm2
	LONG $0xc374fdc5               // vpcmpeqb    ymm0, ymm0, ymm3
	LONG $0x467ffec5; BYTE $0x20   // vmovdqu    yword [rsi + 32], ymm0
	JMP  LBB0_3

LBB0_1:
	WORD $0xc031 // xor    eax, eax

LBB0_2:
	WORD $0x8b48; BYTE $0x0f // mov    rcx, qword [rdi]
	LONG $0xc1a30f48         // bt    rcx, rax
	WORD $0xc919             // sbb    ecx, ecx
	WORD $0x0c88; BYTE $0x06 // mov    byte [rsi + rax], cl
	WORD $0x8b48; BYTE $0x0f // mov    rcx, qword [rdi]
	WORD $0x508d; BYTE $0x01 // lea    edx, [rax + 1]
	WORD $0xb60f; BYTE $0xd2 // movzx    edx, dl
	LONG $0xd1a30f48         // bt    rcx, rdx
	WORD $0xc919             // sbb    ecx, ecx
	LONG $0x01064c88         // mov    byte [rsi + rax + 1], cl
	LONG $0x02c08348         // add    rax, 2
	LONG $0x40f88348         // cmp    rax, 64
	JNE  LBB0_2

LBB0_3:
	VZEROUPPER
	RET
