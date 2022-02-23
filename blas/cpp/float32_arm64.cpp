#include <arm_neon.h>
#include <stdint.h>

void add_float32(float32x4_t *input1, float32x4_t *input2, float32x4_t *output, uint64_t size) {
    int chunks = size / 4;

    for (int i = 0; i < chunks; i++) {
        output[i] = vaddq_f32(input1[i], input2[i]);
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        ((float32_t *) output)[i] = ((float32_t *) input1)[i] + ((float32_t *) input2)[i];
    }
}

void sub_float32(float32x4_t *input1, float32x4_t *input2, float32x4_t *output, uint64_t size) {
    int chunks = size / 4;

    for (int i = 0; i < chunks; i++) {
        output[i] = vsubq_f32(input1[i], input2[i]);
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        ((float32_t *) output)[i] = ((float32_t *) input1)[i] - ((float32_t *) input2)[i];
    }
}

void mul_float32(float32x4_t *input1, float32x4_t *input2, float32x4_t *output, uint64_t size) {
    int chunks = size / 4;

    for (int i = 0; i < chunks; i++) {
        output[i] = vmulq_f32(input1[i], input2[i]);
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        ((float32_t *) output)[i] = ((float32_t *) input1)[i] * ((float32_t *) input2)[i];
    }
}

void div_float32(float32x4_t *input1, float32x4_t *input2, float32x4_t *output, uint64_t size) {
    int chunks = size / 4;

    for (int i = 0; i < chunks; i++) {
        output[i] = vdivq_f32(input1[i], input2[i]);
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        ((float32_t *) output)[i] = ((float32_t *) input1)[i] / ((float32_t *) input2)[i];
    }
}
