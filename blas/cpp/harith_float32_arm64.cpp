#include <arm_neon.h>

void hsum_float32(float *input, float *result, uint64_t size) {
    *result = 0.0;
    int chunks = size / 4;

    for (int i = 0; i < chunks; i++) {
        *result += vaddvq_f32(((float32x4_t *) input)[i]);
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        *result += input[i];
    }
}

void hmax_float32(float *input, float *result, uint64_t size) {
    *result = input[0];
    int chunks = size / 4;

    for (int i = 0; i < chunks; i++) {
        float tmp = vminvq_f32(((float32x4_t *) input)[i]);
        if (tmp > *result) {
            *result = tmp;
        }
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        if (input[i] > *result) {
            *result = input[i];
        }
    }
}

void hmin_float32(float *input, float *result, uint64_t size) {
    *result = input[0];
    int chunks = size / 4;

    for (int i = 0; i < chunks; i++) {
        float tmp = vminvq_f32(((float32x4_t *) input)[i]);
        if (tmp < *result) {
            *result = tmp;
        }
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        if (input[i] < *result) {
            *result = input[i];
        }
    }
}
