#include <immintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

void add_int64(int64_t *input1, int64_t *input2, int64_t *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 4;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_si256((__m256i *) (output + 4 * i),
                                _mm256_add_epi64(_mm256_loadu_si256((__m256i *) (input1 + 4 * i)),
                                                 _mm256_loadu_si256((__m256i *) (input2 + 4 * i))));
        }
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

void sub_int64(int64_t *input1, int64_t *input2, int64_t *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 4;
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_si256((__m256i *) (output + 4 * i),
                                _mm256_sub_epi64(_mm256_loadu_si256((__m256i *) (input1 + 4 * i)),
                                                 _mm256_loadu_si256((__m256i *) (input2 + 4 * i))));
        }
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

void mul_int64(int64_t *input1, int64_t *input2, int64_t *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 4;
        for (int i = 0; i < chunks; i++) {
            __m256i vec1 = _mm256_loadu_si256((__m256i *) (input1 + 4 * i));
            __m256i vec2 = _mm256_loadu_si256((__m256i *) (input2 + 4 * i));
            __m256i bswap = _mm256_shuffle_epi32(vec2, 0xB1);
            __m256i prodlh = _mm256_mullo_epi32(vec1, bswap);
            __m256i prodlh2 = _mm256_hadd_epi32(prodlh, _mm256_setzero_si256());
            __m256i prodlh3 = _mm256_shuffle_epi32(prodlh2, 0x73);
            __m256i prodll = _mm256_mul_epu32(vec1, vec2);
            _mm256_storeu_si256((__m256i *) (output + 4 * i), _mm256_add_epi64(prodll, prodlh3));
        }
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        output[i] = input1[i] * input2[i];
    }
}
