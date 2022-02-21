#include <arm_neon.h>

void expand_mask_neon2(uint64_t *input, uint64x2_t *output) {
    uint8x16_t bit_mask = vreinterpretq_u8_u64(vdupq_n_u64(0x7fbfdfeff7fbfdfe));
    uint64x2_t table = {0, 0};
    table[0] = *input;
    uint64x2_t index[4] = {
            {0,                  0x0101010101010101},
            {0x0202020202020202, 0x0303030303030303},
            {0x0404040404040404, 0x0505050505050505},
            {0x0606060606060606, 0x0707070707070707}
    };
    for (int i = 0; i < 4; i++) {
        uint8x16_t vmask = vqtbl1q_u8(vreinterpretq_u8_u64(table), vreinterpretq_u8_u64(index[i]));
        vmask = vorrq_u8(vmask, bit_mask);
        output[i] = vreinterpretq_u64_u8(vceqq_u8(vmask, vdupq_n_u8(0xff)));
    }
}
