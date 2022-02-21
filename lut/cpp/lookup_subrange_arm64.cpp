#include <arm_neon.h>

void
lookup_subrange_neon(uint8x16_t *input, uint8x16_t *output, char *range_upper, uint8x16_t *lut, int ranges,
                     int size) {
    int chunks = size / 16;

    for (int i = 0; i < chunks; i++) {
        uint8x16_t index = vdupq_n_u8(0);
        for (int j = 0; j < ranges; j++) {
            index = vsubq_u8(index, vcgtq_u8(input[i], vdupq_n_u8(range_upper[j])));
        }
        output[i] = vqtbl1q_u8(*lut, index);
    }

    // Leftovers
    for (int i = chunks * 16; i < size; i++) {
        uint8_t n = 0;
        for (int j = 0; j < ranges; j++) {
            if (((uint8_t *) input)[i] > range_upper[j]) {
                n += 1;
            }
        }
        ((uint8_t *) output)[i] = ((uint8_t *) lut)[n];
    }
}
