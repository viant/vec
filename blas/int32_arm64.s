TEXT ·_inc_int32(SB), $0-32

   MOVD vec+0(FP),  R0
   MOVD constant+8(FP), R1
   MOVD size+16(FP), R2

   WORD $0xd342fc48         // LSR $2, R2, R8
   WORD $0x4e040c20         // VDUP R1, V0.S4
   WORD $0x7100051f         // CMPW $1, R8
   WORD $0x5400010b         // BLT 8(PC)
   WORD $0xd3428449         // UBFX $2, R2, $32, R9
   WORD $0xaa0003ea         // MOVD R0, R10
   WORD $0x3dc00141         // FMOVQ (R10), F1
   WORD $0x4ea08421         // VADD V0.S4, V1.S4, V1.S4
   WORD $0x3c810541         // FMOVQ.P F1, 16(R10)
   WORD $0xf1000529         // SUBS $1, R9, R9
   WORD $0x54ffff81         // BNE -4(PC)
   WORD $0x531e7508         // LSLW $2, R8, R8
   WORD $0x93407d0b         // SXTW R8, R11
   WORD $0xeb02017f         // CMP R2, R11
   WORD $0x540003a2         // BCS 29(PC)
   WORD $0xcb0b0049         // SUB R11, R2, R9
   WORD $0xaa0b03e8         // MOVD R11, R8
   WORD $0xf100413f         // CMP $16, R9
   WORD $0x54000243         // BCC 18(PC)
   WORD $0x927ced2a         // AND $-16, R9, R10
   WORD $0x8b0b0148         // ADD R11, R10, R8
   WORD $0x8b0b080b         // ADD R11<<2, R0, R11
   WORD $0x9100816b         // ADD $32, R11, R11
   WORD $0xaa0a03ec         // MOVD R10, R12
   WORD $0xad7f0961         // FLDPQ -32(R11), (F1, F2)
   WORD $0xad401163         // FLDPQ (R11), (F3, F4)
   WORD $0x4ea08421         // VADD V0.S4, V1.S4, V1.S4
   WORD $0x4ea08442         // VADD V0.S4, V2.S4, V2.S4
   WORD $0x4ea08463         // VADD V0.S4, V3.S4, V3.S4
   WORD $0x4ea08484         // VADD V0.S4, V4.S4, V4.S4
   WORD $0xad3f0961         // FSTPQ (F1, F2), -32(R11)
   WORD $0xac821163         // FSTPQ.P (F3, F4), 64(R11)
   WORD $0xf100418c         // SUBS $16, R12, R12
   WORD $0x54fffee1         // BNE -9(PC)
   WORD $0xeb0a013f         // CMP R10, R9
   WORD $0x54000100         // BEQ 8(PC)
   WORD $0x8b080809         // ADD R8<<2, R0, R9
   WORD $0xcb080048         // SUB R8, R2, R8
   WORD $0xb940012a         // MOVWU (R9), R10
   WORD $0x0b01014a         // ADDW R1, R10, R10
   WORD $0xb800452a         // MOVW.P R10, 4(R9)
   WORD $0xf1000508         // SUBS $1, R8, R8
   WORD $0x54ffff81         // BNE -4(PC)
   WORD $0xd65f03c0         // RET

TEXT ·_add_int32(SB), $0-32

	MOVD input1+0(FP),  R0
	MOVD input2+8(FP),  R1
	MOVD output+16(FP), R2
	MOVD size+24(FP),   R3

    WORD $0xd342fc66		// LSR $2, R3, R6
    WORD $0x710000df		// CMPW $0, R6
    WORD $0x540001ad		// BLE 13(PC)
    WORD $0x510004c5		// SUBW $1, R6, R5
    WORD $0xd2800004		// MOVD $0, R4
    WORD $0x910004a5		// ADD $1, R5, R5
    WORD $0xd37ceca5		// LSL $4, R5, R5
    WORD $0xd503201f		// NOOP
    WORD $0x3ce46800		// MOVD (R0)(R4), V0
    WORD $0x3ce46821		// MOVD (R1)(R4), V1
    WORD $0x4ea18400		// VADD V1.S4, V0.S4, V0.S4
    WORD $0x3ca46840		// MOVD V0, (R2)(R4)
    WORD $0x91004084		// ADD $16, R4, R4
    WORD $0xeb05009f		// CMP R5, R4
    WORD $0x54ffff41		// BNE -6(PC)
    WORD $0x531e74c6		// LSLW $2, R6, R6
    WORD $0x93407ccb		// SXTW R6, R11
    WORD $0xeb0b007f		// CMP R11, R3
    WORD $0x54000ee9		// BLS 119(PC)
    WORD $0x91001167		// ADD $4, R11, R7
    WORD $0xcb0b006a		// SUB R11, R3, R10
    WORD $0xd37ef4e7		// LSL $2, R7, R7
    WORD $0xd10040e9		// SUB $16, R7, R9
    WORD $0x8b070048		// ADD R7, R2, R8
    WORD $0x8b090025		// ADD R9, R1, R5
    WORD $0x8b070024		// ADD R7, R1, R4
    WORD $0xeb05011f		// CMP R5, R8
    WORD $0x8b09004c		// ADD R9, R2, R12
    WORD $0xfa448182		// CCMP HI, R12, R4, $2
    WORD $0x8b090004		// ADD R9, R0, R4
    WORD $0x1a9f37e5		// CSETW HS, R5
    WORD $0x8b070007		// ADD R7, R0, R7
    WORD $0xeb08009f		// CMP R8, R4
    WORD $0xaa0903e8		// MOVD R9, R8
    WORD $0xfa473182		// CCMP LO, R12, R7, $2
    WORD $0x1a9f37e7		// CSETW HS, R7
    WORD $0xf1001d5f		// CMP $7, R10
    WORD $0x0a0700a5		// ANDW R7, R5, R5
    WORD $0x1a9f97e7		// CSETW HI, R7
    WORD $0x6a0500ff		// TSTW R5, R7
    WORD $0x54000c40		// BEQ 98(PC)
    WORD $0xcb440be4		// NEG R4>>2, R4
    WORD $0xd1000467		// SUB $1, R3, R7
    WORD $0x92400485		// AND $3, R4, R5
    WORD $0xcb0b00e4		// SUB R11, R7, R4
    WORD $0x91000ca7		// ADD $3, R5, R7
    WORD $0xeb07009f		// CMP R7, R4
    WORD $0x540005c3		// BCC 46(PC)
    WORD $0x2a0603ec		// MOVW R6, R12
    WORD $0xb40002e5		// CBZ R5, 23(PC)
    WORD $0xb86b7804		// MOVWU (R0)(R11<<2), R4
    WORD $0x110004cc		// ADDW $1, R6, R12
    WORD $0xb86b7827		// MOVWU (R1)(R11<<2), R7
    WORD $0xf10004bf		// CMP $1, R5
    WORD $0x0b070084		// ADDW R7, R4, R4
    WORD $0xb82b7844		// MOVW R4, (R2)(R11<<2)
    WORD $0x93407d8b		// SXTW R12, R11
    WORD $0x540001e0		// BEQ 15(PC)
    WORD $0xb86b7804		// MOVWU (R0)(R11<<2), R4
    WORD $0x110008cc		// ADDW $2, R6, R12
    WORD $0xb86b7827		// MOVWU (R1)(R11<<2), R7
    WORD $0xf1000cbf		// CMP $3, R5
    WORD $0x0b070084		// ADDW R7, R4, R4
    WORD $0xb82b7844		// MOVW R4, (R2)(R11<<2)
    WORD $0x93407d8b		// SXTW R12, R11
    WORD $0x540000e1		// BNE 7(PC)
    WORD $0xb86b7804		// MOVWU (R0)(R11<<2), R4
    WORD $0x11000ccc		// ADDW $3, R6, R12
    WORD $0xb86b7827		// MOVWU (R1)(R11<<2), R7
    WORD $0x0b070084		// ADDW R7, R4, R4
    WORD $0xb82b7844		// MOVW R4, (R2)(R11<<2)
    WORD $0x93407d8b		// SXTW R12, R11
    WORD $0xcb050147		// SUB R5, R10, R7
    WORD $0x8b050925		// ADD R5<<2, R9, R5
    WORD $0x8b05000a		// ADD R5, R0, R10
    WORD $0x8b050029		// ADD R5, R1, R9
    WORD $0xd342fce8		// LSR $2, R7, R8
    WORD $0x8b050045		// ADD R5, R2, R5
    WORD $0xd2800004		// MOVD $0, R4
    WORD $0xd2800006		// MOVD $0, R6
    WORD $0x3ce46920		// MOVD (R9)(R4), V0
    WORD $0x910004c6		// ADD $1, R6, R6
    WORD $0x3ce46941		// MOVD (R10)(R4), V1
    WORD $0xeb06011f		// CMP R6, R8
    WORD $0x4ea18400		// VADD V1.S4, V0.S4, V0.S4
    WORD $0x3ca468a0		// MOVD V0, (R5)(R4)
    WORD $0x91004084		// ADD $16, R4, R4
    WORD $0x54ffff28		// BHI -7(PC)
    WORD $0x927ef4e4		// AND $-4, R7, R4
    WORD $0x8b04016b		// ADD R4, R11, R11
    WORD $0x0b040186		// ADDW R4, R12, R6
    WORD $0xeb0400ff		// CMP R4, R7
    WORD $0x540005a0		// BEQ 45(PC)
    WORD $0xb86b7805		// MOVWU (R0)(R11<<2), R5
    WORD $0x110004c4		// ADDW $1, R6, R4
    WORD $0xb86b7827		// MOVWU (R1)(R11<<2), R7
    WORD $0x93407c84		// SXTW R4, R4
    WORD $0x0b0700a5		// ADDW R7, R5, R5
    WORD $0xb82b7845		// MOVW R5, (R2)(R11<<2)
    WORD $0xeb03009f		// CMP R3, R4
    WORD $0x540004a2		// BCS 37(PC)
    WORD $0xb8647807		// MOVWU (R0)(R4<<2), R7
    WORD $0x110008c5		// ADDW $2, R6, R5
    WORD $0xb8647828		// MOVWU (R1)(R4<<2), R8
    WORD $0x0b0800e7		// ADDW R8, R7, R7
    WORD $0xb8247847		// MOVW R7, (R2)(R4<<2)
    WORD $0x93407ca4		// SXTW R5, R4
    WORD $0xeb04007f		// CMP R4, R3
    WORD $0x540003a9		// BLS 29(PC)
    WORD $0xb8647807		// MOVWU (R0)(R4<<2), R7
    WORD $0x11000cc5		// ADDW $3, R6, R5
    WORD $0xb8647828		// MOVWU (R1)(R4<<2), R8
    WORD $0x0b0800e7		// ADDW R8, R7, R7
    WORD $0xb8247847		// MOVW R7, (R2)(R4<<2)
    WORD $0x93407ca4		// SXTW R5, R4
    WORD $0xeb04007f		// CMP R4, R3
    WORD $0x540002a9		// BLS 21(PC)
    WORD $0xb8647807		// MOVWU (R0)(R4<<2), R7
    WORD $0x110010c5		// ADDW $4, R6, R5
    WORD $0xb8647828		// MOVWU (R1)(R4<<2), R8
    WORD $0x0b0800e7		// ADDW R8, R7, R7
    WORD $0xb8247847		// MOVW R7, (R2)(R4<<2)
    WORD $0x93407ca4		// SXTW R5, R4
    WORD $0xeb04007f		// CMP R4, R3
    WORD $0x540001a9		// BLS 13(PC)
    WORD $0xb8647805		// MOVWU (R0)(R4<<2), R5
    WORD $0x110014c6		// ADDW $5, R6, R6
    WORD $0xb8647827		// MOVWU (R1)(R4<<2), R7
    WORD $0x93407cc6		// SXTW R6, R6
    WORD $0x0b0700a5		// ADDW R7, R5, R5
    WORD $0xb8247845		// MOVW R5, (R2)(R4<<2)
    WORD $0xeb06007f		// CMP R6, R3
    WORD $0x540000a9		// BLS 5(PC)
    WORD $0xb8667800		// MOVWU (R0)(R6<<2), R0
    WORD $0xb8667821		// MOVWU (R1)(R6<<2), R1
    WORD $0x0b010000		// ADDW R1, R0, R0
    WORD $0xb8267840		// MOVW R0, (R2)(R6<<2)
    WORD $0xd65f03c0		// RET
    WORD $0xd37ef463		// LSL $2, R3, R3
    WORD $0xd503201f		// NOOP
    WORD $0xb8686804		// MOVWU (R0)(R8), R4
    WORD $0xb8686825		// MOVWU (R1)(R8), R5
    WORD $0x0b050084		// ADDW R5, R4, R4
    WORD $0xb8286844		// MOVW R4, (R2)(R8)
    WORD $0x91001108		// ADD $4, R8, R8
    WORD $0xeb03011f		// CMP R3, R8
    WORD $0x54ffff41		// BNE -6(PC)
    WORD $0xd65f03c0		// RET

TEXT ·_sub_int32(SB), $0-32

	MOVD input1+0(FP),  R0
	MOVD input2+8(FP),  R1
	MOVD output+16(FP), R2
	MOVD size+24(FP),   R3

    WORD $0xd342fc66		// LSR $2, R3, R6
    WORD $0x710000df		// CMPW $0, R6
    WORD $0x540001ad		// BLE 13(PC)
    WORD $0x510004c5		// SUBW $1, R6, R5
    WORD $0xd2800004		// MOVD $0, R4
    WORD $0x910004a5		// ADD $1, R5, R5
    WORD $0xd37ceca5		// LSL $4, R5, R5
    WORD $0xd503201f		// NOOP
    WORD $0x3ce46800		// MOVD (R0)(R4), V0
    WORD $0x3ce46821		// MOVD (R1)(R4), V1
    WORD $0x6ea18400		// VSUB V1.S4, V0.S4, V0.S4
    WORD $0x3ca46840		// MOVD V0, (R2)(R4)
    WORD $0x91004084		// ADD $16, R4, R4
    WORD $0xeb05009f		// CMP R5, R4
    WORD $0x54ffff41		// BNE -6(PC)
    WORD $0x531e74c6		// LSLW $2, R6, R6
    WORD $0x93407ccb		// SXTW R6, R11
    WORD $0xeb0b007f		// CMP R11, R3
    WORD $0x54000ee9		// BLS 119(PC)
    WORD $0x91001167		// ADD $4, R11, R7
    WORD $0xcb0b006a		// SUB R11, R3, R10
    WORD $0xd37ef4e7		// LSL $2, R7, R7
    WORD $0xd10040e9		// SUB $16, R7, R9
    WORD $0x8b070048		// ADD R7, R2, R8
    WORD $0x8b090025		// ADD R9, R1, R5
    WORD $0x8b070024		// ADD R7, R1, R4
    WORD $0xeb05011f		// CMP R5, R8
    WORD $0x8b09004c		// ADD R9, R2, R12
    WORD $0xfa448182		// CCMP HI, R12, R4, $2
    WORD $0x8b090004		// ADD R9, R0, R4
    WORD $0x1a9f37e5		// CSETW HS, R5
    WORD $0x8b070007		// ADD R7, R0, R7
    WORD $0xeb08009f		// CMP R8, R4
    WORD $0xaa0903e8		// MOVD R9, R8
    WORD $0xfa473182		// CCMP LO, R12, R7, $2
    WORD $0x1a9f37e7		// CSETW HS, R7
    WORD $0xf1001d5f		// CMP $7, R10
    WORD $0x0a0700a5		// ANDW R7, R5, R5
    WORD $0x1a9f97e7		// CSETW HI, R7
    WORD $0x6a0500ff		// TSTW R5, R7
    WORD $0x54000c40		// BEQ 98(PC)
    WORD $0xcb440be4		// NEG R4>>2, R4
    WORD $0xd1000467		// SUB $1, R3, R7
    WORD $0x92400485		// AND $3, R4, R5
    WORD $0xcb0b00e4		// SUB R11, R7, R4
    WORD $0x91000ca7		// ADD $3, R5, R7
    WORD $0xeb07009f		// CMP R7, R4
    WORD $0x540005c3		// BCC 46(PC)
    WORD $0x2a0603ec		// MOVW R6, R12
    WORD $0xb40002e5		// CBZ R5, 23(PC)
    WORD $0xb86b7804		// MOVWU (R0)(R11<<2), R4
    WORD $0x110004cc		// ADDW $1, R6, R12
    WORD $0xb86b7827		// MOVWU (R1)(R11<<2), R7
    WORD $0xf10004bf		// CMP $1, R5
    WORD $0x4b070084		// SUBW R7, R4, R4
    WORD $0xb82b7844		// MOVW R4, (R2)(R11<<2)
    WORD $0x93407d8b		// SXTW R12, R11
    WORD $0x540001e0		// BEQ 15(PC)
    WORD $0xb86b7804		// MOVWU (R0)(R11<<2), R4
    WORD $0x110008cc		// ADDW $2, R6, R12
    WORD $0xb86b7827		// MOVWU (R1)(R11<<2), R7
    WORD $0xf1000cbf		// CMP $3, R5
    WORD $0x4b070084		// SUBW R7, R4, R4
    WORD $0xb82b7844		// MOVW R4, (R2)(R11<<2)
    WORD $0x93407d8b		// SXTW R12, R11
    WORD $0x540000e1		// BNE 7(PC)
    WORD $0xb86b7804		// MOVWU (R0)(R11<<2), R4
    WORD $0x11000ccc		// ADDW $3, R6, R12
    WORD $0xb86b7827		// MOVWU (R1)(R11<<2), R7
    WORD $0x4b070084		// SUBW R7, R4, R4
    WORD $0xb82b7844		// MOVW R4, (R2)(R11<<2)
    WORD $0x93407d8b		// SXTW R12, R11
    WORD $0xcb050147		// SUB R5, R10, R7
    WORD $0x8b050925		// ADD R5<<2, R9, R5
    WORD $0x8b05000a		// ADD R5, R0, R10
    WORD $0x8b050029		// ADD R5, R1, R9
    WORD $0xd342fce8		// LSR $2, R7, R8
    WORD $0x8b050045		// ADD R5, R2, R5
    WORD $0xd2800004		// MOVD $0, R4
    WORD $0xd2800006		// MOVD $0, R6
    WORD $0x3ce46921		// MOVD (R9)(R4), V1
    WORD $0x910004c6		// ADD $1, R6, R6
    WORD $0x3ce46940		// MOVD (R10)(R4), V0
    WORD $0xeb06011f		// CMP R6, R8
    WORD $0x6ea18400		// VSUB V1.S4, V0.S4, V0.S4
    WORD $0x3ca468a0		// MOVD V0, (R5)(R4)
    WORD $0x91004084		// ADD $16, R4, R4
    WORD $0x54ffff28		// BHI -7(PC)
    WORD $0x927ef4e4		// AND $-4, R7, R4
    WORD $0x8b04016b		// ADD R4, R11, R11
    WORD $0x0b040186		// ADDW R4, R12, R6
    WORD $0xeb0400ff		// CMP R4, R7
    WORD $0x540005a0		// BEQ 45(PC)
    WORD $0xb86b7805		// MOVWU (R0)(R11<<2), R5
    WORD $0x110004c4		// ADDW $1, R6, R4
    WORD $0xb86b7827		// MOVWU (R1)(R11<<2), R7
    WORD $0x93407c84		// SXTW R4, R4
    WORD $0x4b0700a5		// SUBW R7, R5, R5
    WORD $0xb82b7845		// MOVW R5, (R2)(R11<<2)
    WORD $0xeb03009f		// CMP R3, R4
    WORD $0x540004a2		// BCS 37(PC)
    WORD $0xb8647807		// MOVWU (R0)(R4<<2), R7
    WORD $0x110008c5		// ADDW $2, R6, R5
    WORD $0xb8647828		// MOVWU (R1)(R4<<2), R8
    WORD $0x4b0800e7		// SUBW R8, R7, R7
    WORD $0xb8247847		// MOVW R7, (R2)(R4<<2)
    WORD $0x93407ca4		// SXTW R5, R4
    WORD $0xeb04007f		// CMP R4, R3
    WORD $0x540003a9		// BLS 29(PC)
    WORD $0xb8647807		// MOVWU (R0)(R4<<2), R7
    WORD $0x11000cc5		// ADDW $3, R6, R5
    WORD $0xb8647828		// MOVWU (R1)(R4<<2), R8
    WORD $0x4b0800e7		// SUBW R8, R7, R7
    WORD $0xb8247847		// MOVW R7, (R2)(R4<<2)
    WORD $0x93407ca4		// SXTW R5, R4
    WORD $0xeb04007f		// CMP R4, R3
    WORD $0x540002a9		// BLS 21(PC)
    WORD $0xb8647807		// MOVWU (R0)(R4<<2), R7
    WORD $0x110010c5		// ADDW $4, R6, R5
    WORD $0xb8647828		// MOVWU (R1)(R4<<2), R8
    WORD $0x4b0800e7		// SUBW R8, R7, R7
    WORD $0xb8247847		// MOVW R7, (R2)(R4<<2)
    WORD $0x93407ca4		// SXTW R5, R4
    WORD $0xeb04007f		// CMP R4, R3
    WORD $0x540001a9		// BLS 13(PC)
    WORD $0xb8647805		// MOVWU (R0)(R4<<2), R5
    WORD $0x110014c6		// ADDW $5, R6, R6
    WORD $0xb8647827		// MOVWU (R1)(R4<<2), R7
    WORD $0x93407cc6		// SXTW R6, R6
    WORD $0x4b0700a5		// SUBW R7, R5, R5
    WORD $0xb8247845		// MOVW R5, (R2)(R4<<2)
    WORD $0xeb06007f		// CMP R6, R3
    WORD $0x540000a9		// BLS 5(PC)
    WORD $0xb8667800		// MOVWU (R0)(R6<<2), R0
    WORD $0xb8667821		// MOVWU (R1)(R6<<2), R1
    WORD $0x4b010000		// SUBW R1, R0, R0
    WORD $0xb8267840		// MOVW R0, (R2)(R6<<2)
    WORD $0xd65f03c0		// RET
    WORD $0xd37ef463		// LSL $2, R3, R3
    WORD $0xd503201f		// NOOP
    WORD $0xb8686804		// MOVWU (R0)(R8), R4
    WORD $0xb8686825		// MOVWU (R1)(R8), R5
    WORD $0x4b050084		// SUBW R5, R4, R4
    WORD $0xb8286844		// MOVW R4, (R2)(R8)
    WORD $0x91001108		// ADD $4, R8, R8
    WORD $0xeb03011f		// CMP R3, R8
    WORD $0x54ffff41		// BNE -6(PC)
    WORD $0xd65f03c0		// RET

TEXT ·_scalar_mul_int32(SB), $0-32

   MOVD vec+0(FP),  R0
   MOVD constant+8(FP), R1
   MOVD size+16(FP), R2

   WORD $0xd342fc48                // LSR $2, R2, R8
   WORD $0x4e040c20                // VDUP R1, V0.S4
   WORD $0x7100051f                // CMPW $1, R8
   WORD $0x5400010b                // BLT 8(PC)
   WORD $0xd3428449                // UBFX $2, R2, $32, R9
   WORD $0xaa0003ea                // MOVD R0, R10
   WORD $0x3dc00141                // FMOVQ (R10), F1
   WORD $0x4ea09c21                // VMUL V0.S4, V1.S4, V1.S4
   WORD $0x3c810541                // FMOVQ.P F1, 16(R10)
   WORD $0xf1000529                // SUBS $1, R9, R9
   WORD $0x54ffff81                // BNE -4(PC)
   WORD $0x531e7508                // LSLW $2, R8, R8
   WORD $0x93407d0b                // SXTW R8, R11
   WORD $0xeb02017f                // CMP R2, R11
   WORD $0x540003a2                // BCS 29(PC)
   WORD $0xcb0b0049                // SUB R11, R2, R9
   WORD $0xaa0b03e8                // MOVD R11, R8
   WORD $0xf100413f                // CMP $16, R9
   WORD $0x54000243                // BCC 18(PC)
   WORD $0x927ced2a                // AND $-16, R9, R10
   WORD $0x8b0b0148                // ADD R11, R10, R8
   WORD $0x8b0b080b                // ADD R11<<2, R0, R11
   WORD $0x9100816b                // ADD $32, R11, R11
   WORD $0xaa0a03ec                // MOVD R10, R12
   WORD $0xad7f0961                // FLDPQ -32(R11), (F1, F2)
   WORD $0xad401163                // FLDPQ (R11), (F3, F4)
   WORD $0x4ea09c21                // VMUL V0.S4, V1.S4, V1.S4
   WORD $0x4ea09c42                // VMUL V0.S4, V2.S4, V2.S4
   WORD $0x4ea09c63                // VMUL V0.S4, V3.S4, V3.S4
   WORD $0x4ea09c84                // VMUL V0.S4, V4.S4, V4.S4
   WORD $0xad3f0961                // FSTPQ (F1, F2), -32(R11)
   WORD $0xac821163                // FSTPQ.P (F3, F4), 64(R11)
   WORD $0xf100418c                // SUBS $16, R12, R12
   WORD $0x54fffee1                // BNE -9(PC)
   WORD $0xeb0a013f                // CMP R10, R9
   WORD $0x54000100                // BEQ 8(PC)
   WORD $0x8b080809                // ADD R8<<2, R0, R9
   WORD $0xcb080048                // SUB R8, R2, R8
   WORD $0xb940012a                // MOVWU (R9), R10
   WORD $0x1b017d4a                // MULW R1, R10, R10
   WORD $0xb800452a                // MOVW.P R10, 4(R9)
   WORD $0xf1000508                // SUBS $1, R8, R8
   WORD $0x54ffff81                // BNE -4(PC)
   WORD $0xd65f03c0                // RET

TEXT ·_mul_int32(SB), $0-32

	MOVD input1+0(FP),  R0
	MOVD input2+8(FP),  R1
	MOVD output+16(FP), R2
	MOVD size+24(FP),   R3

    WORD $0xd342fc66		// LSR $2, R3, R6
    WORD $0x710000df		// CMPW $0, R6
    WORD $0x540001ad		// BLE 13(PC)
    WORD $0x510004c5		// SUBW $1, R6, R5
    WORD $0xd2800004		// MOVD $0, R4
    WORD $0x910004a5		// ADD $1, R5, R5
    WORD $0xd37ceca5		// LSL $4, R5, R5
    WORD $0xd503201f		// NOOP
    WORD $0x3ce46800		// MOVD (R0)(R4), V0
    WORD $0x3ce46821		// MOVD (R1)(R4), V1
    WORD $0x4ea19c00		// VMUL V1.S4, V0.S4, V0.S4
    WORD $0x3ca46840		// MOVD V0, (R2)(R4)
    WORD $0x91004084		// ADD $16, R4, R4
    WORD $0xeb05009f		// CMP R5, R4
    WORD $0x54ffff41		// BNE -6(PC)
    WORD $0x531e74c6		// LSLW $2, R6, R6
    WORD $0x93407ccb		// SXTW R6, R11
    WORD $0xeb0b007f		// CMP R11, R3
    WORD $0x54000ee9		// BLS 119(PC)
    WORD $0x91001167		// ADD $4, R11, R7
    WORD $0xcb0b006a		// SUB R11, R3, R10
    WORD $0xd37ef4e7		// LSL $2, R7, R7
    WORD $0xd10040e9		// SUB $16, R7, R9
    WORD $0x8b070048		// ADD R7, R2, R8
    WORD $0x8b090025		// ADD R9, R1, R5
    WORD $0x8b070024		// ADD R7, R1, R4
    WORD $0xeb05011f		// CMP R5, R8
    WORD $0x8b09004c		// ADD R9, R2, R12
    WORD $0xfa448182		// CCMP HI, R12, R4, $2
    WORD $0x8b090004		// ADD R9, R0, R4
    WORD $0x1a9f37e5		// CSETW HS, R5
    WORD $0x8b070007		// ADD R7, R0, R7
    WORD $0xeb08009f		// CMP R8, R4
    WORD $0xaa0903e8		// MOVD R9, R8
    WORD $0xfa473182		// CCMP LO, R12, R7, $2
    WORD $0x1a9f37e7		// CSETW HS, R7
    WORD $0xf1001d5f		// CMP $7, R10
    WORD $0x0a0700a5		// ANDW R7, R5, R5
    WORD $0x1a9f97e7		// CSETW HI, R7
    WORD $0x6a0500ff		// TSTW R5, R7
    WORD $0x54000c40		// BEQ 98(PC)
    WORD $0xcb440be4		// NEG R4>>2, R4
    WORD $0xd1000467		// SUB $1, R3, R7
    WORD $0x92400485		// AND $3, R4, R5
    WORD $0xcb0b00e4		// SUB R11, R7, R4
    WORD $0x91000ca7		// ADD $3, R5, R7
    WORD $0xeb07009f		// CMP R7, R4
    WORD $0x540005c3		// BCC 46(PC)
    WORD $0x2a0603ec		// MOVW R6, R12
    WORD $0xb40002e5		// CBZ R5, 23(PC)
    WORD $0xb86b7804		// MOVWU (R0)(R11<<2), R4
    WORD $0x110004cc		// ADDW $1, R6, R12
    WORD $0xb86b7827		// MOVWU (R1)(R11<<2), R7
    WORD $0xf10004bf		// CMP $1, R5
    WORD $0x1b077c84		// MULW R7, R4, R4
    WORD $0xb82b7844		// MOVW R4, (R2)(R11<<2)
    WORD $0x93407d8b		// SXTW R12, R11
    WORD $0x540001e0		// BEQ 15(PC)
    WORD $0xb86b7804		// MOVWU (R0)(R11<<2), R4
    WORD $0x110008cc		// ADDW $2, R6, R12
    WORD $0xb86b7827		// MOVWU (R1)(R11<<2), R7
    WORD $0xf1000cbf		// CMP $3, R5
    WORD $0x1b077c84		// MULW R7, R4, R4
    WORD $0xb82b7844		// MOVW R4, (R2)(R11<<2)
    WORD $0x93407d8b		// SXTW R12, R11
    WORD $0x540000e1		// BNE 7(PC)
    WORD $0xb86b7804		// MOVWU (R0)(R11<<2), R4
    WORD $0x11000ccc		// ADDW $3, R6, R12
    WORD $0xb86b7827		// MOVWU (R1)(R11<<2), R7
    WORD $0x1b077c84		// MULW R7, R4, R4
    WORD $0xb82b7844		// MOVW R4, (R2)(R11<<2)
    WORD $0x93407d8b		// SXTW R12, R11
    WORD $0xcb050147		// SUB R5, R10, R7
    WORD $0x8b050925		// ADD R5<<2, R9, R5
    WORD $0x8b05000a		// ADD R5, R0, R10
    WORD $0x8b050029		// ADD R5, R1, R9
    WORD $0xd342fce8		// LSR $2, R7, R8
    WORD $0x8b050045		// ADD R5, R2, R5
    WORD $0xd2800004		// MOVD $0, R4
    WORD $0xd2800006		// MOVD $0, R6
    WORD $0x3ce46920		// MOVD (R9)(R4), V0
    WORD $0x910004c6		// ADD $1, R6, R6
    WORD $0x3ce46941		// MOVD (R10)(R4), V1
    WORD $0xeb06011f		// CMP R6, R8
    WORD $0x4ea19c00		// VMUL V1.S4, V0.S4, V0.S4
    WORD $0x3ca468a0		// MOVD V0, (R5)(R4)
    WORD $0x91004084		// ADD $16, R4, R4
    WORD $0x54ffff28		// BHI -7(PC)
    WORD $0x927ef4e4		// AND $-4, R7, R4
    WORD $0x8b04016b		// ADD R4, R11, R11
    WORD $0x0b040186		// ADDW R4, R12, R6
    WORD $0xeb0400ff		// CMP R4, R7
    WORD $0x540005a0		// BEQ 45(PC)
    WORD $0xb86b7805		// MOVWU (R0)(R11<<2), R5
    WORD $0x110004c4		// ADDW $1, R6, R4
    WORD $0xb86b7827		// MOVWU (R1)(R11<<2), R7
    WORD $0x93407c84		// SXTW R4, R4
    WORD $0xeb03009f		// CMP R3, R4
    WORD $0x1b077ca5		// MULW R7, R5, R5
    WORD $0xb82b7845		// MOVW R5, (R2)(R11<<2)
    WORD $0x540004a2		// BCS 37(PC)
    WORD $0xb8647807		// MOVWU (R0)(R4<<2), R7
    WORD $0x110008c5		// ADDW $2, R6, R5
    WORD $0xb8647828		// MOVWU (R1)(R4<<2), R8
    WORD $0x93407ca5		// SXTW R5, R5
    WORD $0xeb05007f		// CMP R5, R3
    WORD $0x1b087ce7		// MULW R8, R7, R7
    WORD $0xb8247847		// MOVW R7, (R2)(R4<<2)
    WORD $0x540003a9		// BLS 29(PC)
    WORD $0xb8657807		// MOVWU (R0)(R5<<2), R7
    WORD $0x11000cc4		// ADDW $3, R6, R4
    WORD $0xb8657828		// MOVWU (R1)(R5<<2), R8
    WORD $0x93407c84		// SXTW R4, R4
    WORD $0xeb04007f		// CMP R4, R3
    WORD $0x1b087ce7		// MULW R8, R7, R7
    WORD $0xb8257847		// MOVW R7, (R2)(R5<<2)
    WORD $0x540002a9		// BLS 21(PC)
    WORD $0xb8647807		// MOVWU (R0)(R4<<2), R7
    WORD $0x110010c5		// ADDW $4, R6, R5
    WORD $0xb8647828		// MOVWU (R1)(R4<<2), R8
    WORD $0x93407ca5		// SXTW R5, R5
    WORD $0xeb05007f		// CMP R5, R3
    WORD $0x1b087ce7		// MULW R8, R7, R7
    WORD $0xb8247847		// MOVW R7, (R2)(R4<<2)
    WORD $0x540001a9		// BLS 13(PC)
    WORD $0xb8657804		// MOVWU (R0)(R5<<2), R4
    WORD $0x110014c6		// ADDW $5, R6, R6
    WORD $0xb8657827		// MOVWU (R1)(R5<<2), R7
    WORD $0x93407cc6		// SXTW R6, R6
    WORD $0xeb06007f		// CMP R6, R3
    WORD $0x1b077c84		// MULW R7, R4, R4
    WORD $0xb8257844		// MOVW R4, (R2)(R5<<2)
    WORD $0x540000a9		// BLS 5(PC)
    WORD $0xb8667800		// MOVWU (R0)(R6<<2), R0
    WORD $0xb8667821		// MOVWU (R1)(R6<<2), R1
    WORD $0x1b017c00		// MULW R1, R0, R0
    WORD $0xb8267840		// MOVW R0, (R2)(R6<<2)
    WORD $0xd65f03c0		// RET
    WORD $0xd37ef463		// LSL $2, R3, R3
    WORD $0xd503201f		// NOOP
    WORD $0xb8686804		// MOVWU (R0)(R8), R4
    WORD $0xb8686825		// MOVWU (R1)(R8), R5
    WORD $0x1b057c84		// MULW R5, R4, R4
    WORD $0xb8286844		// MOVW R4, (R2)(R8)
    WORD $0x91001108		// ADD $4, R8, R8
    WORD $0xeb03011f		// CMP R3, R8
    WORD $0x54ffff41		// BNE -6(PC)
    WORD $0xd65f03c0		// RET

TEXT ·_inc_int32_sve(SB), $0-32

   MOVD vec+0(FP),  R0
   MOVD constant+8(FP), R1
   MOVD size+16(FP), R2

  WORD $0x7100045f		// CMPW $1, R2
  WORD $0x5400018b		// BLT 12(PC)
  WORD $0x52800008		// MOVW $0, R8
  WORD $0x05a03820		// ?
  WORD $0x04a0e3e9		// ?
  WORD $0x25a20500		// ?
  WORD $0x93407d0a		// SXTW R8, R10
  WORD $0xa54a4001		// ?
  WORD $0x04800001		// ?
  WORD $0xe54a4001		// ?
  WORD $0x0b090108		// ADDW R9, R8, R8
  WORD $0x6b02011f		// CMPW R2, R8
  WORD $0x54ffff2b		// BLT -7(PC)
  WORD $0xd65f03c0		// RET

TEXT ·_scalar_mul_int32_sve(SB), $0-32

   MOVD vec+0(FP),  R0
   MOVD constant+8(FP), R1
   MOVD size+16(FP), R2

   WORD $0x7100045f		// CMPW $1, R2
   WORD $0x5400018b		// BLT 12(PC)
   WORD $0x52800008		// MOVW $0, R8
   WORD $0x05a03820		// ?
   WORD $0x04a0e3e9		// ?
   WORD $0x25a20500		// ?
   WORD $0x93407d0a		// SXTW R8, R10
   WORD $0xa54a4001		// ?
   WORD $0x04900001		// ?
   WORD $0xe54a4001		// ?
   WORD $0x0b090108		// ADDW R9, R8, R8
   WORD $0x6b02011f		// CMPW R2, R8
   WORD $0x54ffff2b		// BLT -7(PC)
   WORD $0xd65f03c0		// RET
