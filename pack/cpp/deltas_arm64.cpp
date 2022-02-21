#include <arm_neon.h>
#include <stdint.h>
#include <memory.h>

void compute_deltas(const int32_t *input, uint32_t length, int32_t *output, int32_t starting_point) {
    int32x4_t prev = vdupq_n_s32(starting_point);
    uint32_t i = 0;
    for (; i < length / 4; i++) {
        int32x4_t curr = ((int32x4_t *) input)[i];
        int32x4_t delta = vsubq_s32(curr, vextq_s32 (prev, curr, 3));
        ((int32x4_t *) output)[i] = delta;
        prev = curr;
    }
    //leftover
    int32_t lastprev = prev[3];
    for (i = 4 * i; i < length; ++i) {
        int32_t curr = input[i];
        output[i] = curr - lastprev;
        lastprev = curr;
    }
}

void compute_prefix_sum(const int32_t *input, uint32_t length, int32_t *output, int32_t starting_point) {
    int32_t prefix = starting_point;
    int32x4_t zeros = vdupq_n_s32(0);
    uint32_t i = 0;
    for (; i < length / 4; i++) {
        int32x4_t v = ((int32x4_t *) input)[i];
        v = vaddq_s32(v, vextq_s32(zeros, v, 3));
        v = vaddq_s32(v, vextq_s32(zeros, v, 2));
        ((int32x4_t *) output)[i] = vaddq_s32(vdupq_n_s32(prefix), v);
        prefix += v[3];
    }
    //leftover
    for (i = 4 * i; i < length; ++i) {
        prefix += input[i];
        output[i] = prefix;
    }
}
