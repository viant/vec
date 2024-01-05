#include "textflag.h"

TEXT ·_msb_neon(SB), $0-24
  MOVD in+0(FP),   R0
  MOVD size+8(FP), R1
  MOVD out+16(FP), R2

  WORD $0x92800008                // MOVD $-1, R8
  WORD $0xf9000048                // MOVD R8, (R2)
  WORD $0x12000429                // ANDW $3, R1, R9
  WORD $0x531a6428                // LSLW $6, R1, R8
  WORD $0x110ffd08                // ADDW $1023, R8, R8
  WORD $0x5100102a                // SUBW $4, R1, R10
  WORD $0x8b2acc0a                // ADD R10.SXTW<<3, R0, R10
  WORD $0x9100814c                // ADD $32, R10, R12
  WORD $0xaa0103eb                // MOVD R1, R11
  WORD $0x5100116b                // SUBW $4, R11, R11
  WORD $0x6b09017f                // CMPW R9, R11
  WORD $0x5400042b                // BLT 33(PC)
  WORD $0x51040108                // SUBW $256, R8, R8
  WORD $0xd100818a                // SUB $32, R12, R10
  WORD $0xad7f0181                // FLDPQ -32(R12), (F1, F0)
  WORD $0x4ea11c02                // VORR V1.B16, V0.B16, V2.B16
  WORD $0x6eb0a842                // VUMAXV V2.S4, V2
  WORD $0x1e26004d                // FMOVS F2, R13
  WORD $0xaa0a03ec                // MOVD R10, R12
  WORD $0x34fffecd                // CBZW R13, -10(PC)
  WORD $0x4ea09821                // VCMEQ $0, V1.S4, V1.S4
  WORD $0x4f000422                // VMOVI $1, V2.S4
  WORD $0x4e221c21                // VAND V2.B16, V1.B16, V1.B16
  MOVD $·shift(SB), R9            // ADRP 0(PC), R9
  WORD $0x3dc00123                // FMOVQ (R9), F3
  WORD $0x6ea34421                // VUSHL V3.S4, V1.S4, V1.S4
  WORD $0x4eb1b821                // VADDV V1.S4, V1
  WORD $0x1e260029                // FMOVS F1, R9
  WORD $0x4ea09800                // VCMEQ $0, V0.S4, V0.S4
  WORD $0x4e221c00                // VAND V2.B16, V0.B16, V0.B16
  WORD $0x6ea34400                // VUSHL V3.S4, V0.S4, V0.S4
  WORD $0x4eb1b800                // VADDV V0.S4, V0
  WORD $0x1e26000b                // FMOVS F0, R11
  WORD $0x2a0b1129                // ORRW R11<<4, R9, R9
  WORD $0x2a2903e9                // MVNW R9, R9
  WORD $0x12001d29                // ANDW $255, R9, R9
  WORD $0x5ac01129                // CLZW R9, R9
  WORD $0x52800f8b                // MOVW $124, R11
  WORD $0x4b09096b                // SUBW R9<<2, R11, R11
  WORD $0xb86b494a                // MOVWU (R10)(R11.UXTW), R10
  WORD $0x5ac0114a                // CLZW R10, R10
  WORD $0x0b091549                // ADDW R9<<5, R10, R9
  WORD $0x4b090108                // SUBW R9, R8, R8
  WORD $0x1400000b                // JMP 11(PC)
  WORD $0xf240042a                // ANDS $3, R1, R10
  WORD $0x54000140                // BEQ 10(PC)
  WORD $0xd1000548                // SUB $1, R10, R8
  WORD $0xf8685809                // MOVD (R0)(R8.UXTW<<3), R9
  WORD $0xb4000109                // CBZ R9, 8(PC)
  WORD $0x528007ea                // MOVW $63, R10
  WORD $0x331a650a                // BFIW $6, R8, $26, R10
  WORD $0xdac01128                // CLZ R9, R8
  WORD $0x4b080148                // SUBW R8, R10, R8
  WORD $0x93407d08                // SXTW R8, R8
  WORD $0xf9000048                // MOVD R8, (R2)
  WORD $0xd65f03c0                // RET
  WORD $0xf1000948                // SUBS $2, R10, R8
  WORD $0x54ffffc3                // BCC -2(PC)
  WORD $0xf8685809                // MOVD (R0)(R8.UXTW<<3), R9
  WORD $0xb5fffec9                // CBNZ R9, -10(PC)
  WORD $0xf1000d48                // SUBS $3, R10, R8
  WORD $0x54ffff41                // BNE -6(PC)
  WORD $0xf8685809                // MOVD (R0)(R8.UXTW<<3), R9
  WORD $0xb5fffe49                // CBNZ R9, -14(PC)
  WORD $0x17fffff7                // JMP -9(PC)

 // DATA    ·shift+0x00(SB)/4, $0x0
 // DATA    ·shift+0x04(SB)/4, $0x1
 // DATA    ·shift+0x08(SB)/4, $0x2
//  DATA    ·shift+0x0c(SB)/4, $0x3
//  GLOBL   ·shift(SB),  NOPTR|RODATA, $16
