#include "textflag.h"                                                    
                                                                 
TEXT ·popcnt(SB), NOSPLIT, $0                                            
  MOVD src+0(FP), R0                                                     
  MOVD len+8(FP), R1                                                     
  PRFM (R0), PLDL1KEEP
                                                                       
  WORD $0xaa0003e8   // MOVD R0, R8                                      
  WORD $0xf102003f   // CMP $128, R1                                     
  WORD $0x54000463   // BCC 35(PC)                                       
  WORD $0xd2800000   // MOVD $0, R0                                      
  WORD $0xaa0103e9   // MOVD R1, R9                                      
  WORD $0xaa0803ea   // MOVD R8, R10                                     
  WORD $0xaa0a03eb   // MOVD R10, R11                                    
  WORD $0x4cdf2160   // VLD1.P 64(R11), [V0.B16, V1.B16, V2.B16, V3.B16] 
  WORD $0x4c402164   // VLD1 (R11), [V4.B16, V5.B16, V6.B16, V7.B16]     
  WORD $0x4e205810   // VCNT V0.B16, V16.B16                             
  WORD $0x4e205831   // VCNT V1.B16, V17.B16                             
  WORD $0x4e205852   // VCNT V2.B16, V18.B16                             
  WORD $0x4e205860   // VCNT V3.B16, V0.B16                              
  WORD $0x4e205881   // VCNT V4.B16, V1.B16                              
  WORD $0x4e2058a2   // VCNT V5.B16, V2.B16                              
  WORD $0x4e2058c3   // VCNT V6.B16, V3.B16                              
  WORD $0x4e2058e4   // VCNT V7.B16, V4.B16                              
  WORD $0x4e328400   // VADD V18.B16, V0.B16, V0.B16                     
  WORD $0x4e308400   // VADD V16.B16, V0.B16, V0.B16                     
  WORD $0x4e318400   // VADD V17.B16, V0.B16, V0.B16                     
  WORD $0x4e238400   // VADD V3.B16, V0.B16, V0.B16                      
  WORD $0x4e248400   // VADD V4.B16, V0.B16, V0.B16                      
  WORD $0x4e218400   // VADD V1.B16, V0.B16, V0.B16                      
  WORD $0x4e228400   // VADD V2.B16, V0.B16, V0.B16                      
  WORD $0x6e303800   // VUADDLV V0.B16, V0                               
  WORD $0x1e26000b   // FMOVS F0, R11                                    
  WORD $0x12003d6b   // ANDW $65535, R11, R11                            
  WORD $0x8b0b0000   // ADD R11, R0, R0                                  
  WORD $0x9102014a   // ADD $128, R10, R10                               
  WORD $0xd1020129   // SUB $128, R9, R9                                 
  WORD $0xf101fd3f   // CMP $127, R9                                     
  WORD $0x54fffce8   // BHI -25(PC)                                      
  WORD $0x9279e029   // AND $-128, R1, R9                                
  WORD $0xcb09002a   // SUB R9, R1, R10                                  
  WORD $0xf100215f   // CMP $8, R10                                      
  WORD $0x540000e2   // BCS 7(PC)                                        
  WORD $0x1400000f   // JMP 15(PC)                                       
  WORD $0xd2800009   // MOVD $0, R9                                      
  WORD $0xd2800000   // MOVD $0, R0                                      
  WORD $0xcb09002a   // SUB R9, R1, R10                                  
  WORD $0xf100215f   // CMP $8, R10                                      
  WORD $0x54000143   // BCC 10(PC)                                       
  WORD $0xfc696900   // FMOVD (R8)(R9), F0                               
  WORD $0x0e205800   // VCNT V0.B8, V0.B8                                
  WORD $0x0e31b800   // VADDV V0.B8, V0                                  
  WORD $0x1e26000b   // FMOVS F0, R11                                    
  WORD $0x8b0b0000   // ADD R11, R0, R0                                  
  WORD $0x91002129   // ADD $8, R9, R9                                   
  WORD $0xd100214a   // SUB $8, R10, R10                                 
  WORD $0xf1001d5f   // CMP $7, R10                                      
  WORD $0x54ffff08   // BHI -8(PC)                                       
  WORD $0xb400034a   // CBZ R10, 26(PC)                                  
  WORD $0xf100115f   // CMP $4, R10                                      
  WORD $0x54000123   // BCC 9(PC)                                        
  WORD $0x8b09010a   // ADD R9, R8, R10                                  
  WORD $0x2f00e400   // VMOVI $0, V0                                     
  WORD $0x0d408140   // VLD1 (R10), V0.S[0]                              
  WORD $0xb27e0129   // ORR $4, R9, R9                                   
  WORD $0xcb09002a   // SUB R9, R1, R10                                  
  WORD $0xf100095f   // CMP $2, R10                                      
  WORD $0x540000c2   // BCS 6(PC)                                        
  WORD $0x14000008   // JMP 8(PC)                                        
  WORD $0x4f00e400   // VMOVI $0, V0.B16                                 
  WORD $0xcb09002a   // SUB R9, R1, R10                                  
  WORD $0xf100095f   // CMP $2, R10                                      
  WORD $0x54000083   // BCC 4(PC)                                        
  WORD $0x8b09010a   // ADD R9, R8, R10                                  
  WORD $0x0d405140   // VLD1 (R10), V0.H[2]                              
  WORD $0x91000929   // ADD $2, R9, R9                                   
  WORD $0xeb01013f   // CMP R1, R9                                       
  WORD $0x54000060   // BEQ 3(PC)                                        
  WORD $0x8b090108   // ADD R9, R8, R8                                   
  WORD $0x0d401900   // VLD1 (R8), V0.B[6]                               
  WORD $0x0e205800   // VCNT V0.B8, V0.B8                                
  WORD $0x0e31b800   // VADDV V0.B8, V0                                  
  WORD $0x1e260008   // FMOVS F0, R8                                     
  WORD $0x8b080000   // ADD R8, R0, R0
  MOVD R0, ret+16(FP)                                   
  WORD $0xd65f03c0   // RET                                              
                                                                             
                                                                             
                                                                             
                                                                             
                                                                             
                                                                             
                                                                             
                                                                             
                                                                             
                                                                             
                                                                             
                                                                             
