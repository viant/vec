#include <arm_neon.h>
#include <stdint.h>

void add_float64(float64x2_t *input1, float64x2_t *input2, float64x2_t *output, uint64_t size) {
    int chunks = size / 2;

    for (int i = 0; i < chunks; i++) {
        output[i] = vaddq_f64(input1[i], input2[i]);
    }
    // Leftovers
    for (int i = chunks * 2; i < size; i++) {
        ((float64_t *) output)[i] = ((float64_t *) input1)[i] + ((float64_t *) input2)[i];
    }
}

void sub_float64(float64x2_t *input1, float64x2_t *input2, float64x2_t *output, uint64_t size) {
    int chunks = size / 2;

    for (int i = 0; i < chunks; i++) {
        output[i] = vsubq_f64(input1[i], input2[i]);
    }
    // Leftovers
    for (int i = chunks * 2; i < size; i++) {
        ((float64_t *) output)[i] = ((float64_t *) input1)[i] - ((float64_t *) input2)[i];
    }
}

void mul_float64(float64x2_t *input1, float64x2_t *input2, float64x2_t *output, uint64_t size) {
    int chunks = size / 2;

    for (int i = 0; i < chunks; i++) {
        output[i] = vmulq_f64(input1[i], input2[i]);
    }
    // Leftovers
    for (int i = chunks * 2; i < size; i++) {
        ((float64_t *) output)[i] = ((float64_t *) input1)[i] * ((float64_t *) input2)[i];
    }
}

void div_float64(float64x2_t *input1, float64x2_t *input2, float64x2_t *output, uint64_t size) {
    int chunks = size / 2;

    for (int i = 0; i < chunks; i++) {
        output[i] = vdivq_f64(input1[i], input2[i]);
    }
    // Leftovers
    for (int i = chunks * 2; i < size; i++) {
        ((float64_t *) output)[i] = ((float64_t *) input1)[i] / ((float64_t *) input2)[i];
    }
}
