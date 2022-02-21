#include <arm_neon.h>

void lookup_neon(const uint8x16_t *input, uint8x16_t *output, const uint8x16x4_t *table, const int size) {
    int chunks = size / 16;

    for (int i = 0; i < chunks; i++) {
        const uint8x16_t t1 = vqtbl4q_u8(table[0], input[i]);
        const uint8x16_t t2 = vqtbl4q_u8(table[1], veorq_u8(input[i], vdupq_n_u8(0x40)));
        const uint8x16_t t3 = vqtbl4q_u8(table[2], veorq_u8(input[i], vdupq_n_u8(0x80)));
        const uint8x16_t t4 = vqtbl4q_u8(table[3], veorq_u8(input[i], vdupq_n_u8(0xc0)));
        output[i] = vorrq_u8(vorrq_u8(t1, t2), vorrq_u8(t3, t4));
    }

    // leftovers
    for (int i = chunks * 16; i < size; i++) {
        ((uint8_t *) output)[i] = ((uint8_t *) table)[((uint8_t *) input)[i]];
    }
}
