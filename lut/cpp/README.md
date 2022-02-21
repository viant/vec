# Building ARM assembly code

### The C function implementing neon vectorized table lookup:
```c
#include <arm_neon.h>

void lookup_neon(const uint8x16_t *input, uint8x16_t *output, const uint8x16x4_t *table, const int chunks) {
    for (int i = 0; i < chunks; i++) {
        const uint8x16_t t1 = vqtbl4q_u8(table[0], input[i]);
        const uint8x16_t t2 = vqtbl4q_u8(table[1], veorq_u8(input[i], vdupq_n_u8(0x40)));
        const uint8x16_t t3 = vqtbl4q_u8(table[2], veorq_u8(input[i], vdupq_n_u8(0x80)));
        const uint8x16_t t4 = vqtbl4q_u8(table[3], veorq_u8(input[i], vdupq_n_u8(0xc0)));
        output[i] = vorrq_u8(vorrq_u8(t1, t2), vorrq_u8(t3, t4));
    }
}
```
### Step 1: Compile with clang on Mac OSX:

clang -c -arch arm64 -march=armv8  lookup_neon.cpp -Ofast

### Step 2: Disassemble the output object file using the Go disasembler:

go tool objdump  lookup_neon.o
```asm
  :0			0x0			    7100047f		CMPW $1, R3							
  :0			0x4			    5400036b		BLT 27(PC)							
  :0			0x8			    2a0303e8		MOVW R3, R8							
  :0			0xc			    4f02e400		VMOVI $64, V0.B16						
  :0			0x10			4f04e401		VMOVI $128, V1.B16						
  :0			0x14			4f06e402		VMOVI $192, V2.B16						
  :0			0x18			3cc10403		FMOVQ.P 16(R0), F3						
  :0			0x1c			ad401444		FLDPQ (R2), (F4, F5)						
  :0			0x20			ad411c46		FLDPQ 32(R2), (F6, F7)						
  :0			0x24			6e201c70		VEOR V0.B16, V3.B16, V16.B16					
  :0			0x28			ad424851		FLDPQ 64(R2), (F17, F18)					
  :0			0x2c			ad435053		FLDPQ 96(R2), (F19, F20)					
  :0			0x30			ad445855		FLDPQ 128(R2), (F21, F22)					
  :0			0x34			4e106230		VTBL V16.B16, [V17.B16, V18.B16, V19.B16, V20.B16], V16.B16	
  :0			0x38			6e211c71		VEOR V1.B16, V3.B16, V17.B16					
  :0			0x3c			ad456057		FLDPQ 160(R2), (F23, F24)					
  :0			0x40			ad466859		FLDPQ 192(R2), (F25, F26)					
  :0			0x44			4e1162b1		VTBL V17.B16, [V21.B16, V22.B16, V23.B16, V24.B16], V17.B16	
  :0			0x48			4e036084		VTBL V3.B16, [V4.B16, V5.B16, V6.B16, V7.B16], V4.B16		
  :0			0x4c			6e221c63		VEOR V2.B16, V3.B16, V3.B16					
  :0			0x50			ad47705b		FLDPQ 224(R2), (F27, F28)					
  :0			0x54			4e036323		VTBL V3.B16, [V25.B16, V26.B16, V27.B16, V28.B16], V3.B16	
  :0			0x58			4ea41e04		VORR V4.B16, V16.B16, V4.B16					
  :0			0x5c			4eb11c84		VORR V17.B16, V4.B16, V4.B16					
  :0			0x60			4ea31c83		VORR V3.B16, V4.B16, V3.B16					
  :0			0x64			3c810423		FMOVQ.P F3, 16(R1)						
  :0			0x68			f1000508		SUBS $1, R8, R8							
  :0			0x6c			54fffd61		BNE -21(PC)							
  :0			0x70			d65f03c0		RET								

```

### Step 3: Put together a Go assembler compliant file

We need to combine an argument passing prefix with 
hexadecimal instruction constants from Step 2:

```asm
TEXT ·_transform16_neon(SB), $0-32                                                          
                                   
  MOVD input+0(FP),  R0                         
  MOVD output+8(FP), R1                         
  MOVD table+16(FP), R2                         
  MOVD chunks+24(FP),R3                         

  WORD $0x7100047f
  WORD $0x5400036b
  WORD $0x2a0303e8
  WORD $0x4f02e400
  WORD $0x4f04e401
  WORD $0x4f06e402
  WORD $0xad401043
  WORD $0xad411845
  WORD $0x3cc10407
  WORD $0xad424450
  WORD $0xad434c52
  WORD $0x6e201cf4
  WORD $0xad445855
  WORD $0x4e146210
  WORD $0xad456057
  WORD $0x6e211cf1
  WORD $0xad466859
  WORD $0x4e1162b1
  WORD $0x4e076063
  WORD $0xad47705b
  WORD $0x6e221ce4
  WORD $0x4e046324
  WORD $0x4ea31e03
  WORD $0x4eb11c63
  WORD $0x4ea41c63
  WORD $0x3c810423
  WORD $0xf1000508
  WORD $0x54fffd61
  WORD $0xd65f03c0
```

### Note:                                                                                        
In Go version 1.17 and above it is possible to use the mneumonic decompiler output instead of hex
constants like this:

```asm
TEXT ·_transform16_neon(SB), $0-32                                                          
                                   
  MOVD input+0(FP),  R0                         
  MOVD output+8(FP), R1                         
  MOVD table+16(FP), R2                         
  MOVD chunks+24(FP),R3                         

  CMPW $1, R3                                                
  BLT 27(PC)                                                 
  MOVW R3, R8                                                
  VMOVI $64, V0.B16                                          
  VMOVI $128, V1.B16                                         
  VMOVI $192, V2.B16                                         
  FLDPQ (R2), (F3, F4)                                       
  FLDPQ 32(R2), (F5, F6)                                     
  FMOVQ.P 16(R0), F7                                         
  FLDPQ 64(R2), (F16, F17)                                   
  FLDPQ 96(R2), (F18, F19)                                   
  VEOR V0.B16, V7.B16, V20.B16                               
  FLDPQ 128(R2), (F21, F22)                                  
  VTBL V20.B16, [V16.B16, V17.B16, V18.B16, V19.B16], V16.B16
  FLDPQ 160(R2), (F23, F24)                                  
  VEOR V1.B16, V7.B16, V17.B16                               
  FLDPQ 192(R2), (F25, F26)                                  
  VTBL V17.B16, [V21.B16, V22.B16, V23.B16, V24.B16], V17.B16
  VTBL V7.B16, [V3.B16, V4.B16, V5.B16, V6.B16], V3.B16      
  FLDPQ 224(R2), (F27, F28)                                  
  VEOR V2.B16, V7.B16, V4.B16                                
  VTBL V4.B16, [V25.B16, V26.B16, V27.B16, V28.B16], V4.B16  
  VORR V3.B16, V16.B16, V3.B16                               
  VORR V17.B16, V3.B16, V3.B16                               
  VORR V4.B16, V3.B16, V3.B16                                
  FMOVQ.P F3, 16(R1)                                         
  SUBS $1, R8, R8                                            
  BNE -21(PC)                                                
  RET                                                        
```

However, the 1.17 Go assembler may incorrectly compile decompiled instructions. E.g.
'MOVW R3, R8'  will be replaced with 'SXTW R3, R8'.  Using hex WORD constants ensures that the original machine 
code will be unchanged.

# Building Intel assembly code

There a tool for the Intel architecture that makes the above manual process easier.

### Step 1: Compile the original C code to assembler 

clang  -mavx512vbmi -S -masm=intel -fno-asynchronous-unwind-tables lookup_512vbmi.cpp -Ofast

### Step 2: Transform the assembler code into the Go Assembler form:

c2goasm -f -a cpp/lookup_512vbmi.s lookup_amd64.s

### Note:

For more information on c2goasm see https://github.com/minio/c2goasm

