#include <arm_neon.h>
#include <stdint.h>


void add_scalar(int32x4_t *input, int32x4_t *output, int32_t value, uint64_t size) {
    int chunks = size / 4;

    int32x4_t vValue = vdupq_n_s32(value);
    for (int i = 0; i < chunks; i++) {
        output[i] = vaddq_s32(input[i], vdupq_n_s32(value));
    }

    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        ((uint32_t *) output)[i] = ((uint32_t *) input)[i] + value;
    }
}
