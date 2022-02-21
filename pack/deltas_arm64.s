TEXT ·_compute_deltas(SB), $0-32
	MOVD input+0(FP),           R0
	MOVD length+8(FP),          R1
	MOVD output+16(FP),         R2
	MOVD starting_point+24(FP), R3

  WORD $0x53027c28	// lsr w8, w1, #2				
  WORD $0x4e040c60	// dup v0.4s, w3				
  WORD $0x34000188	// cbz w8, .+0x30				
  WORD $0x2a0803e8	// mov w8, w8				
  WORD $0xaa0203e9	// mov x9, x2				
  WORD $0xaa0003ea	// mov x10, x0				
  WORD $0x4ea01c01	// mov v1.16b, v0.16b			
  WORD $0x3cc10540	// ldr q0, [x10],#16			
  WORD $0xf1000508	// subs x8, x8, #0x1			
  WORD $0x6e006021	// ext v1.16b, v1.16b, v0.16b, #12		
  WORD $0x6ea18401	// sub v1.4s, v0.4s, v1.4s			
  WORD $0x3c810521	// str q1, [x9],#16			
  WORD $0x54ffff41	// b.ne .+0xffffffffffffffe8		
  WORD $0x121e7428	// and w8, w1, #0xfffffffc			
  WORD $0x6b01011f	// cmp w8, w1				
  WORD $0x540005c2	// b.cs .+0xb8				
  WORD $0x2a0803e8	// mov w8, w8				
  WORD $0x2a0103e9	// mov w9, w1				
  WORD $0xcb08012a	// sub x10, x9, x8				
  WORD $0xf100215f	// cmp x10, #0x8				
  WORD $0x0e1c3c0b	// mov w11, v0.s[3]			
  WORD $0x540003c3	// b.cc .+0x78				
  WORD $0xd37ef50c	// lsl x12, x8, #2				
  WORD $0xd37ef52d	// lsl x13, x9, #2				
  WORD $0x8b0c004e	// add x14, x2, x12			
  WORD $0x8b0d000f	// add x15, x0, x13			
  WORD $0xeb0f01df	// cmp x14, x15				
  WORD $0x540000a2	// b.cs .+0x14				
  WORD $0x8b0d004d	// add x13, x2, x13			
  WORD $0x8b0c000e	// add x14, x0, x12			
  WORD $0xeb0d01df	// cmp x14, x13				
  WORD $0x54000283	// b.cc .+0x50				
  WORD $0x927df14b	// and x11, x10, #0xfffffffffffffff8	
  WORD $0x9100418d	// add x13, x12, #0x10			
  WORD $0x8b080168	// add x8, x11, x8				
  WORD $0x8b0d004c	// add x12, x2, x13			
  WORD $0x8b0d000d	// add x13, x0, x13			
  WORD $0xaa0b03ee	// mov x14, x11				
  WORD $0x3cdf01a1	// ldur q1, [x13,#-16]			
  WORD $0xf10021ce	// subs x14, x14, #0x8			
  WORD $0x6e016002	// ext v2.16b, v0.16b, v1.16b, #12		
  WORD $0x3cc205a0	// ldr q0, [x13],#32			
  WORD $0x6ea28422	// sub v2.4s, v1.4s, v2.4s			
  WORD $0x6e006021	// ext v1.16b, v1.16b, v0.16b, #12		
  WORD $0x6ea18401	// sub v1.4s, v0.4s, v1.4s			
  WORD $0xad3f8582	// stp q2, q1, [x12,#-16]			
  WORD $0x9100818c	// add x12, x12, #0x20			
  WORD $0x54fffee1	// b.ne .+0xffffffffffffffdc		
  WORD $0xeb0b015f	// cmp x10, x11				
  WORD $0x54000180	// b.eq .+0x30				
  WORD $0x0e1c3c0b	// mov w11, v0.s[3]			
  WORD $0xd37ef50c	// lsl x12, x8, #2				
  WORD $0x8b0c000a	// add x10, x0, x12			
  WORD $0x8b0c004c	// add x12, x2, x12			
  WORD $0xcb080128	// sub x8, x9, x8				
  WORD $0xb8404549	// ldr w9, [x10],#4			
  WORD $0xf1000508	// subs x8, x8, #0x1			
  WORD $0x4b0b012b	// sub w11, w9, w11			
  WORD $0xb800458b	// str w11, [x12],#4			
  WORD $0x2a0903eb	// mov w11, w9				
  WORD $0x54ffff61	// b.ne .+0xffffffffffffffec		
  WORD $0xd65f03c0	// ret					

TEXT ·_compute_prefix_sum(SB), $0-32                       
	MOVD input+0(FP),           R0                           
	MOVD length+8(FP),          R1                           
	MOVD output+16(FP),         R2                           
	MOVD starting_point+24(FP), R3                           

  WORD $0x53027c28  // lsr w8, w1, #2			
  WORD $0x34000248  // cbz w8, .+0x48			
  WORD $0x2a0803e8  // mov w8, w8			
  WORD $0x6f00e400  // movi v0.2d, #0x0		
  WORD $0xaa0203e9  // mov x9, x2			
  WORD $0xaa0003ea  // mov x10, x0			
  WORD $0x3cc10541  // ldr q1, [x10],#16		
  WORD $0x4e040c62  // dup v2.4s, w3			
  WORD $0xf1000508  // subs x8, x8, #0x1		
  WORD $0x6e016003  // ext v3.16b, v0.16b, v1.16b, #12	
  WORD $0x4ea18461  // add v1.4s, v3.4s, v1.4s		
  WORD $0x6e014003  // ext v3.16b, v0.16b, v1.16b, #8	
  WORD $0x4ea18461  // add v1.4s, v3.4s, v1.4s		
  WORD $0x4ea28422  // add v2.4s, v1.4s, v2.4s		
  WORD $0x0e1c3c2b  // mov w11, v1.s[3]		
  WORD $0x3c810522  // str q2, [x9],#16		
  WORD $0x0b030163  // add w3, w11, w3			
  WORD $0x54fffea1  // b.ne .+0xffffffffffffffd4	
  WORD $0x121e7428  // and w8, w1, #0xfffffffc		
  WORD $0x6b01011f  // cmp w8, w1			
  WORD $0x54000182  // b.cs .+0x30			
  WORD $0x2a0803ea  // mov w10, w8			
  WORD $0x2a0103eb  // mov w11, w1			
  WORD $0xd37ef549  // lsl x9, x10, #2			
  WORD $0x8b090008  // add x8, x0, x9			
  WORD $0x8b090049  // add x9, x2, x9			
  WORD $0xcb0a016a  // sub x10, x11, x10		
  WORD $0xb840450b  // ldr w11, [x8],#4		
  WORD $0xf100054a  // subs x10, x10, #0x1		
  WORD $0x0b030163  // add w3, w11, w3			
  WORD $0xb8004523  // str w3, [x9],#4			
  WORD $0x54ffff81  // b.ne .+0xfffffffffffffff0	
  WORD $0xd65f03c0  // ret				
