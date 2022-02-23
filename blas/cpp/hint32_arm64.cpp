#include <arm_neon.h>

void hsum_int32(int32_t *input, int32_t *result, uint64_t size) {
    *result = 0.0;
    int chunks = size / 4;

    for (int i = 0; i < chunks; i++) {
        *result += vaddvq_s32(((int32x4_t *) input)[i]);
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        *result += input[i];
    }
}

void hmax_int32(int32_t *input, int32_t *result, uint64_t size) {
    *result = input[0];
    int chunks = size / 4;

    for (int i = 0; i < chunks; i++) {
        int32_t tmp = vmaxvq_s32(((int32x4_t *) input)[i]);
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

void hmin_int32(int32_t *input, int32_t *result, uint64_t size) {
    *result = input[0];
    int chunks = size / 4;

    for (int i = 0; i < chunks; i++) {
        int32_t tmp = vminvq_s32(((int32x4_t *) input)[i]);
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
