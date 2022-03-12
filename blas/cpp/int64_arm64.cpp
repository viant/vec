#include <arm_neon.h>
#include <stdint.h>

void add_int64(int64x2_t *input1, int64x2_t *input2, int64x2_t *output, uint64_t size) {
    int chunks = size / 2;

    for (int i = 0; i < chunks; i++) {
        output[i] = vaddq_s64(input1[i], input2[i]);
    }
    // Leftovers
    for (int i = chunks * 2; i < size; i++) {
        ((int64_t *) output)[i] = ((int64_t *) input1)[i] + ((int64_t *) input2)[i];
    }
}

void sub_int64(int64x2_t *input1, int64x2_t *input2, int64x2_t *output, uint64_t size) {
    int chunks = size / 2;

    for (int i = 0; i < chunks; i++) {
        output[i] = vsubq_s64(input1[i], input2[i]);
    }
    // Leftovers
    for (int i = chunks * 2; i < size; i++) {
        ((int64_t *) output)[i] = ((int64_t *) input1)[i] - ((int64_t *) input2)[i];
    }
}

void mul_int64(int64x2_t *input1, int64x2_t *input2, int64x2_t *output, uint64_t size) {
    int chunks = size / 2;

//    for (int i = 0; i < chunks; i++) {
//        output[i] = vmulq_s64(input1[i], input2[i]);
//    }
    // Leftovers
    chunks = 0;
    for (int i = chunks * 2; i < size; i++) {
        ((int64_t *) output)[i] = ((int64_t *) input1)[i] * ((int64_t *) input2)[i];
    }
}
