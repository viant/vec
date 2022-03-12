#include <arm_neon.h>

void hsum_float64(double *input, double *result, uint64_t size) {
    *result = 0.0;
    int chunks = size / 2;

    for (int i = 0; i < chunks; i++) {
        *result += vaddvq_f64(((float64x2_t *) input)[i]);
    }
    // Leftovers
    for (int i = chunks * 2; i < size; i++) {
        *result += input[i];
    }
}

void hmax_float64(double *input, double *result, uint64_t size) {
    *result = input[0];
    int chunks = size / 2;

    for (int i = 0; i < chunks; i++) {
        double tmp = vmaxvq_f64(((float64x2_t *) input)[i]);
        if (tmp > *result) {
            *result = tmp;
        }
    }
    // Leftovers
    for (int i = chunks * 2; i < size; i++) {
        if (input[i] > *result) {
            *result = input[i];
        }
    }
}

void hmin_float64(double *input, double *result, uint64_t size) {
    *result = input[0];
    int chunks = size / 2;

    for (int i = 0; i < chunks; i++) {
        double tmp = vminvq_f64(((float64x2_t *) input)[i]);
        if (tmp < *result) {
            *result = tmp;
        }
    }
    // Leftovers
    for (int i = chunks * 2; i < size; i++) {
        if (input[i] < *result) {
            *result = input[i];
        }
    }
}
