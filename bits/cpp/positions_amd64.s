	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 11, 0
	.intel_syntax noprefix
	.section	__TEXT,__literal4,4byte_literals
	.p2align	2                               ## -- Begin function populate_positions_avx2
LCPI0_0:
	.long	7                               ## 0x7
LCPI0_1:
	.long	15                              ## 0xf
LCPI0_2:
	.long	23                              ## 0x17
LCPI0_3:
	.long	31                              ## 0x1f
LCPI0_4:
	.long	39                              ## 0x27
LCPI0_5:
	.long	47                              ## 0x2f
LCPI0_6:
	.long	55                              ## 0x37
	.section	__TEXT,__const
	.p2align	5                               ## @_ZL14vecDecodeTable
LCPI0_7:
	.space	32
	.long	1                               ## 0x1
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	0                               ## 0x0
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.long	0                               ## 0x0
	.long	1                               ## 0x1
	.long	2                               ## 0x2
	.long	3                               ## 0x3
	.long	4                               ## 0x4
	.long	5                               ## 0x5
	.long	6                               ## 0x6
	.long	7                               ## 0x7
	.long	8                               ## 0x8
	.p2align	4                               ## @_ZL11lengthTable
LCPI0_8:
	.byte	000
	.byte	001
	.byte	001
	.byte	002
	.byte	001
	.byte	002
	.byte	002
	.byte	003
	.byte	001
	.byte	002
	.byte	002
	.byte	003
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	001
	.byte	002
	.byte	002
	.byte	003
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	001
	.byte	002
	.byte	002
	.byte	003
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	001
	.byte	002
	.byte	002
	.byte	003
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	005
	.byte	006
	.byte	006
	.byte	007
	.byte	001
	.byte	002
	.byte	002
	.byte	003
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	005
	.byte	006
	.byte	006
	.byte	007
	.byte	002
	.byte	003
	.byte	003
	.byte	004
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	005
	.byte	006
	.byte	006
	.byte	007
	.byte	003
	.byte	004
	.byte	004
	.byte	005
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	005
	.byte	006
	.byte	006
	.byte	007
	.byte	004
	.byte	005
	.byte	005
	.byte	006
	.byte	005
	.byte	006
	.byte	006
	.byte	007
	.byte	005
	.byte	006
	.byte	006
	.byte	007
	.byte	006
	.byte	007
	.byte	007
	.byte	008
	.section	__TEXT,__text,regular,pure_instructions
	.globl	populate_positions_avx2
	.p2align	4, 0x90
populate_positions_avx2:                ## @"\01populate_positions_avx2"
## %bb.0:
	push	rbp
	mov	rbp, rsp
	push	r15
	push	r14
	push	r12
	push	rbx
	and	rsp, -8
	mov	rcx, rsi
	test	rdi, rdi
	je	LBB0_2
## %bb.1:
	mov	rbx, rdi
	movzx	edi, bh
	mov	ecx, ebx
	shr	ecx, 16
	movzx	r10d, cl
	movzx	ecx, bl
	lea	r8, [rip + LCPI0_7]
	lea	r9, [rip + LCPI0_8]
	movzx	r11d, byte ptr [rcx + r9]
	shl	rcx, 5
	movzx	r14d, byte ptr [rdi + r9]
	shl	rdi, 5
	vpcmpeqd	ymm0, ymm0, ymm0
	vpaddd	ymm0, ymm0, ymmword ptr [rcx + r8]
	vpbroadcastd	ymm1, dword ptr [rip + LCPI0_0] ## ymm1 = [7,7,7,7,7,7,7,7]
	vpaddd	ymm1, ymm1, ymmword ptr [rdi + r8]
	vmovdqu	ymmword ptr [rsi], ymm0
	lea	r15, [rsi + 4*r11]
	vmovdqu	ymmword ptr [rsi + 4*r11], ymm1
	lea	r12, [r15 + 4*r14]
	mov	eax, ebx
	shr	eax, 24
	mov	rcx, rbx
	shr	rcx, 32
	movzx	edi, byte ptr [r10 + r9]
	shl	r10, 5
	movzx	r11d, byte ptr [rax + r9]
	shl	rax, 5
	vpbroadcastd	ymm0, dword ptr [rip + LCPI0_1] ## ymm0 = [15,15,15,15,15,15,15,15]
	vpaddd	ymm0, ymm0, ymmword ptr [r10 + r8]
	vpbroadcastd	ymm1, dword ptr [rip + LCPI0_2] ## ymm1 = [23,23,23,23,23,23,23,23]
	vpaddd	ymm1, ymm1, ymmword ptr [rax + r8]
	vmovdqu	ymmword ptr [r15 + 4*r14], ymm0
	lea	r10, [r12 + 4*rdi]
	vmovdqu	ymmword ptr [r12 + 4*rdi], ymm1
	lea	r14, [r10 + 4*r11]
	mov	rax, rbx
	shr	rax, 40
	mov	rdi, rbx
	shr	rdi, 48
	movzx	ecx, cl
	movzx	eax, al
	movzx	r15d, byte ptr [rcx + r9]
	shl	rcx, 5
	movzx	r12d, byte ptr [rax + r9]
	shl	rax, 5
	vpbroadcastd	ymm0, dword ptr [rip + LCPI0_3] ## ymm0 = [31,31,31,31,31,31,31,31]
	vpaddd	ymm0, ymm0, ymmword ptr [rcx + r8]
	vpbroadcastd	ymm1, dword ptr [rip + LCPI0_4] ## ymm1 = [39,39,39,39,39,39,39,39]
	vpaddd	ymm1, ymm1, ymmword ptr [rax + r8]
	vmovdqu	ymmword ptr [r10 + 4*r11], ymm0
	lea	r10, [r14 + 4*r15]
	vmovdqu	ymmword ptr [r14 + 4*r15], ymm1
	lea	rcx, [r10 + 4*r12]
	shr	rbx, 56
	movzx	edi, dil
	movzx	eax, byte ptr [rdi + r9]
	shl	rdi, 5
	movzx	r9d, byte ptr [rbx + r9]
	shl	rbx, 5
	vpbroadcastd	ymm0, dword ptr [rip + LCPI0_5] ## ymm0 = [47,47,47,47,47,47,47,47]
	vpaddd	ymm0, ymm0, ymmword ptr [rdi + r8]
	vpbroadcastd	ymm1, dword ptr [rip + LCPI0_6] ## ymm1 = [55,55,55,55,55,55,55,55]
	vpaddd	ymm1, ymm1, ymmword ptr [rbx + r8]
	vmovdqu	ymmword ptr [r10 + 4*r12], ymm0
	lea	rdi, [rcx + 4*rax]
	vmovdqu	ymmword ptr [rcx + 4*rax], ymm1
	lea	rcx, [rdi + 4*r9]
LBB0_2:
	sub	rcx, rsi
	shr	rcx, 2
	mov	dword ptr [rdx], ecx
	lea	rsp, [rbp - 32]
	pop	rbx
	pop	r12
	pop	r14
	pop	r15
	pop	rbp
	vzeroupper
    	ret
                                            ## -- End function
.subsections_via_symbols
