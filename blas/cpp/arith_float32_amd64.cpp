#include <immintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

void add_float32(float *input1, float *input2, float *output, uint64_t info) asm("add_float32");
void add_float32(float *input1, float *input2, float *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_ps(output + 8*i, _mm256_add_ps(_mm256_loadu_ps(input1 + 8*i), _mm256_loadu_ps(input2 + 8*i)));
        }
    }
    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

void sub_float32(float *input1, float *input2, float *output, uint64_t info) asm("sub_float32");
void sub_float32(float *input1, float *input2, float *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_ps(output + 8*i, _mm256_sub_ps(_mm256_loadu_ps(input1 + 8*i), _mm256_loadu_ps(input2 + 8*i)));
        }
    }
    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

void mul_float32(float *input1, float *input2, float *output, uint64_t info) asm("mul_float32");
void mul_float32(float *input1, float *input2, float *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_ps(output + 8*i, _mm256_mul_ps(_mm256_loadu_ps(input1 + 8*i), _mm256_loadu_ps(input2 + 8*i)));
        }
    }
    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        output[i] = input1[i] * input2[i];
    }
}

void div_float32(float *input1, float *input2, float *output, uint64_t info) asm("div_float32");
void div_float32(float *input1, float *input2, float *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_ps(output + 8*i, _mm256_div_ps(_mm256_loadu_ps(input1 + 8*i), _mm256_loadu_ps(input2 + 8*i)));
        }
    }
    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        output[i] = input1[i] / input2[i];
    }
}


