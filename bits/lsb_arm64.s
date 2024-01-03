#include "textflag.h"

TEXT ·_lsb_neon(SB), $0-24
  MOVD in+0(FP),   R0
  MOVD size+8(FP), R1
  MOVD out+16(FP), R2

  WORD $0x92800008                // MOVD $-1, R8
  WORD $0xf9000048                // MOVD R8, (R2)
  WORD $0xd342fc28                // LSR $2, R1, R8
  WORD $0x340001c8                // CBZW R8, 14(PC)
  WORD $0xd2800008                // MOVD $0, R8
  WORD $0xd3428429                // UBFX $2, R1, $32, R9
  WORD $0xd378dd2a                // LSL $8, R9, R10
  WORD $0xaa0003e9                // MOVD R0, R9
  WORD $0xad400121                // FLDPQ (R9), (F1, F0)
  WORD $0x4ea11c02                // VORR V1.B16, V0.B16, V2.B16
  WORD $0x6eb0a842                // VUMAXV V2.S4, V2
  WORD $0x1e26004b                // FMOVS F2, R11
  WORD $0x350001cb                // CBNZW R11, 14(PC)
  WORD $0xd1040108                // SUB $256, R8, R8
  WORD $0x91008129                // ADD $32, R9, R9
  WORD $0xab08015f                // CMN R8, R10
  WORD $0x54ffff01                // BNE -8(PC)
  WORD $0x927e7428                // AND $4294967292, R1, R8
  WORD $0xeb01011f                // CMP R1, R8
  WORD $0x540000c2                // BCS 6(PC)
  WORD $0xf8687809                // MOVD (R0)(R8<<3), R9
  WORD $0xb50003a9                // CBNZ R9, 29(PC)
  WORD $0x11000508                // ADDW $1, R8, R8
  WORD $0xeb01011f                // CMP R1, R8
  WORD $0x54ffff83                // BCC -4(PC)
  WORD $0xd65f03c0                // RET
  WORD $0x4ea09821                // VCMEQ $0, V1.S4, V1.S4
  WORD $0x4f000422                // VMOVI $1, V2.S4
  WORD $0x4e221c21                // VAND V2.B16, V1.B16, V1.B16
  MOVD $·shift(SB), R10           // ADRP 0(PC), R10
  WORD $0x3dc00143                // FMOVQ (R10), F3
  WORD $0x6ea34421                // VUSHL V3.S4, V1.S4, V1.S4
  WORD $0x4eb1b821                // VADDV V1.S4, V1
  WORD $0x1e26002a                // FMOVS F1, R10
  WORD $0x4ea09800                // VCMEQ $0, V0.S4, V0.S4
  WORD $0x4e221c00                // VAND V2.B16, V0.B16, V0.B16
  WORD $0x6ea34400                // VUSHL V3.S4, V0.S4, V0.S4
  WORD $0x4eb1b800                // VADDV V0.S4, V0
  WORD $0x1e26000b                // FMOVS F0, R11
  WORD $0x2a0b114a                // ORRW R11<<4, R10, R10
  WORD $0x2a2a03ea                // MVNW R10, R10
  WORD $0x5ac0014a                // RBITW R10, R10
  WORD $0x5ac0114a                // CLZW R10, R10
  WORD $0xb86a5929                // MOVWU (R9)(R10.UXTW<<2), R9
  WORD $0x5ac00129                // RBITW R9, R9
  WORD $0x5ac01129                // CLZW R9, R9
  WORD $0x0b0a1529                // ADDW R10<<5, R9, R9
  WORD $0x4b080128                // SUBW R8, R9, R8
  WORD $0xf9000048                // MOVD R8, (R2)
  WORD $0xd65f03c0                // RET
  WORD $0x531a6508                // LSLW $6, R8, R8
  WORD $0xdac00129                // RBIT R9, R9
  WORD $0xdac01129                // CLZ R9, R9
  WORD $0xaa090108                // ORR R9, R8, R8
  WORD $0xf9000048                // MOVD R8, (R2)
  WORD $0xd65f03c0                // RET

  DATA    ·shift+0x00(SB)/4, $0x0
  DATA    ·shift+0x04(SB)/4, $0x1
  DATA    ·shift+0x08(SB)/4, $0x2
  DATA    ·shift+0x0c(SB)/4, $0x3
  GLOBL   ·shift(SB),  NOPTR|RODATA, $16
