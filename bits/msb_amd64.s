//+build !noasm !appengine
// AUTO-GENERATED BY C2GOASM -- DO NOT EDIT

TEXT ·_msb_avx2(SB), $0-24

	MOVQ data+0(FP), DI
	MOVQ len+8(FP), SI
	MOVQ result+16(FP), DX

	LONG $0xff02c748; WORD $0xffff; BYTE $0xff // mov    qword [rdx], -1
	WORD $0x8948; BYTE $0xf0                   // mov    rax, rsi
	LONG $0x20e8c148                           // shr    rax, 32
	WORD $0xf883; BYTE $0x01                   // cmp    eax, 1
	JNE  LBB0_4
	WORD $0x8941; BYTE $0xf0                   // mov    r8d, esi
	LONG $0x07e08341                           // and    r8d, 7
	WORD $0x468d; BYTE $0xfc                   // lea    eax, [rsi - 4]
	WORD $0x8941; BYTE $0xf1                   // mov    r9d, esi
	LONG $0xffe18141; WORD $0xffff; BYTE $0x1f // and    r9d, 536870911
	LONG $0x03e1c149                           // shl    r9, 3
	WORD $0x4e8d; BYTE $0xf8                   // lea    ecx, [rsi - 8]
	WORD $0x6348; BYTE $0xc9                   // movsxd    rcx, ecx
	LONG $0xcf0c8d48                           // lea    rcx, [rdi + 8*rcx]
	LONG $0x40c18348                           // add    rcx, 64

LBB0_2:
	WORD $0x708d; BYTE $0xfc       // lea    esi, [rax - 4]
	WORD $0x3944; BYTE $0xc6       // cmp    esi, r8d
	JL   LBB0_3
	LONG $0x416ffec5; BYTE $0xc0   // vmovdqu    ymm0, yword [rcx - 64]
	LONG $0xc0c18348               // add    rcx, -64
	WORD $0x6348; BYTE $0xf0       // movsxd    rsi, eax
	LONG $0x0c6ffec5; BYTE $0xf7   // vmovdqu    ymm1, yword [rdi + 8*rsi]
	LONG $0xd0ebf5c5               // vpor    ymm2, ymm1, ymm0
	WORD $0xc083; BYTE $0xf8       // add    eax, -8
	LONG $0xc0c18349               // add    r9, -64
	LONG $0x177de2c4; BYTE $0xd2   // vptest    ymm2, ymm2
	JE   LBB0_2
	LONG $0xd2efe9c5               // vpxor    xmm2, xmm2, xmm2
	LONG $0xc276fdc5               // vpcmpeqd    ymm0, ymm0, ymm2
	LONG $0xc050fcc5               // vmovmskps    eax, ymm0
	LONG $0xc276f5c5               // vpcmpeqd    ymm0, ymm1, ymm2
	LONG $0xf050fcc5               // vmovmskps    esi, ymm0
	WORD $0xe6c1; BYTE $0x08       // shl    esi, 8
	WORD $0xc609                   // or    esi, eax
	LONG $0xfffff681; WORD $0x0000 // xor    esi, 65535
	WORD $0xbd0f; BYTE $0xc6       // bsr    eax, esi
	WORD $0xf083; BYTE $0x1f       // xor    eax, 31
	WORD $0xe0c1; BYTE $0x02       // shl    eax, 2
	LONG $0x7cf08348               // xor    rax, 124
	WORD $0x0141; BYTE $0xc1       // add    r9d, eax
	LONG $0x03e1c141               // shl    r9d, 3
	LONG $0x1fc98341               // or    r9d, 31
	LONG $0x0104bd0f               // bsr    eax, dword [rcx + rax]
	WORD $0xf083; BYTE $0x1f       // xor    eax, 31
	WORD $0x2941; BYTE $0xc1       // sub    r9d, eax
	JMP  LBB0_10

LBB0_3:
	WORD $0x8944; BYTE $0xc6 // mov    esi, r8d

LBB0_4:
	WORD $0xf089             // mov    eax, esi
	WORD $0xe6c1; BYTE $0x06 // shl    esi, 6
	WORD $0xce83; BYTE $0x3f // or    esi, 63

LBB0_5:
	WORD $0xc085                 // test    eax, eax
	JLE  LBB0_11
	LONG $0xc74c8b48; BYTE $0xf8 // mov    rcx, qword [rdi + 8*rax - 8]
	WORD $0xc683; BYTE $0xc0     // add    esi, -64
	LONG $0xffc08348             // add    rax, -1
	WORD $0x8548; BYTE $0xc9     // test    rcx, rcx
	JE   LBB0_5
	LONG $0xc1bd0f48             // bsr    rax, rcx
	WORD $0xf083; BYTE $0x3f     // xor    eax, 63
	WORD $0xc629                 // sub    esi, eax
	WORD $0x634c; BYTE $0xce     // movsxd    r9, esi

LBB0_10:
	WORD $0x894c; BYTE $0x0a // mov    qword [rdx], r9

LBB0_11:
	VZEROUPPER
	RET
