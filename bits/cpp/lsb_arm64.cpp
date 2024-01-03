#include <memory.h>
#include <arm_neon.h>

void lsb_neon(uint64_t *in, uint64_t size, int64_t *out) {
    uint32_t parts = 0;
    *out = -1;

    parts = size / 4;
    if (parts > 0) {
        for (int i = 0; i < parts; i++) {
            uint64x2_t v1 = vld1q_u64(in + 4 * i);
            uint64x2_t v2 = vld1q_u64(in + 4 * i + 2);
            uint32x4_t or1 = vreinterpretq_u64_u32(vorrq_u64(v1, v2));
            if (vmaxvq_u32(or1)) {
                uint32x4_t v1z = vceqq_u32(v1, vdupq_n_u32(0x00));
                uint32x4_t v2z = vceqq_u32(v2, vdupq_n_u32(0x00));

                static const int32x4_t shift = {0, 1, 2, 3};
                uint32x4_t tmp = vshrq_n_u32(v1z, 31);
                uint32_t zero_map = vaddvq_u32(vshlq_u32(tmp, shift));
                tmp = vshrq_n_u32(v2z, 31);
                zero_map |= vaddvq_u32(vshlq_u32(tmp, shift)) << 4;

                unsigned idx = __builtin_ctz(~zero_map);
                uint32_t nonzero_chunk;
                memcpy(&nonzero_chunk, (char *) &in[4 * i] + 4 * idx, sizeof(nonzero_chunk));

                *out = (32 * i + 4 * idx) * 8 + __builtin_ctz(nonzero_chunk);
                return;
            }
        }
    }

    // Leftover:
    for (uint32_t i = parts * 4; i < size; i++) {
        if (in[i] != 0) {
            *out = 64 * i + __builtin_ctzll(in[i]);
            return;
        }
    }
}
