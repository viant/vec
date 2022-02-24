//+build !noasm !appengine
// AUTO-GENERATED BY C2GOASM -- DO NOT EDIT

TEXT ·_add_int32(SB), $0-32

	MOVQ input1+0(FP), DI
	MOVQ input2+8(FP), SI
	MOVQ output+16(FP), DX
	MOVQ info+24(FP), CX

	WORD $0x8948; BYTE $0xc8                   // mov    rax, rcx
	LONG $0x20e8c148                           // shr    rax, 32
	WORD $0x3145; BYTE $0xc0                   // xor    r8d, r8d
	WORD $0xf883; BYTE $0x01                   // cmp    eax, 1
	JNE  LBB0_8
	WORD $0xf983; BYTE $0x08                   // cmp    ecx, 8
	JB   LBB0_7
	WORD $0x8949; BYTE $0xc8                   // mov    r8, rcx
	LONG $0x03e8c149                           // shr    r8, 3
	WORD $0x8945; BYTE $0xc1                   // mov    r9d, r8d
	LONG $0xffe18141; WORD $0xffff; BYTE $0x1f // and    r9d, 536870911
	LONG $0xff418d49                           // lea    rax, [r9 - 1]
	LONG $0x03e08341                           // and    r8d, 3
	LONG $0x03f88348                           // cmp    rax, 3
	JAE  LBB0_11
	WORD $0x3145; BYTE $0xd2                   // xor    r10d, r10d
	JMP  LBB0_4

LBB0_11:
	WORD $0x294d; BYTE $0xc1     // sub    r9, r8
	LONG $0x000060b8; BYTE $0x00 // mov    eax, 96
	WORD $0x3145; BYTE $0xd2     // xor    r10d, r10d

LBB0_12:
	LONG $0x446ffec5; WORD $0xa006 // vmovdqu    ymm0, yword [rsi + rax - 96]
	LONG $0x44fefdc5; WORD $0xa007 // vpaddd    ymm0, ymm0, yword [rdi + rax - 96]
	LONG $0x447ffec5; WORD $0xa002 // vmovdqu    yword [rdx + rax - 96], ymm0
	LONG $0x446ffec5; WORD $0xc006 // vmovdqu    ymm0, yword [rsi + rax - 64]
	LONG $0x44fefdc5; WORD $0xc007 // vpaddd    ymm0, ymm0, yword [rdi + rax - 64]
	LONG $0x447ffec5; WORD $0xc002 // vmovdqu    yword [rdx + rax - 64], ymm0
	LONG $0x446ffec5; WORD $0xe006 // vmovdqu    ymm0, yword [rsi + rax - 32]
	LONG $0x44fefdc5; WORD $0xe007 // vpaddd    ymm0, ymm0, yword [rdi + rax - 32]
	LONG $0x447ffec5; WORD $0xe002 // vmovdqu    yword [rdx + rax - 32], ymm0
	LONG $0x046ffec5; BYTE $0x06   // vmovdqu    ymm0, yword [rsi + rax]
	LONG $0x04fefdc5; BYTE $0x07   // vpaddd    ymm0, ymm0, yword [rdi + rax]
	LONG $0x047ffec5; BYTE $0x02   // vmovdqu    yword [rdx + rax], ymm0
	LONG $0x04c28349               // add    r10, 4
	LONG $0x80e88348               // sub    rax, -128
	WORD $0x394d; BYTE $0xd1       // cmp    r9, r10
	JNE  LBB0_12

LBB0_4:
	WORD $0x854d; BYTE $0xc0 // test    r8, r8
	JE   LBB0_7
	LONG $0x05e2c149         // shl    r10, 5
	LONG $0x160c8d4e         // lea    r9, [rsi + r10]
	LONG $0x171c8d4a         // lea    rbx, [rdi + r10]
	WORD $0x0149; BYTE $0xd2 // add    r10, rdx
	LONG $0x05e0c149         // shl    r8, 5
	WORD $0xc031             // xor    eax, eax

LBB0_6:
	LONG $0x6f7ec1c4; WORD $0x0104 // vmovdqu    ymm0, yword [r9 + rax]
	LONG $0x04fefdc5; BYTE $0x03   // vpaddd    ymm0, ymm0, yword [rbx + rax]
	LONG $0x7f7ec1c4; WORD $0x0204 // vmovdqu    yword [r10 + rax], ymm0
	LONG $0x20c08348               // add    rax, 32
	WORD $0x3949; BYTE $0xc0       // cmp    r8, rax
	JNE  LBB0_6

LBB0_7:
	WORD $0x8941; BYTE $0xc8 // mov    r8d, ecx
	LONG $0xf8e08341         // and    r8d, -8

LBB0_8:
	WORD $0x3941; BYTE $0xc8 // cmp    r8d, ecx
	JAE  LBB0_21
	WORD $0x634d; BYTE $0xc8 // movsxd    r9, r8d
	WORD $0xf741; BYTE $0xd0 // not    r8d
	WORD $0x0141; BYTE $0xc8 // add    r8d, ecx
	LONG $0x1ff88341         // cmp    r8d, 31
	JAE  LBB0_13
	WORD $0x894c; BYTE $0xc8 // mov    rax, r9
	JMP  LBB0_20

LBB0_13:
	LONG $0x8a148d4e         // lea    r10, [rdx + 4*r9]
	LONG $0x081c8d4b         // lea    rbx, [r8 + r9]
	LONG $0x9a048d48         // lea    rax, [rdx + 4*rbx]
	LONG $0x04c08348         // add    rax, 4
	LONG $0x8f1c8d4e         // lea    r11, [rdi + 4*r9]
	LONG $0x9f348d4c         // lea    r14, [rdi + 4*rbx]
	LONG $0x04c68349         // add    r14, 4
	LONG $0x8e3c8d4e         // lea    r15, [rsi + 4*r9]
	LONG $0x9e1c8d48         // lea    rbx, [rsi + 4*rbx]
	LONG $0x04c38348         // add    rbx, 4
	WORD $0x394d; BYTE $0xf2 // cmp    r10, r14
	LONG $0xc6920f41         // setb    r14b
	WORD $0x3949; BYTE $0xc3 // cmp    r11, rax
	LONG $0xc3920f41         // setb    r11b
	WORD $0x3949; BYTE $0xda // cmp    r10, rbx
	WORD $0x920f; BYTE $0xc3 // setb    bl
	WORD $0x3949; BYTE $0xc7 // cmp    r15, rax
	WORD $0x920f; BYTE $0xc0 // setb    al
	WORD $0x8445; BYTE $0xde // test    r14b, r11b
	JNE  LBB0_14
	WORD $0xc320             // and    bl, al
	JNE  LBB0_16
	LONG $0x01c08349         // add    r8, 1
	WORD $0x894d; BYTE $0xc2 // mov    r10, r8
	LONG $0xe0e28349         // and    r10, -32
	LONG $0x0a048d4b         // lea    rax, [r10 + r9]
	LONG $0x8f1c8d4e         // lea    r11, [rdi + 4*r9]
	LONG $0x60c38349         // add    r11, 96
	LONG $0x8e348d4e         // lea    r14, [rsi + 4*r9]
	LONG $0x60c68349         // add    r14, 96
	LONG $0x8a0c8d4e         // lea    r9, [rdx + 4*r9]
	LONG $0x60c18349         // add    r9, 96
	WORD $0xdb31             // xor    ebx, ebx

LBB0_18:
	LONG $0x6f7ec1c4; WORD $0x9e44; BYTE $0xa0 // vmovdqu    ymm0, yword [r14 + 4*rbx - 96]
	LONG $0x6f7ec1c4; WORD $0x9e4c; BYTE $0xc0 // vmovdqu    ymm1, yword [r14 + 4*rbx - 64]
	LONG $0x6f7ec1c4; WORD $0x9e54; BYTE $0xe0 // vmovdqu    ymm2, yword [r14 + 4*rbx - 32]
	LONG $0x6f7ec1c4; WORD $0x9e1c             // vmovdqu    ymm3, yword [r14 + 4*rbx]
	LONG $0xfe7dc1c4; WORD $0x9b44; BYTE $0xa0 // vpaddd    ymm0, ymm0, yword [r11 + 4*rbx - 96]
	LONG $0xfe75c1c4; WORD $0x9b4c; BYTE $0xc0 // vpaddd    ymm1, ymm1, yword [r11 + 4*rbx - 64]
	LONG $0xfe6dc1c4; WORD $0x9b54; BYTE $0xe0 // vpaddd    ymm2, ymm2, yword [r11 + 4*rbx - 32]
	LONG $0xfe65c1c4; WORD $0x9b1c             // vpaddd    ymm3, ymm3, yword [r11 + 4*rbx]
	LONG $0x7f7ec1c4; WORD $0x9944; BYTE $0xa0 // vmovdqu    yword [r9 + 4*rbx - 96], ymm0
	LONG $0x7f7ec1c4; WORD $0x994c; BYTE $0xc0 // vmovdqu    yword [r9 + 4*rbx - 64], ymm1
	LONG $0x7f7ec1c4; WORD $0x9954; BYTE $0xe0 // vmovdqu    yword [r9 + 4*rbx - 32], ymm2
	LONG $0x7f7ec1c4; WORD $0x991c             // vmovdqu    yword [r9 + 4*rbx], ymm3
	LONG $0x20c38348                           // add    rbx, 32
	WORD $0x3949; BYTE $0xda                   // cmp    r10, rbx
	JNE  LBB0_18
	WORD $0x394d; BYTE $0xd0                   // cmp    r8, r10
	JNE  LBB0_20
	JMP  LBB0_21

LBB0_14:
	WORD $0x894c; BYTE $0xc8 // mov    rax, r9
	JMP  LBB0_20

LBB0_16:
	WORD $0x894c; BYTE $0xc8 // mov    rax, r9

LBB0_20:
	WORD $0x1c8b; BYTE $0x86 // mov    ebx, dword [rsi + 4*rax]
	WORD $0x1c03; BYTE $0x87 // add    ebx, dword [rdi + 4*rax]
	WORD $0x1c89; BYTE $0x82 // mov    dword [rdx + 4*rax], ebx
	LONG $0x01c08348         // add    rax, 1
	WORD $0xc839             // cmp    eax, ecx
	JB   LBB0_20

LBB0_21:
	VZEROUPPER
	RET

TEXT ·_sub_int32(SB), $0-32

	MOVQ input1+0(FP), DI
	MOVQ input2+8(FP), SI
	MOVQ output+16(FP), DX
	MOVQ info+24(FP), CX

	WORD $0x8948; BYTE $0xc8                   // mov    rax, rcx
	LONG $0x20e8c148                           // shr    rax, 32
	WORD $0x3145; BYTE $0xc0                   // xor    r8d, r8d
	WORD $0xf883; BYTE $0x01                   // cmp    eax, 1
	JNE  LBB1_8
	WORD $0xf983; BYTE $0x08                   // cmp    ecx, 8
	JB   LBB1_7
	WORD $0x8949; BYTE $0xc8                   // mov    r8, rcx
	LONG $0x03e8c149                           // shr    r8, 3
	WORD $0x8945; BYTE $0xc1                   // mov    r9d, r8d
	LONG $0xffe18141; WORD $0xffff; BYTE $0x1f // and    r9d, 536870911
	LONG $0xff418d49                           // lea    rax, [r9 - 1]
	LONG $0x03e08341                           // and    r8d, 3
	LONG $0x03f88348                           // cmp    rax, 3
	JAE  LBB1_11
	WORD $0x3145; BYTE $0xd2                   // xor    r10d, r10d
	JMP  LBB1_4

LBB1_11:
	WORD $0x294d; BYTE $0xc1     // sub    r9, r8
	LONG $0x000060b8; BYTE $0x00 // mov    eax, 96
	WORD $0x3145; BYTE $0xd2     // xor    r10d, r10d

LBB1_12:
	LONG $0x446ffec5; WORD $0xa007 // vmovdqu    ymm0, yword [rdi + rax - 96]
	LONG $0x44fafdc5; WORD $0xa006 // vpsubd    ymm0, ymm0, yword [rsi + rax - 96]
	LONG $0x447ffec5; WORD $0xa002 // vmovdqu    yword [rdx + rax - 96], ymm0
	LONG $0x446ffec5; WORD $0xc007 // vmovdqu    ymm0, yword [rdi + rax - 64]
	LONG $0x44fafdc5; WORD $0xc006 // vpsubd    ymm0, ymm0, yword [rsi + rax - 64]
	LONG $0x447ffec5; WORD $0xc002 // vmovdqu    yword [rdx + rax - 64], ymm0
	LONG $0x446ffec5; WORD $0xe007 // vmovdqu    ymm0, yword [rdi + rax - 32]
	LONG $0x44fafdc5; WORD $0xe006 // vpsubd    ymm0, ymm0, yword [rsi + rax - 32]
	LONG $0x447ffec5; WORD $0xe002 // vmovdqu    yword [rdx + rax - 32], ymm0
	LONG $0x046ffec5; BYTE $0x07   // vmovdqu    ymm0, yword [rdi + rax]
	LONG $0x04fafdc5; BYTE $0x06   // vpsubd    ymm0, ymm0, yword [rsi + rax]
	LONG $0x047ffec5; BYTE $0x02   // vmovdqu    yword [rdx + rax], ymm0
	LONG $0x04c28349               // add    r10, 4
	LONG $0x80e88348               // sub    rax, -128
	WORD $0x394d; BYTE $0xd1       // cmp    r9, r10
	JNE  LBB1_12

LBB1_4:
	WORD $0x854d; BYTE $0xc0 // test    r8, r8
	JE   LBB1_7
	LONG $0x05e2c149         // shl    r10, 5
	LONG $0x160c8d4e         // lea    r9, [rsi + r10]
	LONG $0x171c8d4a         // lea    rbx, [rdi + r10]
	WORD $0x0149; BYTE $0xd2 // add    r10, rdx
	LONG $0x05e0c149         // shl    r8, 5
	WORD $0xc031             // xor    eax, eax

LBB1_6:
	LONG $0x046ffec5; BYTE $0x03   // vmovdqu    ymm0, yword [rbx + rax]
	LONG $0xfa7dc1c4; WORD $0x0104 // vpsubd    ymm0, ymm0, yword [r9 + rax]
	LONG $0x7f7ec1c4; WORD $0x0204 // vmovdqu    yword [r10 + rax], ymm0
	LONG $0x20c08348               // add    rax, 32
	WORD $0x3949; BYTE $0xc0       // cmp    r8, rax
	JNE  LBB1_6

LBB1_7:
	WORD $0x8941; BYTE $0xc8 // mov    r8d, ecx
	LONG $0xf8e08341         // and    r8d, -8

LBB1_8:
	WORD $0x3941; BYTE $0xc8 // cmp    r8d, ecx
	JAE  LBB1_21
	WORD $0x634d; BYTE $0xc8 // movsxd    r9, r8d
	WORD $0xf741; BYTE $0xd0 // not    r8d
	WORD $0x0141; BYTE $0xc8 // add    r8d, ecx
	LONG $0x1ff88341         // cmp    r8d, 31
	JAE  LBB1_13
	WORD $0x894c; BYTE $0xc8 // mov    rax, r9
	JMP  LBB1_20

LBB1_13:
	LONG $0x8a148d4e         // lea    r10, [rdx + 4*r9]
	LONG $0x081c8d4b         // lea    rbx, [r8 + r9]
	LONG $0x9a048d48         // lea    rax, [rdx + 4*rbx]
	LONG $0x04c08348         // add    rax, 4
	LONG $0x8f1c8d4e         // lea    r11, [rdi + 4*r9]
	LONG $0x9f348d4c         // lea    r14, [rdi + 4*rbx]
	LONG $0x04c68349         // add    r14, 4
	LONG $0x8e3c8d4e         // lea    r15, [rsi + 4*r9]
	LONG $0x9e1c8d48         // lea    rbx, [rsi + 4*rbx]
	LONG $0x04c38348         // add    rbx, 4
	WORD $0x394d; BYTE $0xf2 // cmp    r10, r14
	LONG $0xc6920f41         // setb    r14b
	WORD $0x3949; BYTE $0xc3 // cmp    r11, rax
	LONG $0xc3920f41         // setb    r11b
	WORD $0x3949; BYTE $0xda // cmp    r10, rbx
	WORD $0x920f; BYTE $0xc3 // setb    bl
	WORD $0x3949; BYTE $0xc7 // cmp    r15, rax
	WORD $0x920f; BYTE $0xc0 // setb    al
	WORD $0x8445; BYTE $0xde // test    r14b, r11b
	JNE  LBB1_14
	WORD $0xc320             // and    bl, al
	JNE  LBB1_16
	LONG $0x01c08349         // add    r8, 1
	WORD $0x894d; BYTE $0xc2 // mov    r10, r8
	LONG $0xe0e28349         // and    r10, -32
	LONG $0x0a048d4b         // lea    rax, [r10 + r9]
	LONG $0x8f1c8d4e         // lea    r11, [rdi + 4*r9]
	LONG $0x60c38349         // add    r11, 96
	LONG $0x8e348d4e         // lea    r14, [rsi + 4*r9]
	LONG $0x60c68349         // add    r14, 96
	LONG $0x8a0c8d4e         // lea    r9, [rdx + 4*r9]
	LONG $0x60c18349         // add    r9, 96
	WORD $0xdb31             // xor    ebx, ebx

LBB1_18:
	LONG $0x6f7ec1c4; WORD $0x9b44; BYTE $0xa0 // vmovdqu    ymm0, yword [r11 + 4*rbx - 96]
	LONG $0x6f7ec1c4; WORD $0x9b4c; BYTE $0xc0 // vmovdqu    ymm1, yword [r11 + 4*rbx - 64]
	LONG $0x6f7ec1c4; WORD $0x9b54; BYTE $0xe0 // vmovdqu    ymm2, yword [r11 + 4*rbx - 32]
	LONG $0x6f7ec1c4; WORD $0x9b1c             // vmovdqu    ymm3, yword [r11 + 4*rbx]
	LONG $0xfa7dc1c4; WORD $0x9e44; BYTE $0xa0 // vpsubd    ymm0, ymm0, yword [r14 + 4*rbx - 96]
	LONG $0xfa75c1c4; WORD $0x9e4c; BYTE $0xc0 // vpsubd    ymm1, ymm1, yword [r14 + 4*rbx - 64]
	LONG $0xfa6dc1c4; WORD $0x9e54; BYTE $0xe0 // vpsubd    ymm2, ymm2, yword [r14 + 4*rbx - 32]
	LONG $0xfa65c1c4; WORD $0x9e1c             // vpsubd    ymm3, ymm3, yword [r14 + 4*rbx]
	LONG $0x7f7ec1c4; WORD $0x9944; BYTE $0xa0 // vmovdqu    yword [r9 + 4*rbx - 96], ymm0
	LONG $0x7f7ec1c4; WORD $0x994c; BYTE $0xc0 // vmovdqu    yword [r9 + 4*rbx - 64], ymm1
	LONG $0x7f7ec1c4; WORD $0x9954; BYTE $0xe0 // vmovdqu    yword [r9 + 4*rbx - 32], ymm2
	LONG $0x7f7ec1c4; WORD $0x991c             // vmovdqu    yword [r9 + 4*rbx], ymm3
	LONG $0x20c38348                           // add    rbx, 32
	WORD $0x3949; BYTE $0xda                   // cmp    r10, rbx
	JNE  LBB1_18
	WORD $0x394d; BYTE $0xd0                   // cmp    r8, r10
	JNE  LBB1_20
	JMP  LBB1_21

LBB1_14:
	WORD $0x894c; BYTE $0xc8 // mov    rax, r9
	JMP  LBB1_20

LBB1_16:
	WORD $0x894c; BYTE $0xc8 // mov    rax, r9

LBB1_20:
	WORD $0x1c8b; BYTE $0x87 // mov    ebx, dword [rdi + 4*rax]
	WORD $0x1c2b; BYTE $0x86 // sub    ebx, dword [rsi + 4*rax]
	WORD $0x1c89; BYTE $0x82 // mov    dword [rdx + 4*rax], ebx
	LONG $0x01c08348         // add    rax, 1
	WORD $0xc839             // cmp    eax, ecx
	JB   LBB1_20

LBB1_21:
	VZEROUPPER
	RET

TEXT ·_mul_int32(SB), $0-32

	MOVQ input1+0(FP), DI
	MOVQ input2+8(FP), SI
	MOVQ output+16(FP), DX
	MOVQ info+24(FP), CX

	WORD $0x8948; BYTE $0xc8                   // mov    rax, rcx
	LONG $0x20e8c148                           // shr    rax, 32
	WORD $0x3145; BYTE $0xc0                   // xor    r8d, r8d
	WORD $0xf883; BYTE $0x01                   // cmp    eax, 1
	JNE  LBB2_8
	WORD $0xf983; BYTE $0x08                   // cmp    ecx, 8
	JB   LBB2_7
	WORD $0x8949; BYTE $0xc8                   // mov    r8, rcx
	LONG $0x03e8c149                           // shr    r8, 3
	WORD $0x8945; BYTE $0xc1                   // mov    r9d, r8d
	LONG $0xffe18141; WORD $0xffff; BYTE $0x1f // and    r9d, 536870911
	LONG $0xff418d49                           // lea    rax, [r9 - 1]
	LONG $0x03e08341                           // and    r8d, 3
	LONG $0x03f88348                           // cmp    rax, 3
	JAE  LBB2_11
	WORD $0x3145; BYTE $0xd2                   // xor    r10d, r10d
	JMP  LBB2_4

LBB2_11:
	WORD $0x294d; BYTE $0xc1     // sub    r9, r8
	LONG $0x000060b8; BYTE $0x00 // mov    eax, 96
	WORD $0x3145; BYTE $0xd2     // xor    r10d, r10d

LBB2_12:
	LONG $0x446ffec5; WORD $0xa006             // vmovdqu    ymm0, yword [rsi + rax - 96]
	LONG $0x407de2c4; WORD $0x0744; BYTE $0xa0 // vpmulld    ymm0, ymm0, yword [rdi + rax - 96]
	LONG $0x447ffec5; WORD $0xa002             // vmovdqu    yword [rdx + rax - 96], ymm0
	LONG $0x446ffec5; WORD $0xc006             // vmovdqu    ymm0, yword [rsi + rax - 64]
	LONG $0x407de2c4; WORD $0x0744; BYTE $0xc0 // vpmulld    ymm0, ymm0, yword [rdi + rax - 64]
	LONG $0x447ffec5; WORD $0xc002             // vmovdqu    yword [rdx + rax - 64], ymm0
	LONG $0x446ffec5; WORD $0xe006             // vmovdqu    ymm0, yword [rsi + rax - 32]
	LONG $0x407de2c4; WORD $0x0744; BYTE $0xe0 // vpmulld    ymm0, ymm0, yword [rdi + rax - 32]
	LONG $0x447ffec5; WORD $0xe002             // vmovdqu    yword [rdx + rax - 32], ymm0
	LONG $0x046ffec5; BYTE $0x06               // vmovdqu    ymm0, yword [rsi + rax]
	LONG $0x407de2c4; WORD $0x0704             // vpmulld    ymm0, ymm0, yword [rdi + rax]
	LONG $0x047ffec5; BYTE $0x02               // vmovdqu    yword [rdx + rax], ymm0
	LONG $0x04c28349                           // add    r10, 4
	LONG $0x80e88348                           // sub    rax, -128
	WORD $0x394d; BYTE $0xd1                   // cmp    r9, r10
	JNE  LBB2_12

LBB2_4:
	WORD $0x854d; BYTE $0xc0 // test    r8, r8
	JE   LBB2_7
	LONG $0x05e2c149         // shl    r10, 5
	LONG $0x160c8d4e         // lea    r9, [rsi + r10]
	LONG $0x171c8d4a         // lea    rbx, [rdi + r10]
	WORD $0x0149; BYTE $0xd2 // add    r10, rdx
	LONG $0x05e0c149         // shl    r8, 5
	WORD $0xc031             // xor    eax, eax

LBB2_6:
	LONG $0x6f7ec1c4; WORD $0x0104 // vmovdqu    ymm0, yword [r9 + rax]
	LONG $0x407de2c4; WORD $0x0304 // vpmulld    ymm0, ymm0, yword [rbx + rax]
	LONG $0x7f7ec1c4; WORD $0x0204 // vmovdqu    yword [r10 + rax], ymm0
	LONG $0x20c08348               // add    rax, 32
	WORD $0x3949; BYTE $0xc0       // cmp    r8, rax
	JNE  LBB2_6

LBB2_7:
	WORD $0x8941; BYTE $0xc8 // mov    r8d, ecx
	LONG $0xf8e08341         // and    r8d, -8

LBB2_8:
	WORD $0x3941; BYTE $0xc8 // cmp    r8d, ecx
	JAE  LBB2_21
	WORD $0x634d; BYTE $0xc8 // movsxd    r9, r8d
	WORD $0xf741; BYTE $0xd0 // not    r8d
	WORD $0x0141; BYTE $0xc8 // add    r8d, ecx
	LONG $0x1ff88341         // cmp    r8d, 31
	JAE  LBB2_13
	WORD $0x894c; BYTE $0xc8 // mov    rax, r9
	JMP  LBB2_20

LBB2_13:
	LONG $0x8a148d4e         // lea    r10, [rdx + 4*r9]
	LONG $0x081c8d4b         // lea    rbx, [r8 + r9]
	LONG $0x9a048d48         // lea    rax, [rdx + 4*rbx]
	LONG $0x04c08348         // add    rax, 4
	LONG $0x8f1c8d4e         // lea    r11, [rdi + 4*r9]
	LONG $0x9f348d4c         // lea    r14, [rdi + 4*rbx]
	LONG $0x04c68349         // add    r14, 4
	LONG $0x8e3c8d4e         // lea    r15, [rsi + 4*r9]
	LONG $0x9e1c8d48         // lea    rbx, [rsi + 4*rbx]
	LONG $0x04c38348         // add    rbx, 4
	WORD $0x394d; BYTE $0xf2 // cmp    r10, r14
	LONG $0xc6920f41         // setb    r14b
	WORD $0x3949; BYTE $0xc3 // cmp    r11, rax
	LONG $0xc3920f41         // setb    r11b
	WORD $0x3949; BYTE $0xda // cmp    r10, rbx
	WORD $0x920f; BYTE $0xc3 // setb    bl
	WORD $0x3949; BYTE $0xc7 // cmp    r15, rax
	WORD $0x920f; BYTE $0xc0 // setb    al
	WORD $0x8445; BYTE $0xde // test    r14b, r11b
	JNE  LBB2_14
	WORD $0xc320             // and    bl, al
	JNE  LBB2_16
	LONG $0x01c08349         // add    r8, 1
	WORD $0x894d; BYTE $0xc2 // mov    r10, r8
	LONG $0xe0e28349         // and    r10, -32
	LONG $0x0a048d4b         // lea    rax, [r10 + r9]
	LONG $0x8f1c8d4e         // lea    r11, [rdi + 4*r9]
	LONG $0x60c38349         // add    r11, 96
	LONG $0x8e348d4e         // lea    r14, [rsi + 4*r9]
	LONG $0x60c68349         // add    r14, 96
	LONG $0x8a0c8d4e         // lea    r9, [rdx + 4*r9]
	LONG $0x60c18349         // add    r9, 96
	WORD $0xdb31             // xor    ebx, ebx

LBB2_18:
	LONG $0x6f7ec1c4; WORD $0x9e44; BYTE $0xa0 // vmovdqu    ymm0, yword [r14 + 4*rbx - 96]
	LONG $0x6f7ec1c4; WORD $0x9e4c; BYTE $0xc0 // vmovdqu    ymm1, yword [r14 + 4*rbx - 64]
	LONG $0x6f7ec1c4; WORD $0x9e54; BYTE $0xe0 // vmovdqu    ymm2, yword [r14 + 4*rbx - 32]
	LONG $0x6f7ec1c4; WORD $0x9e1c             // vmovdqu    ymm3, yword [r14 + 4*rbx]
	LONG $0x407dc2c4; WORD $0x9b44; BYTE $0xa0 // vpmulld    ymm0, ymm0, yword [r11 + 4*rbx - 96]
	LONG $0x4075c2c4; WORD $0x9b4c; BYTE $0xc0 // vpmulld    ymm1, ymm1, yword [r11 + 4*rbx - 64]
	LONG $0x406dc2c4; WORD $0x9b54; BYTE $0xe0 // vpmulld    ymm2, ymm2, yword [r11 + 4*rbx - 32]
	LONG $0x4065c2c4; WORD $0x9b1c             // vpmulld    ymm3, ymm3, yword [r11 + 4*rbx]
	LONG $0x7f7ec1c4; WORD $0x9944; BYTE $0xa0 // vmovdqu    yword [r9 + 4*rbx - 96], ymm0
	LONG $0x7f7ec1c4; WORD $0x994c; BYTE $0xc0 // vmovdqu    yword [r9 + 4*rbx - 64], ymm1
	LONG $0x7f7ec1c4; WORD $0x9954; BYTE $0xe0 // vmovdqu    yword [r9 + 4*rbx - 32], ymm2
	LONG $0x7f7ec1c4; WORD $0x991c             // vmovdqu    yword [r9 + 4*rbx], ymm3
	LONG $0x20c38348                           // add    rbx, 32
	WORD $0x3949; BYTE $0xda                   // cmp    r10, rbx
	JNE  LBB2_18
	WORD $0x394d; BYTE $0xd0                   // cmp    r8, r10
	JNE  LBB2_20
	JMP  LBB2_21

LBB2_14:
	WORD $0x894c; BYTE $0xc8 // mov    rax, r9
	JMP  LBB2_20

LBB2_16:
	WORD $0x894c; BYTE $0xc8 // mov    rax, r9

LBB2_20:
	WORD $0x1c8b; BYTE $0x86 // mov    ebx, dword [rsi + 4*rax]
	LONG $0x871caf0f         // imul    ebx, dword [rdi + 4*rax]
	WORD $0x1c89; BYTE $0x82 // mov    dword [rdx + 4*rax], ebx
	LONG $0x01c08348         // add    rax, 1
	WORD $0xc839             // cmp    eax, ecx
	JB   LBB2_20

LBB2_21:
	VZEROUPPER
	RET