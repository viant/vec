#include <stdint.h>


inline uint32_t bits(const uint32_t v) {
    return v == 0 ? 0 : 32 - __builtin_clz(v);
}

uint32_t maxbits(const uint32_t *in, const uint32_t size) asm("maxbits");

uint32_t maxbits(const uint32_t *in, const uint32_t size) {
    uint32_t acc = 0;
    for (int i = 0; i < size; i++) {
        acc |= in[i];
    }
    return bits(acc);
}

