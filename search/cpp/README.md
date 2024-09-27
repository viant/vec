

Generating amd64 assembler files

1. Upgrade clang to version 18.0.0 or later to produce automatically vectorized code.

2. Generate the assembler file

```bash
 cd cpp
 clang -S -mavx2 -mfma -masm=intel -fno-asynchronous-unwind-tables -mstackrealign -Ofast   file_name_1.cpp
 
 cd ..
     c2goasm -a -f cpp/file_name_1.s file_name1_2.s
```

3. Edit the resulting files to remove lines like below (since c2 cannot handle them as it did in earlier versions of clang):
   // push	rbp
   // mov	rbp, rsp
   // and	rsp, -8

