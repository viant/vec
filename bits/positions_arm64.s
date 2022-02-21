#include "textflag.h"

TEXT ·_populate_positions_neon(SB), NOSPLIT, $0
  MOVD input+0(FP), R0
  MOVD out+8(FP),   R1
  MOVD count+16(FP), R2


	CBZ R0, L1
	UBFIZ $4, R0, $8, R6
	MOVD $vecDecodeTable(SB), R14
	MOVD R14, R8
	LSR $8, R0, R5
	WORD $0x6f000403		// VMVNI $0, V3.S4
	MOVD $lengthTable(SB), R15
	MOVD R15, R7
	UBFIZ $4, R5, $8, R9
	WORD $0x3ce868c1		// MOVD (R6)(R8), V1
	ANDW $255, R0, R6
	WORD $0x4f0084e2		// VMOVI $7, V2.H8
	ANDW $255, R5, R5
	WORD $0x3ce86920		// MOVD (R9)(R8), V0
	LSR $16, R0, R10
	WORD $0x4e638421		// VADD V3.H8, V1.H8, V1.H8
	MOVBU (R7)(R6.SXTW), R6
	MOVBU (R7)(R5.SXTW), R7
	WORD $0x4e628400		// VADD V2.H8, V0.H8, V0.H8
	CMPW $4, R6
	WORD $0x5e080422		// VMOV V1.D[0], V2
	WORD $0x2f10a442		// VUXTL V2.H4, V2.S4
	WORD $0x3d800022		// MOVD V2, (R1)
	BLS L3
	WORD $0x5e180421		// VMOV V1.D[1], V1
	WORD $0x2f10a421		// VUXTL V1.H4, V1.S4
	WORD $0x3d800421		// MOVD V1, 16(R1)
L3:	UBFIZ $2, R6, $8, R5
	WORD $0x5e080401		// VMOV V0.D[0], V1
	ADD R5, R1, R11
	WORD $0x2f10a421		// VUXTL V1.H4, V1.S4
	WORD $0x3ca56821		// MOVD V1, (R1)(R5)
	CMPW $4, R7
	BLS L4
	WORD $0x5e180400		// VMOV V0.D[1], V0
	WORD $0x2f10a400		// VUXTL V0.H4, V0.S4
	WORD $0x3d800560		// MOVD V0, 16(R11)
L4:	MOVD R14, R9
	UBFIZ $4, R10, $8, R6
	LSR $24, R0, R5
	WORD $0x4f0085e3		// VMOVI $15, V3.H8
	MOVD R15, R8
	UBFIZ $4, R5, $8, R12
	ANDW $255, R10, R10
	WORD $0x3ce968c1		// MOVD (R6)(R9), V1
	UBFIZ $2, R7, $8, R7
	WORD $0x3ce96980		// MOVD (R12)(R9), V0
	ANDW $255, R5, R5
	WORD $0x4f0086e2		// VMOVI $23, V2.H8
	MOVBU (R8)(R10.SXTW), R6
	WORD $0x4e638421		// VADD V3.H8, V1.H8, V1.H8
	ADD R7, R11, R9
	LSR $32, R0, R10
	MOVBU (R8)(R5.SXTW), R8
	CMPW $4, R6
	WORD $0x4e628400		// VADD V2.H8, V0.H8, V0.H8
	WORD $0x5e080422		// VMOV V1.D[0], V2
	WORD $0x2f10a442		// VUXTL V2.H4, V2.S4
	WORD $0x3ca76962		// MOVD V2, (R11)(R7)
	BLS L5
	WORD $0x5e180421		// VMOV V1.D[1], V1
	WORD $0x2f10a421		// VUXTL V1.H4, V1.S4
	WORD $0x3d800521		// MOVD V1, 16(R9)
L5:	UBFIZ $2, R6, $8, R5
	WORD $0x5e080401		// VMOV V0.D[0], V1
	ADD R5, R9, R11
	WORD $0x2f10a421		// VUXTL V1.H4, V1.S4
	WORD $0x3ca56921		// MOVD V1, (R9)(R5)
	CMPW $4, R8
	BLS L6
	WORD $0x5e180400		// VMOV V0.D[1], V0
	WORD $0x2f10a400		// VUXTL V0.H4, V0.S4
	WORD $0x3d800560		// MOVD V0, 16(R11)
L6: MOVD R14, R9
	UBFIZ $4, R10, $8, R7
	LSR $40, R0, R5
	WORD $0x4f0087e3		// VMOVI $31, V3.H8
	MOVD R15, R6
	UBFIZ $4, R5, $8, R12
	ANDW $255, R10, R10
	WORD $0x3ce968e1		// MOVD (R7)(R9), V1
	UBFIZ $2, R8, $8, R8
	WORD $0x3ce96980		// MOVD (R12)(R9), V0
	ANDW $255, R5, R5
	WORD $0x4f0184e2		// VMOVI $39, V2.H8
	MOVBU (R6)(R10.SXTW), R7
	WORD $0x4e638421		// VADD V3.H8, V1.H8, V1.H8
	ADD R8, R11, R10
	LSR $48, R0, R9
	MOVBU (R6)(R5.SXTW), R6
	CMPW $4, R7
	WORD $0x4e628400		// VADD V2.H8, V0.H8, V0.H8
	WORD $0x5e080422		// VMOV V1.D[0], V2
	WORD $0x2f10a442		// VUXTL V2.H4, V2.S4
	WORD $0x3ca86962		// MOVD V2, (R11)(R8)
	BLS L7
	WORD $0x5e180421		// VMOV V1.D[1], V1
	WORD $0x2f10a421		// VUXTL V1.H4, V1.S4
	WORD $0x3d800541		// MOVD V1, 16(R10)
L7:	UBFIZ $2, R7, $8, R5
	WORD $0x5e080401		// VMOV V0.D[0], V1
	ADD R5, R10, R8
	WORD $0x2f10a421		// VUXTL V1.H4, V1.S4
	WORD $0x3ca56941		// MOVD V1, (R10)(R5)
	CMPW $4, R6
	BLS L8
	WORD $0x5e180400		// VMOV V0.D[1], V0
	WORD $0x2f10a400		// VUXTL V0.H4, V0.S4
	WORD $0x3d800500		// MOVD V0, 16(R8)
L8:	UBFIZ $4, R9, $8, R5
	MOVD R14, R4
	LSR $56, R0, R0
	WORD $0x4f0185e3		// VMOVI $47, V3.H8
	MOVD R15, R3
	LSL $4, R0, R10
	ANDW $255, R9, R9
	WORD $0x3ce468a1		// MOVD (R5)(R4), V1
	UBFIZ $2, R6, $8, R6
	WORD $0x4f0186e2		// VMOVI $55, V2.H8
	MOVBU (R3)(R9.SXTW), R5
	WORD $0x3ce46940		// MOVD (R10)(R4), V0
	ADD R6, R8, R7
	WORD $0x4e638421		// VADD V3.H8, V1.H8, V1.H8
	MOVBU (R3)(R0), R0
	CMPW $4, R5
	WORD $0x4e628400		// VADD V2.H8, V0.H8, V0.H8
	WORD $0x5e080422		// VMOV V1.D[0], V2
	WORD $0x2f10a442		// VUXTL V2.H4, V2.S4
	WORD $0x3ca66902		// MOVD V2, (R8)(R6)
	BLS L9
	WORD $0x5e180421		// VMOV V1.D[1], V1
	WORD $0x2f10a421		// VUXTL V1.H4, V1.S4
	WORD $0x3d8004e1		// MOVD V1, 16(R7)
L9:	UBFIZ $2, R5, $8, R3
	WORD $0x5e080401		// VMOV V0.D[0], V1
	ADD R3, R7, R4
	WORD $0x2f10a421		// VUXTL V1.H4, V1.S4
	WORD $0x3ca368e1		// MOVD V1, (R7)(R3)
	CMPW $4, R0
	BLS L10
	WORD $0x5e180400		// VMOV V0.D[1], V0
	WORD $0x2f10a400		// VUXTL V0.H4, V0.S4
	WORD $0x3d800480		// MOVD V0, 16(R4)
L10:	ADD R0.UXTB<<2, R4, R0
L1: 	SUB R1, R0, R0
	ASR $2, R0, R0
	MOVW R0, (R2)
	RET

        DATA vecDecodeTable+0x000(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x008(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x010(SB)/8, $0x0000000000000001
        DATA vecDecodeTable+0x018(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x020(SB)/8, $0x0000000000000002
        DATA vecDecodeTable+0x028(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x030(SB)/8, $0x0000000000020001
        DATA vecDecodeTable+0x038(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x040(SB)/8, $0x0000000000000003
        DATA vecDecodeTable+0x048(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x050(SB)/8, $0x0000000000030001
        DATA vecDecodeTable+0x058(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x060(SB)/8, $0x0000000000030002
        DATA vecDecodeTable+0x068(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x070(SB)/8, $0x0000000300020001
        DATA vecDecodeTable+0x078(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x080(SB)/8, $0x0000000000000004
        DATA vecDecodeTable+0x088(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x090(SB)/8, $0x0000000000040001
        DATA vecDecodeTable+0x098(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x0a0(SB)/8, $0x0000000000040002
        DATA vecDecodeTable+0x0a8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x0b0(SB)/8, $0x0000000400020001
        DATA vecDecodeTable+0x0b8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x0c0(SB)/8, $0x0000000000040003
        DATA vecDecodeTable+0x0c8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x0d0(SB)/8, $0x0000000400030001
        DATA vecDecodeTable+0x0d8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x0e0(SB)/8, $0x0000000400030002
        DATA vecDecodeTable+0x0e8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x0f0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0x0f8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x100(SB)/8, $0x0000000000000005
        DATA vecDecodeTable+0x108(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x110(SB)/8, $0x0000000000050001
        DATA vecDecodeTable+0x118(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x120(SB)/8, $0x0000000000050002
        DATA vecDecodeTable+0x128(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x130(SB)/8, $0x0000000500020001
        DATA vecDecodeTable+0x138(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x140(SB)/8, $0x0000000000050003
        DATA vecDecodeTable+0x148(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x150(SB)/8, $0x0000000500030001
        DATA vecDecodeTable+0x158(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x160(SB)/8, $0x0000000500030002
        DATA vecDecodeTable+0x168(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x170(SB)/8, $0x0005000300020001
        DATA vecDecodeTable+0x178(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x180(SB)/8, $0x0000000000050004
        DATA vecDecodeTable+0x188(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x190(SB)/8, $0x0000000500040001
        DATA vecDecodeTable+0x198(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x1a0(SB)/8, $0x0000000500040002
        DATA vecDecodeTable+0x1a8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x1b0(SB)/8, $0x0005000400020001
        DATA vecDecodeTable+0x1b8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x1c0(SB)/8, $0x0000000500040003
        DATA vecDecodeTable+0x1c8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x1d0(SB)/8, $0x0005000400030001
        DATA vecDecodeTable+0x1d8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x1e0(SB)/8, $0x0005000400030002
        DATA vecDecodeTable+0x1e8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x1f0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0x1f8(SB)/8, $0x0000000000000005
        DATA vecDecodeTable+0x200(SB)/8, $0x0000000000000006
        DATA vecDecodeTable+0x208(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x210(SB)/8, $0x0000000000060001
        DATA vecDecodeTable+0x218(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x220(SB)/8, $0x0000000000060002
        DATA vecDecodeTable+0x228(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x230(SB)/8, $0x0000000600020001
        DATA vecDecodeTable+0x238(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x240(SB)/8, $0x0000000000060003
        DATA vecDecodeTable+0x248(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x250(SB)/8, $0x0000000600030001
        DATA vecDecodeTable+0x258(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x260(SB)/8, $0x0000000600030002
        DATA vecDecodeTable+0x268(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x270(SB)/8, $0x0006000300020001
        DATA vecDecodeTable+0x278(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x280(SB)/8, $0x0000000000060004
        DATA vecDecodeTable+0x288(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x290(SB)/8, $0x0000000600040001
        DATA vecDecodeTable+0x298(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x2a0(SB)/8, $0x0000000600040002
        DATA vecDecodeTable+0x2a8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x2b0(SB)/8, $0x0006000400020001
        DATA vecDecodeTable+0x2b8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x2c0(SB)/8, $0x0000000600040003
        DATA vecDecodeTable+0x2c8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x2d0(SB)/8, $0x0006000400030001
        DATA vecDecodeTable+0x2d8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x2e0(SB)/8, $0x0006000400030002
        DATA vecDecodeTable+0x2e8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x2f0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0x2f8(SB)/8, $0x0000000000000006
        DATA vecDecodeTable+0x300(SB)/8, $0x0000000000060005
        DATA vecDecodeTable+0x308(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x310(SB)/8, $0x0000000600050001
        DATA vecDecodeTable+0x318(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x320(SB)/8, $0x0000000600050002
        DATA vecDecodeTable+0x328(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x330(SB)/8, $0x0006000500020001
        DATA vecDecodeTable+0x338(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x340(SB)/8, $0x0000000600050003
        DATA vecDecodeTable+0x348(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x350(SB)/8, $0x0006000500030001
        DATA vecDecodeTable+0x358(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x360(SB)/8, $0x0006000500030002
        DATA vecDecodeTable+0x368(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x370(SB)/8, $0x0005000300020001
        DATA vecDecodeTable+0x378(SB)/8, $0x0000000000000006
        DATA vecDecodeTable+0x380(SB)/8, $0x0000000600050004
        DATA vecDecodeTable+0x388(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x390(SB)/8, $0x0006000500040001
        DATA vecDecodeTable+0x398(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x3a0(SB)/8, $0x0006000500040002
        DATA vecDecodeTable+0x3a8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x3b0(SB)/8, $0x0005000400020001
        DATA vecDecodeTable+0x3b8(SB)/8, $0x0000000000000006
        DATA vecDecodeTable+0x3c0(SB)/8, $0x0006000500040003
        DATA vecDecodeTable+0x3c8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x3d0(SB)/8, $0x0005000400030001
        DATA vecDecodeTable+0x3d8(SB)/8, $0x0000000000000006
        DATA vecDecodeTable+0x3e0(SB)/8, $0x0005000400030002
        DATA vecDecodeTable+0x3e8(SB)/8, $0x0000000000000006
        DATA vecDecodeTable+0x3f0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0x3f8(SB)/8, $0x0000000000060005
        DATA vecDecodeTable+0x400(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x408(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x410(SB)/8, $0x0000000000070001
        DATA vecDecodeTable+0x418(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x420(SB)/8, $0x0000000000070002
        DATA vecDecodeTable+0x428(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x430(SB)/8, $0x0000000700020001
        DATA vecDecodeTable+0x438(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x440(SB)/8, $0x0000000000070003
        DATA vecDecodeTable+0x448(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x450(SB)/8, $0x0000000700030001
        DATA vecDecodeTable+0x458(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x460(SB)/8, $0x0000000700030002
        DATA vecDecodeTable+0x468(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x470(SB)/8, $0x0007000300020001
        DATA vecDecodeTable+0x478(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x480(SB)/8, $0x0000000000070004
        DATA vecDecodeTable+0x488(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x490(SB)/8, $0x0000000700040001
        DATA vecDecodeTable+0x498(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x4a0(SB)/8, $0x0000000700040002
        DATA vecDecodeTable+0x4a8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x4b0(SB)/8, $0x0007000400020001
        DATA vecDecodeTable+0x4b8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x4c0(SB)/8, $0x0000000700040003
        DATA vecDecodeTable+0x4c8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x4d0(SB)/8, $0x0007000400030001
        DATA vecDecodeTable+0x4d8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x4e0(SB)/8, $0x0007000400030002
        DATA vecDecodeTable+0x4e8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x4f0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0x4f8(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x500(SB)/8, $0x0000000000070005
        DATA vecDecodeTable+0x508(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x510(SB)/8, $0x0000000700050001
        DATA vecDecodeTable+0x518(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x520(SB)/8, $0x0000000700050002
        DATA vecDecodeTable+0x528(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x530(SB)/8, $0x0007000500020001
        DATA vecDecodeTable+0x538(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x540(SB)/8, $0x0000000700050003
        DATA vecDecodeTable+0x548(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x550(SB)/8, $0x0007000500030001
        DATA vecDecodeTable+0x558(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x560(SB)/8, $0x0007000500030002
        DATA vecDecodeTable+0x568(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x570(SB)/8, $0x0005000300020001
        DATA vecDecodeTable+0x578(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x580(SB)/8, $0x0000000700050004
        DATA vecDecodeTable+0x588(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x590(SB)/8, $0x0007000500040001
        DATA vecDecodeTable+0x598(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x5a0(SB)/8, $0x0007000500040002
        DATA vecDecodeTable+0x5a8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x5b0(SB)/8, $0x0005000400020001
        DATA vecDecodeTable+0x5b8(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x5c0(SB)/8, $0x0007000500040003
        DATA vecDecodeTable+0x5c8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x5d0(SB)/8, $0x0005000400030001
        DATA vecDecodeTable+0x5d8(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x5e0(SB)/8, $0x0005000400030002
        DATA vecDecodeTable+0x5e8(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x5f0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0x5f8(SB)/8, $0x0000000000070005
        DATA vecDecodeTable+0x600(SB)/8, $0x0000000000070006
        DATA vecDecodeTable+0x608(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x610(SB)/8, $0x0000000700060001
        DATA vecDecodeTable+0x618(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x620(SB)/8, $0x0000000700060002
        DATA vecDecodeTable+0x628(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x630(SB)/8, $0x0007000600020001
        DATA vecDecodeTable+0x638(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x640(SB)/8, $0x0000000700060003
        DATA vecDecodeTable+0x648(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x650(SB)/8, $0x0007000600030001
        DATA vecDecodeTable+0x658(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x660(SB)/8, $0x0007000600030002
        DATA vecDecodeTable+0x668(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x670(SB)/8, $0x0006000300020001
        DATA vecDecodeTable+0x678(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x680(SB)/8, $0x0000000700060004
        DATA vecDecodeTable+0x688(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x690(SB)/8, $0x0007000600040001
        DATA vecDecodeTable+0x698(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x6a0(SB)/8, $0x0007000600040002
        DATA vecDecodeTable+0x6a8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x6b0(SB)/8, $0x0006000400020001
        DATA vecDecodeTable+0x6b8(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x6c0(SB)/8, $0x0007000600040003
        DATA vecDecodeTable+0x6c8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x6d0(SB)/8, $0x0006000400030001
        DATA vecDecodeTable+0x6d8(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x6e0(SB)/8, $0x0006000400030002
        DATA vecDecodeTable+0x6e8(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x6f0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0x6f8(SB)/8, $0x0000000000070006
        DATA vecDecodeTable+0x700(SB)/8, $0x0000000700060005
        DATA vecDecodeTable+0x708(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x710(SB)/8, $0x0007000600050001
        DATA vecDecodeTable+0x718(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x720(SB)/8, $0x0007000600050002
        DATA vecDecodeTable+0x728(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x730(SB)/8, $0x0006000500020001
        DATA vecDecodeTable+0x738(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x740(SB)/8, $0x0007000600050003
        DATA vecDecodeTable+0x748(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x750(SB)/8, $0x0006000500030001
        DATA vecDecodeTable+0x758(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x760(SB)/8, $0x0006000500030002
        DATA vecDecodeTable+0x768(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x770(SB)/8, $0x0005000300020001
        DATA vecDecodeTable+0x778(SB)/8, $0x0000000000070006
        DATA vecDecodeTable+0x780(SB)/8, $0x0007000600050004
        DATA vecDecodeTable+0x788(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x790(SB)/8, $0x0006000500040001
        DATA vecDecodeTable+0x798(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x7a0(SB)/8, $0x0006000500040002
        DATA vecDecodeTable+0x7a8(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x7b0(SB)/8, $0x0005000400020001
        DATA vecDecodeTable+0x7b8(SB)/8, $0x0000000000070006
        DATA vecDecodeTable+0x7c0(SB)/8, $0x0006000500040003
        DATA vecDecodeTable+0x7c8(SB)/8, $0x0000000000000007
        DATA vecDecodeTable+0x7d0(SB)/8, $0x0005000400030001
        DATA vecDecodeTable+0x7d8(SB)/8, $0x0000000000070006
        DATA vecDecodeTable+0x7e0(SB)/8, $0x0005000400030002
        DATA vecDecodeTable+0x7e8(SB)/8, $0x0000000000070006
        DATA vecDecodeTable+0x7f0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0x7f8(SB)/8, $0x0000000700060005
        DATA vecDecodeTable+0x800(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0x808(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x810(SB)/8, $0x0000000000080001
        DATA vecDecodeTable+0x818(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x820(SB)/8, $0x0000000000080002
        DATA vecDecodeTable+0x828(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x830(SB)/8, $0x0000000800020001
        DATA vecDecodeTable+0x838(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x840(SB)/8, $0x0000000000080003
        DATA vecDecodeTable+0x848(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x850(SB)/8, $0x0000000800030001
        DATA vecDecodeTable+0x858(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x860(SB)/8, $0x0000000800030002
        DATA vecDecodeTable+0x868(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x870(SB)/8, $0x0008000300020001
        DATA vecDecodeTable+0x878(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x880(SB)/8, $0x0000000000080004
        DATA vecDecodeTable+0x888(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x890(SB)/8, $0x0000000800040001
        DATA vecDecodeTable+0x898(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x8a0(SB)/8, $0x0000000800040002
        DATA vecDecodeTable+0x8a8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x8b0(SB)/8, $0x0008000400020001
        DATA vecDecodeTable+0x8b8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x8c0(SB)/8, $0x0000000800040003
        DATA vecDecodeTable+0x8c8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x8d0(SB)/8, $0x0008000400030001
        DATA vecDecodeTable+0x8d8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x8e0(SB)/8, $0x0008000400030002
        DATA vecDecodeTable+0x8e8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x8f0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0x8f8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0x900(SB)/8, $0x0000000000080005
        DATA vecDecodeTable+0x908(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x910(SB)/8, $0x0000000800050001
        DATA vecDecodeTable+0x918(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x920(SB)/8, $0x0000000800050002
        DATA vecDecodeTable+0x928(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x930(SB)/8, $0x0008000500020001
        DATA vecDecodeTable+0x938(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x940(SB)/8, $0x0000000800050003
        DATA vecDecodeTable+0x948(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x950(SB)/8, $0x0008000500030001
        DATA vecDecodeTable+0x958(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x960(SB)/8, $0x0008000500030002
        DATA vecDecodeTable+0x968(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x970(SB)/8, $0x0005000300020001
        DATA vecDecodeTable+0x978(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0x980(SB)/8, $0x0000000800050004
        DATA vecDecodeTable+0x988(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x990(SB)/8, $0x0008000500040001
        DATA vecDecodeTable+0x998(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x9a0(SB)/8, $0x0008000500040002
        DATA vecDecodeTable+0x9a8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x9b0(SB)/8, $0x0005000400020001
        DATA vecDecodeTable+0x9b8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0x9c0(SB)/8, $0x0008000500040003
        DATA vecDecodeTable+0x9c8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0x9d0(SB)/8, $0x0005000400030001
        DATA vecDecodeTable+0x9d8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0x9e0(SB)/8, $0x0005000400030002
        DATA vecDecodeTable+0x9e8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0x9f0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0x9f8(SB)/8, $0x0000000000080005
        DATA vecDecodeTable+0xa00(SB)/8, $0x0000000000080006
        DATA vecDecodeTable+0xa08(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xa10(SB)/8, $0x0000000800060001
        DATA vecDecodeTable+0xa18(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xa20(SB)/8, $0x0000000800060002
        DATA vecDecodeTable+0xa28(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xa30(SB)/8, $0x0008000600020001
        DATA vecDecodeTable+0xa38(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xa40(SB)/8, $0x0000000800060003
        DATA vecDecodeTable+0xa48(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xa50(SB)/8, $0x0008000600030001
        DATA vecDecodeTable+0xa58(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xa60(SB)/8, $0x0008000600030002
        DATA vecDecodeTable+0xa68(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xa70(SB)/8, $0x0006000300020001
        DATA vecDecodeTable+0xa78(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xa80(SB)/8, $0x0000000800060004
        DATA vecDecodeTable+0xa88(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xa90(SB)/8, $0x0008000600040001
        DATA vecDecodeTable+0xa98(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xaa0(SB)/8, $0x0008000600040002
        DATA vecDecodeTable+0xaa8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xab0(SB)/8, $0x0006000400020001
        DATA vecDecodeTable+0xab8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xac0(SB)/8, $0x0008000600040003
        DATA vecDecodeTable+0xac8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xad0(SB)/8, $0x0006000400030001
        DATA vecDecodeTable+0xad8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xae0(SB)/8, $0x0006000400030002
        DATA vecDecodeTable+0xae8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xaf0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0xaf8(SB)/8, $0x0000000000080006
        DATA vecDecodeTable+0xb00(SB)/8, $0x0000000800060005
        DATA vecDecodeTable+0xb08(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xb10(SB)/8, $0x0008000600050001
        DATA vecDecodeTable+0xb18(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xb20(SB)/8, $0x0008000600050002
        DATA vecDecodeTable+0xb28(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xb30(SB)/8, $0x0006000500020001
        DATA vecDecodeTable+0xb38(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xb40(SB)/8, $0x0008000600050003
        DATA vecDecodeTable+0xb48(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xb50(SB)/8, $0x0006000500030001
        DATA vecDecodeTable+0xb58(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xb60(SB)/8, $0x0006000500030002
        DATA vecDecodeTable+0xb68(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xb70(SB)/8, $0x0005000300020001
        DATA vecDecodeTable+0xb78(SB)/8, $0x0000000000080006
        DATA vecDecodeTable+0xb80(SB)/8, $0x0008000600050004
        DATA vecDecodeTable+0xb88(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xb90(SB)/8, $0x0006000500040001
        DATA vecDecodeTable+0xb98(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xba0(SB)/8, $0x0006000500040002
        DATA vecDecodeTable+0xba8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xbb0(SB)/8, $0x0005000400020001
        DATA vecDecodeTable+0xbb8(SB)/8, $0x0000000000080006
        DATA vecDecodeTable+0xbc0(SB)/8, $0x0006000500040003
        DATA vecDecodeTable+0xbc8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xbd0(SB)/8, $0x0005000400030001
        DATA vecDecodeTable+0xbd8(SB)/8, $0x0000000000080006
        DATA vecDecodeTable+0xbe0(SB)/8, $0x0005000400030002
        DATA vecDecodeTable+0xbe8(SB)/8, $0x0000000000080006
        DATA vecDecodeTable+0xbf0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0xbf8(SB)/8, $0x0000000800060005
        DATA vecDecodeTable+0xc00(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xc08(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xc10(SB)/8, $0x0000000800070001
        DATA vecDecodeTable+0xc18(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xc20(SB)/8, $0x0000000800070002
        DATA vecDecodeTable+0xc28(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xc30(SB)/8, $0x0008000700020001
        DATA vecDecodeTable+0xc38(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xc40(SB)/8, $0x0000000800070003
        DATA vecDecodeTable+0xc48(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xc50(SB)/8, $0x0008000700030001
        DATA vecDecodeTable+0xc58(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xc60(SB)/8, $0x0008000700030002
        DATA vecDecodeTable+0xc68(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xc70(SB)/8, $0x0007000300020001
        DATA vecDecodeTable+0xc78(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xc80(SB)/8, $0x0000000800070004
        DATA vecDecodeTable+0xc88(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xc90(SB)/8, $0x0008000700040001
        DATA vecDecodeTable+0xc98(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xca0(SB)/8, $0x0008000700040002
        DATA vecDecodeTable+0xca8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xcb0(SB)/8, $0x0007000400020001
        DATA vecDecodeTable+0xcb8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xcc0(SB)/8, $0x0008000700040003
        DATA vecDecodeTable+0xcc8(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xcd0(SB)/8, $0x0007000400030001
        DATA vecDecodeTable+0xcd8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xce0(SB)/8, $0x0007000400030002
        DATA vecDecodeTable+0xce8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xcf0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0xcf8(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xd00(SB)/8, $0x0000000800070005
        DATA vecDecodeTable+0xd08(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xd10(SB)/8, $0x0008000700050001
        DATA vecDecodeTable+0xd18(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xd20(SB)/8, $0x0008000700050002
        DATA vecDecodeTable+0xd28(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xd30(SB)/8, $0x0007000500020001
        DATA vecDecodeTable+0xd38(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xd40(SB)/8, $0x0008000700050003
        DATA vecDecodeTable+0xd48(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xd50(SB)/8, $0x0007000500030001
        DATA vecDecodeTable+0xd58(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xd60(SB)/8, $0x0007000500030002
        DATA vecDecodeTable+0xd68(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xd70(SB)/8, $0x0005000300020001
        DATA vecDecodeTable+0xd78(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xd80(SB)/8, $0x0008000700050004
        DATA vecDecodeTable+0xd88(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xd90(SB)/8, $0x0007000500040001
        DATA vecDecodeTable+0xd98(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xda0(SB)/8, $0x0007000500040002
        DATA vecDecodeTable+0xda8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xdb0(SB)/8, $0x0005000400020001
        DATA vecDecodeTable+0xdb8(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xdc0(SB)/8, $0x0007000500040003
        DATA vecDecodeTable+0xdc8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xdd0(SB)/8, $0x0005000400030001
        DATA vecDecodeTable+0xdd8(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xde0(SB)/8, $0x0005000400030002
        DATA vecDecodeTable+0xde8(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xdf0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0xdf8(SB)/8, $0x0000000800070005
        DATA vecDecodeTable+0xe00(SB)/8, $0x0000000800070006
        DATA vecDecodeTable+0xe08(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xe10(SB)/8, $0x0008000700060001
        DATA vecDecodeTable+0xe18(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xe20(SB)/8, $0x0008000700060002
        DATA vecDecodeTable+0xe28(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xe30(SB)/8, $0x0007000600020001
        DATA vecDecodeTable+0xe38(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xe40(SB)/8, $0x0008000700060003
        DATA vecDecodeTable+0xe48(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xe50(SB)/8, $0x0007000600030001
        DATA vecDecodeTable+0xe58(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xe60(SB)/8, $0x0007000600030002
        DATA vecDecodeTable+0xe68(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xe70(SB)/8, $0x0006000300020001
        DATA vecDecodeTable+0xe78(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xe80(SB)/8, $0x0008000700060004
        DATA vecDecodeTable+0xe88(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xe90(SB)/8, $0x0007000600040001
        DATA vecDecodeTable+0xe98(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xea0(SB)/8, $0x0007000600040002
        DATA vecDecodeTable+0xea8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xeb0(SB)/8, $0x0006000400020001
        DATA vecDecodeTable+0xeb8(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xec0(SB)/8, $0x0007000600040003
        DATA vecDecodeTable+0xec8(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xed0(SB)/8, $0x0006000400030001
        DATA vecDecodeTable+0xed8(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xee0(SB)/8, $0x0006000400030002
        DATA vecDecodeTable+0xee8(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xef0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0xef8(SB)/8, $0x0000000800070006
        DATA vecDecodeTable+0xf00(SB)/8, $0x0008000700060005
        DATA vecDecodeTable+0xf08(SB)/8, $0x0000000000000000
        DATA vecDecodeTable+0xf10(SB)/8, $0x0007000600050001
        DATA vecDecodeTable+0xf18(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xf20(SB)/8, $0x0007000600050002
        DATA vecDecodeTable+0xf28(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xf30(SB)/8, $0x0006000500020001
        DATA vecDecodeTable+0xf38(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xf40(SB)/8, $0x0007000600050003
        DATA vecDecodeTable+0xf48(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xf50(SB)/8, $0x0006000500030001
        DATA vecDecodeTable+0xf58(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xf60(SB)/8, $0x0006000500030002
        DATA vecDecodeTable+0xf68(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xf70(SB)/8, $0x0005000300020001
        DATA vecDecodeTable+0xf78(SB)/8, $0x0000000800070006
        DATA vecDecodeTable+0xf80(SB)/8, $0x0007000600050004
        DATA vecDecodeTable+0xf88(SB)/8, $0x0000000000000008
        DATA vecDecodeTable+0xf90(SB)/8, $0x0006000500040001
        DATA vecDecodeTable+0xf98(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xfa0(SB)/8, $0x0006000500040002
        DATA vecDecodeTable+0xfa8(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xfb0(SB)/8, $0x0005000400020001
        DATA vecDecodeTable+0xfb8(SB)/8, $0x0000000800070006
        DATA vecDecodeTable+0xfc0(SB)/8, $0x0006000500040003
        DATA vecDecodeTable+0xfc8(SB)/8, $0x0000000000080007
        DATA vecDecodeTable+0xfd0(SB)/8, $0x0005000400030001
        DATA vecDecodeTable+0xfd8(SB)/8, $0x0000000800070006
        DATA vecDecodeTable+0xfe0(SB)/8, $0x0005000400030002
        DATA vecDecodeTable+0xfe8(SB)/8, $0x0000000800070006
        DATA vecDecodeTable+0xff0(SB)/8, $0x0004000300020001
        DATA vecDecodeTable+0xff8(SB)/8, $0x0008000700060005
        GLOBL   vecDecodeTable(SB),  NOPTR|RODATA, $4096

        DATA lengthTable+0x000(SB)/8, $0x0302020102010100
        DATA lengthTable+0x008(SB)/8, $0x0403030203020201
        DATA lengthTable+0x010(SB)/8, $0x0403030203020201
        DATA lengthTable+0x018(SB)/8, $0x0504040304030302
        DATA lengthTable+0x020(SB)/8, $0x0403030203020201
        DATA lengthTable+0x028(SB)/8, $0x0504040304030302
        DATA lengthTable+0x030(SB)/8, $0x0504040304030302
        DATA lengthTable+0x038(SB)/8, $0x0605050405040403
        DATA lengthTable+0x040(SB)/8, $0x0403030203020201
        DATA lengthTable+0x048(SB)/8, $0x0504040304030302
        DATA lengthTable+0x050(SB)/8, $0x0504040304030302
        DATA lengthTable+0x058(SB)/8, $0x0605050405040403
        DATA lengthTable+0x060(SB)/8, $0x0504040304030302
        DATA lengthTable+0x068(SB)/8, $0x0605050405040403
        DATA lengthTable+0x070(SB)/8, $0x0605050405040403
        DATA lengthTable+0x078(SB)/8, $0x0706060506050504
        DATA lengthTable+0x080(SB)/8, $0x0403030203020201
        DATA lengthTable+0x088(SB)/8, $0x0504040304030302
        DATA lengthTable+0x090(SB)/8, $0x0504040304030302
        DATA lengthTable+0x098(SB)/8, $0x0605050405040403
        DATA lengthTable+0x0a0(SB)/8, $0x0504040304030302
        DATA lengthTable+0x0a8(SB)/8, $0x0605050405040403
        DATA lengthTable+0x0b0(SB)/8, $0x0605050405040403
        DATA lengthTable+0x0b8(SB)/8, $0x0706060506050504
        DATA lengthTable+0x0c0(SB)/8, $0x0504040304030302
        DATA lengthTable+0x0c8(SB)/8, $0x0605050405040403
        DATA lengthTable+0x0d0(SB)/8, $0x0605050405040403
        DATA lengthTable+0x0d8(SB)/8, $0x0706060506050504
        DATA lengthTable+0x0e0(SB)/8, $0x0605050405040403
        DATA lengthTable+0x0e8(SB)/8, $0x0706060506050504
        DATA lengthTable+0x0f0(SB)/8, $0x0706060506050504
        DATA lengthTable+0x0f8(SB)/8, $0x0807070607060605
        GLOBL   lengthTable(SB),  NOPTR|RODATA, $256
