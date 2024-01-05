#include <arm_neon.h>
#include <memory.h>

void msb_neon(uint64_t *in, uint64_t size, int64_t *out) {
    *out = -1;

    int leftover = size % 4;
    int32x4_t shift = {0, 1, 2, 3};
    for (int i = size - 4; i >= leftover; i -= 4) {
        uint64x2_t v1 = vld1q_u64(in + i);
        uint64x2_t v2 = vld1q_u64(in + i + 2);
        uint32x4_t or1 = vreinterpretq_u64_u32(vorrq_u64(v1, v2));
        if (vmaxvq_u32(or1)) {
            uint32x4_t v1z = vceqq_u32(v1, vdupq_n_u32(0x00));
            uint32x4_t v2z = vceqq_u32(v2, vdupq_n_u32(0x00));

            uint32x4_t tmp = vshrq_n_u32(v1z, 31);
            uint32_t zero_map = vaddvq_u32(vshlq_u32(tmp, shift));
            tmp = vshrq_n_u32(v2z, 31);
            zero_map |= vaddvq_u32(vshlq_u32(tmp, shift)) << 4;
            uint32_t one_map = ~zero_map & 0xff;

            unsigned idx = 31 - __builtin_clz(one_map);
            uint32_t nonzero_chunk;
            memcpy(&nonzero_chunk, (char *) &in[i] + 4 * idx, sizeof(nonzero_chunk));

            *out = (8 * i + 4 * idx) * 8 + 31 - __builtin_clz(nonzero_chunk);
            return;
        }
    }

    // Leftover:
    for (int i = leftover - 1; i >= 0; i--) {
        if (in[i] != 0) {
            *out = 64 * i + 63 - __builtin_clzll(in[i]);
            return;
        }
    }
}
