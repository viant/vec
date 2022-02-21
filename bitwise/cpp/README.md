

Generating bitwise_amd64.s for bitwise_amd64.go

```bash
 cd cpp
 clang -S -mavx512vbmi -masm=intel -fno-asynchronous-unwind-tables -mstackrealign -Ofast   bitwise_amd64.cpp
 
 cd ..
     c2goasm -a -f cpp/bitwise_amd64.s bitwise_amd64.s
```

Generating bitwise_arm64.s for bitwise_arm64.go

```bash
 cd cpp
 clang -c -arch arm64 -march=armv8+sve bitwise_arm64.cpp -Ofast
 go tool objdump  bitwise_arm64.o
```