#include <immintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

void add_float64(double *input1, double *input2, double *output, uint64_t info) asm("add_float64");
void add_float64(double *input1, double *input2, double *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 4;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_pd(output + 4*i, _mm256_add_pd(_mm256_loadu_pd(input1 + 4*i), _mm256_loadu_pd(input2 + 4*i)));
        }
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

void sub_float64(double *input1, double *input2, double *output, uint64_t info) asm("sub_float64");
void sub_float64(double *input1, double *input2, double *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 4;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_pd(output + 4*i, _mm256_sub_pd(_mm256_loadu_pd(input1 + 4*i), _mm256_loadu_pd(input2 + 4*i)));
        }
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

void mul_float64(double *input1, double *input2, double *output, uint64_t info) asm("mul_float64");
void mul_float64(double *input1, double *input2, double *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 4;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_pd(output + 4*i, _mm256_mul_pd(_mm256_loadu_pd(input1 + 4*i), _mm256_loadu_pd(input2 + 4*i)));
        }
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        output[i] = input1[i] * input2[i];
    }
}

void div_float64(double *input1, double *input2, double *output, uint64_t info) asm("div_float64");
void div_float64(double *input1, double *input2, double *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 4;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_pd(output + 4*i, _mm256_div_pd(_mm256_loadu_pd(input1 + 4*i), _mm256_loadu_pd(input2 + 4*i)));
        }
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        output[i] = input1[i] / input2[i];
    }
}


