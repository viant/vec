#include <immintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

void inc_int32(int32_t* vec, int32_t constant,  uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;
    __m256i constant_vec = _mm256_set1_epi32(constant);

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_si256((__m256i *) (vec + 8 * i),
                                _mm256_add_epi32(_mm256_loadu_si256((__m256i *) (vec + 8 * i)),
                                                 constant_vec));
        }
    }
    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        vec[i] = vec[i] + constant;
    }
}

void add_int32(int32_t *input1, int32_t *input2, int32_t *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_si256((__m256i *) (output + 8 * i),
                                _mm256_add_epi32(_mm256_loadu_si256((__m256i *) (input1 + 8 * i)),
                                                 _mm256_loadu_si256((__m256i *) (input2 + 8 * i))));
        }
    }
    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

void sub_int32(int32_t *input1, int32_t *input2, int32_t *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_si256((__m256i *) (output + 8 * i),
                                _mm256_sub_epi32(_mm256_loadu_si256((__m256i *) (input1 + 8 * i)),
                                                 _mm256_loadu_si256((__m256i *) (input2 + 8 * i))));
        }
    }
    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

void scalar_mul_int32(int32_t* vec, int32_t constant,  uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;
    __m256i constant_vec = _mm256_set1_epi32(constant);

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_si256((__m256i *) (vec + 8 * i),
                                _mm256_mullo_epi32(_mm256_loadu_si256((__m256i *) (vec + 8 * i)),
                                                 constant_vec));
        }
    }
    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        vec[i] = vec[i] * constant;
    }
}

void mul_int32(int32_t *input1, int32_t *input2, int32_t *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_si256((__m256i *) (output + 8 * i),
                                _mm256_mullo_epi32(_mm256_loadu_si256((__m256i *) (input1 + 8 * i)),
                                                 _mm256_loadu_si256((__m256i *) (input2 + 8 * i))));
        }
    }
    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        output[i] = input1[i] * input2[i];
    }
}
