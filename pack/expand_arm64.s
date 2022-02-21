#include "textflag.h"

TEXT ·_expand_mask_arm(SB), $0-16
  MOVD input+0(FP),  R0
  MOVD output+8(FP), R1

WORD $0x4f00e400         // VMOVI $0, V0.B16                     // movi v0.16b, #0x0
WORD $0x0d408400         // VLD1 (R0), V0.D[0]                   // ld1 {v0.d}[0], [x0]
MOVD $·index0(SB), R8    // ADRP 0(PC), R8                       // adrp x8, .+0x0
WORD $0x3dc00101         // FMOVQ (R8), F1                       // ldr q1, [x8]
WORD $0x4e010001         // VTBL V1.B16, [V0.B16], V1.B16        // tbl v1.16b, {v0.16b}, v1.16b
MOVD $·bit_mask(SB), R8  // ADRP 0(PC), R8                       // adrp x8, .+0x0
WORD $0x3dc00102         // FMOVQ (R8), F2                       // ldr q2, [x8]
WORD $0x4ea21c21         // VORR V2.B16, V1.B16, V1.B16          // orr v1.16b, v1.16b, v2.16b
WORD $0x6f07e7e3         // VMOVI $-1, V3.D2                     // movi v3.2d, #0xffffffffffffffff
WORD $0x6e238c21         // VCMEQ V3.B16, V1.B16, V1.B16         // cmeq v1.16b, v1.16b, v3.16b
MOVD $·index1(SB), R8    // ADRP 0(PC), R8                       // adrp x8, .+0x0
WORD $0x3dc00104         // FMOVQ (R8), F4                       // ldr q4, [x8]
WORD $0x4e040004         // VTBL V4.B16, [V0.B16], V4.B16        // tbl v4.16b, {v0.16b}, v4.16b
WORD $0x4ea21c84         // VORR V2.B16, V4.B16, V4.B16          // orr v4.16b, v4.16b, v2.16b
WORD $0x6e238c84         // VCMEQ V3.B16, V4.B16, V4.B16         // cmeq v4.16b, v4.16b, v3.16b
WORD $0xad001021         // FSTPQ (F1, F4), (R1)                 // stp q1, q4, [x1]
MOVD $·index2(SB), R8    // ADRP 0(PC), R8                       // adrp x8, .+0x0
WORD $0x3dc00101         // FMOVQ (R8), F1                       // ldr q1, [x8]
WORD $0x4e010001         // VTBL V1.B16, [V0.B16], V1.B16        // tbl v1.16b, {v0.16b}, v1.16b
MOVD $·index3(SB), R8    // ADRP 0(PC), R8                       // adrp x8, .+0x0
WORD $0x3dc00104         // FMOVQ (R8), F4                       // ldr q4, [x8]
WORD $0x4e040000         // VTBL V4.B16, [V0.B16], V0.B16        // tbl v0.16b, {v0.16b}, v4.16b
WORD $0x4ea21c21         // VORR V2.B16, V1.B16, V1.B16          // orr v1.16b, v1.16b, v2.16b
WORD $0x6e238c21         // VCMEQ V3.B16, V1.B16, V1.B16         // cmeq v1.16b, v1.16b, v3.16b
WORD $0x4ea21c00         // VORR V2.B16, V0.B16, V0.B16          // orr v0.16b, v0.16b, v2.16b
WORD $0x6e238c00         // VCMEQ V3.B16, V0.B16, V0.B16         // cmeq v0.16b, v0.16b, v3.16b
WORD $0xad010021         // FSTPQ (F1, F0), 32(R1)               // stp q1, q0, [x1,#32]
WORD $0xd65f03c0         // RET                                  // ret

DATA    ·bit_mask+0x00(SB)/8, $0x7fbfdfeff7fbfdfe
DATA    ·bit_mask+0x08(SB)/8, $0x7fbfdfeff7fbfdfe
GLOBL   ·bit_mask(SB),  NOPTR|RODATA, $16

DATA    ·index0+0x00(SB)/8, $0x0000000000000000
DATA    ·index0+0x08(SB)/8, $0x0101010101010101
GLOBL   ·index0(SB),  NOPTR|RODATA, $16

DATA    ·index1+0x00(SB)/8, $0x0202020202020202
DATA    ·index1+0x08(SB)/8, $0x0303030303030303
GLOBL   ·index1(SB),  NOPTR|RODATA, $16

DATA    ·index2+0x00(SB)/8, $0x0404040404040404
DATA    ·index2+0x08(SB)/8, $0x0505050505050505
GLOBL   ·index2(SB),  NOPTR|RODATA, $16

DATA    ·index3+0x00(SB)/8, $0x0606060606060606
DATA    ·index3+0x08(SB)/8, $0x0707070707070707
GLOBL   ·index3(SB),  NOPTR|RODATA, $16


