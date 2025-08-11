#include <arm_neon.h>


void find_nonzero_128bit_words(const uint64x2_t* V, uint32_t* mask_buf) {
    const uint32_t N = mask_buf[0];
    uint32_t* out = &mask_buf[3];
    uint32_t count = 0;

    for (uint32_t i = 0; i < N; ++i) {
        uint64x2_t v = V[i];
        uint64_t or_value = vgetq_lane_u64(v, 0) | vgetq_lane_u64(v, 1);
        if (or_value != 0) {
            out[count++] = i;
        }
    }

    mask_buf[2] = count;
}


void and_masked_128bit_neon(const uint64x2_t* A, const uint64x2_t* B, uint64x2_t* C, const uint32_t* mask_buf) {
    const uint32_t N = mask_buf[0];
    const uint32_t count = mask_buf[2];
    const uint32_t* indices = &mask_buf[3];

    for (uint32_t i = 0; i < count; ++i) {
        uint32_t idx = indices[i];
        uint64x2_t a = A[idx];
        uint64x2_t b = B[idx];
        C[idx] = vandq_u64(a, b);
    }
}
