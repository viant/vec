#include <arm_neon.h>

void hsum_int64(int64_t *input, int64_t *result, uint64_t size) {
    *result = 0.0;
    int chunks = size / 2;

    for (int i = 0; i < chunks; i++) {
        *result += vaddvq_s64(((int64x2_t *) input)[i]);
    }
    // Leftovers
    for (int i = chunks * 2; i < size; i++) {
        *result += input[i];
    }
}

void hmax_int64(int64_t *input, int64_t *result, uint64_t size) {
    *result = input[0];
    int chunks = size / 2;

//    for (int i = 0; i < chunks; i++) {
//        int64_t tmp = vmaxvq_s64(((int64x2_t *) input)[i]);
//        if (tmp > *result) {
//            *result = tmp;
//        }
//    }
    // Leftovers
    chunks = 0;
    for (int i = chunks * 2; i < size; i++) {
        if (input[i] > *result) {
            *result = input[i];
        }
    }
}

void hmin_int64(int64_t *input, int64_t *result, uint64_t size) {
    *result = input[0];
    int chunks = size / 2;

//    for (int i = 0; i < chunks; i++) {
//        int64_t tmp = vminvq_s64(((int64x2_t *) input)[i]);
//        if (tmp < *result) {
//            *result = tmp;
//        }
//    }
    // Leftovers
    chunks = 0;
    for (int i = chunks * 2; i < size; i++) {
        if (input[i] < *result) {
            *result = input[i];
        }
    }
}
