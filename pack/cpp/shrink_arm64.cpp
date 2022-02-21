#include <arm_neon.h>

void packed_mask_neon(uint8x16_t *input, uint64_t *output, uint64_t size) {
    size = (size > 8) ? 8 : size;
    uint64_t chunks = size / 2;
    uint64_t remainder = size % 2;

    for (uint64_t i = 0; i < chunks; i++) {
        uint16x8_t high_bits = vreinterpretq_u16_u8(vshrq_n_u8(input[i], 7));
        uint32x4_t paired16 = vreinterpretq_u32_u16(vsraq_n_u16(high_bits, high_bits, 7));
        uint64x2_t paired32 = vreinterpretq_u64_u32(vsraq_n_u32(paired16, paired16, 14));
        uint8x16_t paired64 = vreinterpretq_u8_u64(vsraq_n_u64(paired32, paired32, 28));
        *output |= (uint64_t) vgetq_lane_u8(paired64, 0) << 8 * 2 * i |
                   (uint64_t) vgetq_lane_u8(paired64, 8) << 8 * (2 * i + 1);
    }

    if (remainder > 0) {
        uint64x2_t padding = {0, 0};
        padding[0] = ((uint64_t *) input)[2 * chunks];
        uint16x8_t high_bits = vreinterpretq_u16_u8(vshrq_n_u8(vreinterpretq_u8_u64(padding), 7));
        uint32x4_t paired16 = vreinterpretq_u32_u16(vsraq_n_u16(high_bits, high_bits, 7));
        uint64x2_t paired32 = vreinterpretq_u64_u32(vsraq_n_u32(paired16, paired16, 14));
        uint8x16_t paired64 = vreinterpretq_u8_u64(vsraq_n_u64(paired32, paired32, 28));
        *output |= (uint64_t) vgetq_lane_u8(paired64, 0) << 8 * 2 * chunks |
                   (uint64_t) vgetq_lane_u8(paired64, 8) << (2 * chunks + 1);

    }
}

void packed_mask2_neon(uint8x16_t *input, uint64_t *output, uint64_t value, uint64_t size) {
    size = (size > 8) ? 8 : size;
    uint64_t chunks = size / 2;
    uint64_t remainder = size % 2;

    for (uint64_t i = 0; i < chunks; i++) {
        uint8x16_t vInput = vceqq_u8(input[i], vdupq_n_u8(value));
        uint16x8_t high_bits = vreinterpretq_u16_u8(vshrq_n_u8(vInput, 7));
        uint32x4_t paired16 = vreinterpretq_u32_u16(vsraq_n_u16(high_bits, high_bits, 7));
        uint64x2_t paired32 = vreinterpretq_u64_u32(vsraq_n_u32(paired16, paired16, 14));
        uint8x16_t paired64 = vreinterpretq_u8_u64(vsraq_n_u64(paired32, paired32, 28));
        *output |= (uint64_t) vgetq_lane_u8(paired64, 0) << 8 * 2 * i |
                   (uint64_t) vgetq_lane_u8(paired64, 8) << 8 * (2 * i + 1);
    }

    if (remainder > 0) {
        uint64x2_t padding = {0, 0};
        padding[0] = ((uint64_t *) input)[2 * chunks];
        uint8x16_t vPadding =  vceqq_u8(vreinterpretq_u8_u64(padding), vdupq_n_u8(value));
        uint16x8_t high_bits = vreinterpretq_u16_u8(vshrq_n_u8(vPadding, 7));
        uint32x4_t paired16 = vreinterpretq_u32_u16(vsraq_n_u16(high_bits, high_bits, 7));
        uint64x2_t paired32 = vreinterpretq_u64_u32(vsraq_n_u32(paired16, paired16, 14));
        uint8x16_t paired64 = vreinterpretq_u8_u64(vsraq_n_u64(paired32, paired32, 28));
        *output |= (uint64_t) vgetq_lane_u8(paired64, 0) << 8 * 2 * chunks |
                   (uint64_t) vgetq_lane_u8(paired64, 8) << (2 * chunks + 1);

    }
}
