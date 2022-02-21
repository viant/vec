#include <stdint.h>

void populate_positions_ctz(uint64_t input, uint32_t *out, int *count) {
    int pos = 0;
    while (input != 0) {
        out[pos++] = __builtin_ctzll(input);
        input ^= input & -input;
    }
    *count = pos;
}
