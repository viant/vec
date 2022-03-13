#include <immintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

void hsum_int64(int64_t *input, int64_t *result, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    *result = 0;
    int chunks = 0;

    if (feature > 0) {
        chunks = size / 4;
        for (int i = 0; i < chunks; i++) {
            __m256i vInput = _mm256_loadu_si256((__m256i *) (input + 4 * i));
            __m128i sum = _mm_add_epi64(
                    _mm256_castsi256_si128(vInput),
                    _mm256_extracti128_si256(vInput, 1));
            sum = _mm_add_epi64(sum, _mm_shuffle_epi32(sum, _MM_SHUFFLE(1, 0, 3, 2)));
            *result += _mm_cvtsi128_si64(sum);
        }
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        *result += input[i];
    }
}

void hmax_int64(int64_t *input, int64_t *result, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    *result = input[0];
    int chunks = 0;

    if (feature == AVX512) {
        chunks = size / 4;
        for (int i = 0; i < chunks; i++) {
            __m256i vInput = _mm256_loadu_si256((__m256i *) (input + 4 * i));
            __m128i max = _mm_max_epi64(
                    _mm256_castsi256_si128(vInput),
                    _mm256_extracti128_si256(vInput, 1));
            max = _mm_max_epi64(max, _mm_shuffle_epi32(max, _MM_SHUFFLE(1, 0, 3, 2)));
            int64_t tmp = _mm_cvtsi128_si64(max);
            if (tmp > *result) {
                *result = tmp;
            }
        }
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        if (input[i] > *result) {
            *result = input[i];
        }
    }
}

void hmin_int64(int64_t *input, int64_t *result, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    *result = input[0];
    int chunks = 0;

    if (feature == AVX512) {
        chunks = size / 4;
        for (int i = 0; i < chunks; i++) {
            __m256i vInput = _mm256_loadu_si256((__m256i *) (input + 4 * i));
            __m128i min = _mm_min_epi64(
                    _mm256_castsi256_si128(vInput),
                    _mm256_extracti128_si256(vInput, 1));
            min = _mm_min_epi64(min, _mm_shuffle_epi32(min, _MM_SHUFFLE(1, 0, 3, 2)));
            int64_t tmp = _mm_cvtsi128_si64(min);
            if (tmp < *result) {
                *result = tmp;
            }
        }
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        if (input[i] < *result) {
            *result = input[i];
        }
    }
}
