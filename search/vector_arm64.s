//go:build arm64

TEXT ·_magnitude_f32_neon(SB), $0-24
MOVD input+0(FP),  R0
MOVD size+8(FP), R1
MOVD output+16(FP),R2
WORD 	$0xd342fc28	//	LSR $2, R1, R8
WORD 	$0x7100051f	//	CMPW $1, R8
WORD 	$0x540001cb	//	BLT 14(PC)
WORD 	$0xb27e7fe9	//	MOVD $17179869180, R9
WORD 	$0x8b090029	//	ADD R9, R1, R9
WORD 	$0x927e7d29	//	AND $17179869180, R9, R9
WORD 	$0x6f00e400	//	VMOVI $0, V0.D2
WORD 	$0xaa0003ea	//	MOVD R0, R10
WORD 	$0x3cc10541	//	FMOVQ.P 16(R10), F1
WORD 	$0x4e21cc20	//	VFMLA V1.S4, V1.S4, V0.S4
WORD 	$0x71000508	//	SUBSW $1, R8, R8
WORD 	$0x54ffffa1	//	BNE -3(PC)
WORD 	$0x91001129	//	ADD $4, R9, R9
WORD 	$0xeb01013f	//	CMP R1, R9
WORD 	$0x540000c3	//	BCC 6(PC)
WORD 	$0x1400000d	//	JMP 13(PC)
WORD 	$0xd2800009	//	MOVD $0, R9
WORD 	$0x6f00e400	//	VMOVI $0, V0.D2
WORD 	$0xeb01013f	//	CMP R1, R9
WORD 	$0x54000122	//	BCS 9(PC)
WORD 	$0xcb090028	//	SUB R9, R1, R8
WORD 	$0x8b090809	//	ADD R9<<2, R0, R9
WORD 	$0xbc404521	//	FMOVS.P 4(R9), F1
WORD 	$0x1e210821	//	FMULS F1, F1, F1
WORD 	$0x4e040421	//	VDUP V1.S[0], V1.S4
WORD 	$0x4e20d420	//	FADD V0.S4, V1.S4, V0.S4
WORD 	$0xf1000508	//	SUBS $1, R8, R8
WORD 	$0x54ffff61	//	BNE -5(PC)
WORD 	$0x6e20d400	//	VFADDP V0.S4, V0.S4, V0.S4
WORD 	$0x7e30d800	//	FADDP V0.S2, F0
WORD 	$0x1e21c000	//	FSQRTS F0, F0
WORD 	$0xbd000040	//	FMOVS F0, (R2)
WORD 	$0xd65f03c0	//	RET


TEXT ·_magnitude_f32_sve(SB), $0-24
MOVD input+0(FP),  R0
MOVD size+16(FP), R1
MOVD output+24(FP),R2
WORD 	$0x04a0e3e9	//	?
WORD 	$0x9ac9082a	//	UDIV R9, R1, R10
WORD 	$0x7100055f	//	CMPW $1, R10
WORD 	$0x5400020b	//	BLT 16(PC)
WORD 	$0xd2800008	//	MOVD $0, R8
WORD 	$0x5280000b	//	MOVW $0, R11
WORD 	$0x25b8c000	//	?
WORD 	$0x25a11d00	//	?
WORD 	$0xa5484001	//	?
WORD 	$0x65a10020	//	?
WORD 	$0x8b090108	//	ADD R9, R8, R8
WORD 	$0x1100056b	//	ADDW $1, R11, R11
WORD 	$0x6b0a017f	//	CMPW R10, R11
WORD 	$0x54ffff41	//	BNE -6(PC)
WORD 	$0x2598e3e0	//	?
WORD 	$0x65802000	//	?
WORD 	$0xeb01011f	//	CMP R1, R8
WORD 	$0x54000103	//	BCC 8(PC)
WORD 	$0x14000029	//	JMP 41(PC)
WORD 	$0xd2800008	//	MOVD $0, R8
WORD 	$0x25b8c000	//	?
WORD 	$0x2598e3e0	//	?
WORD 	$0x65802000	//	?
WORD 	$0xeb01011f	//	CMP R1, R8
WORD 	$0x54000462	//	BCS 35(PC)
WORD 	$0xcb080029	//	SUB R8, R1, R9
WORD 	$0x0460e3eb	//	?
WORD 	$0xeb0b013f	//	CMP R11, R9
WORD 	$0x54000062	//	BCS 3(PC)
WORD 	$0xaa0803e9	//	MOVD R8, R9
WORD 	$0x14000017	//	JMP 23(PC)
WORD 	$0xd280000c	//	MOVD $0, R12
WORD 	$0x9acb092a	//	UDIV R11, R9, R10
WORD 	$0x9b0b7d4d	//	MUL R11, R10, R13
WORD 	$0xcb0d012a	//	SUB R13, R9, R10
WORD 	$0x8b0d0109	//	ADD R13, R8, R9
WORD 	$0x25b8c001	//	?
WORD 	$0x2598e021	//	?
WORD 	$0x05a1c400	//	?
WORD 	$0x04bf502e	//	?
WORD 	$0xd344fdce	//	LSR $4, R14, R14
WORD 	$0x8b080808	//	ADD R8<<2, R0, R8
WORD 	$0x8b2e510e	//	ADD R14.UXTW<<4, R8, R14
WORD 	$0xa54c4102	//	?
WORD 	$0xa54c41c3	//	?
WORD 	$0x65a20040	//	?
WORD 	$0x65a30061	//	?
WORD 	$0x8b0b018c	//	ADD R11, R12, R12
WORD 	$0xeb0d019f	//	CMP R13, R12
WORD 	$0x54ffff41	//	BNE -6(PC)
WORD 	$0x65800020	//	?
WORD 	$0x65802000	//	?
WORD 	$0xb40000ea	//	CBZ R10, 7(PC)
WORD 	$0xcb090028	//	SUB R9, R1, R8
WORD 	$0x8b090809	//	ADD R9<<2, R0, R9
WORD 	$0xbc404521	//	FMOVS.P 4(R9), F1
WORD 	$0x1f010020	//	FMADDS F1, F0, F1, F0
WORD 	$0xf1000508	//	SUBS $1, R8, R8
WORD 	$0x54ffffa1	//	BNE -3(PC)
WORD 	$0x1e21c000	//	FSQRTS F0, F0
WORD 	$0xbd000040	//	FMOVS F0, (R2)
WORD 	$0xd65f03c0	//	RET
