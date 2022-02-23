#include <immintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

void hsum_int32(int32_t *input, int32_t *result, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    *result = 0;
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            __m256i vInput = _mm256_loadu_si256((__m256i *) (input + 8 * i));
            __m128i sum128 = _mm_add_epi32(
                    _mm256_castsi256_si128(vInput),
                    _mm256_extracti128_si256(vInput, 1));
            __m128i hi64 = _mm_unpackhi_epi64(sum128, sum128);
            __m128i sum64 = _mm_add_epi32(hi64, sum128);
            __m128i hi32 = _mm_shuffle_epi32(sum64, _MM_SHUFFLE(2, 3, 0, 1));
            __m128i sum32 = _mm_add_epi32(sum64, hi32);
            *result += _mm_cvtsi128_si32(sum32);
        }
    }
    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        *result += input[i];
    }
}

void hmax_int32(int32_t *input, int32_t *result, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    *result = input[0];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            __m256i vInput = _mm256_loadu_si256((__m256i *) (input + 8 * i));
            __m128i sum128 = _mm_max_epi32(
                    _mm256_castsi256_si128(vInput),
                    _mm256_extracti128_si256(vInput, 1));
            __m128i hi64 = _mm_unpackhi_epi64(sum128, sum128);
            __m128i sum64 = _mm_max_epi32(hi64, sum128);
            __m128i hi32 = _mm_shuffle_epi32(sum64, _MM_SHUFFLE(2, 3, 0, 1));
            __m128i sum32 = _mm_max_epi32(sum64, hi32);
            int32_t tmp = _mm_cvtsi128_si32(sum32);
            if (tmp > *result) {
                *result = tmp;
            }
        }
    }
    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        if (input[i] > *result) {
            *result = input[i];
        }
    }
}

void hmin_int32(int32_t *input, int32_t *result, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    *result = input[0];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            __m256i vInput = _mm256_loadu_si256((__m256i *) (input + 8 * i));
            __m128i sum128 = _mm_min_epi32(
                    _mm256_castsi256_si128(vInput),
                    _mm256_extracti128_si256(vInput, 1));
            __m128i hi64 = _mm_unpackhi_epi64(sum128, sum128);
            __m128i sum64 = _mm_min_epi32(hi64, sum128);
            __m128i hi32 = _mm_shuffle_epi32(sum64, _MM_SHUFFLE(2, 3, 0, 1));
            __m128i sum32 = _mm_min_epi32(sum64, hi32);
            int32_t tmp = _mm_cvtsi128_si32(sum32);
            if (tmp < *result) {
                *result = tmp;
            }
        }
    }
    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        if (input[i] < *result) {
            *result = input[i];
        }
    }
}
