#include <immintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

void hsum_float32(float *input, float *result, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    *result = 0.0;
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        __m128 sum;
        for (int i = 0; i < chunks; i++) {
            __m256 vInput = _mm256_loadu_ps(input + 8 * i);
            __m128 hiQuad = _mm256_extractf128_ps(vInput, 1);
            __m128 loQuad = _mm256_castps256_ps128(vInput);
            __m128 sumQuad = _mm_add_ps(loQuad, hiQuad);
            __m128 loDual = sumQuad;
            __m128 hiDual = _mm_movehl_ps(sumQuad, sumQuad);
            __m128 sumDual = _mm_add_ps(loDual, hiDual);
            __m128 lo = sumDual;
            __m128 hi = _mm_shuffle_ps(sumDual, sumDual, 0x1);
            sum = _mm_add_ss(lo, hi);
            *result += _mm_cvtss_f32(sum);
        }
     }
    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        *result += input[i];
    }
}

void hmax_float32(float *input, float *result, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    *result = input[0];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            __m256 vInput = _mm256_loadu_ps(input + 8 * i);
            __m128 hiQuad = _mm256_extractf128_ps(vInput, 1);
            __m128 loQuad = _mm256_castps256_ps128(vInput);
            __m128 maxQuad = _mm_max_ps(loQuad, hiQuad);
            __m128 loDual = maxQuad;
            __m128 hiDual = _mm_movehl_ps(maxQuad, maxQuad);
            __m128 maxDual = _mm_max_ps(loDual, hiDual);
            __m128 lo = maxDual;
            __m128 hi = _mm_shuffle_ps(maxDual, maxDual, 0x1);
            __m128 max = _mm_max_ss(lo, hi);
            float tmp = _mm_cvtss_f32(max);
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

void hmin_float32(float *input, float *result, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    *result = input[0];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        for (int i = 0; i < chunks; i++) {
            __m256 vInput = _mm256_loadu_ps(input + 8 * i);
            __m128 hiQuad = _mm256_extractf128_ps(vInput, 1);
            __m128 loQuad = _mm256_castps256_ps128(vInput);
            __m128 minQuad = _mm_min_ps(loQuad, hiQuad);
            __m128 loDual = minQuad;
            __m128 hiDual = _mm_movehl_ps(minQuad, minQuad);
            __m128 minDual = _mm_min_ps(loDual, hiDual);
            __m128 lo = minDual;
            __m128 hi = _mm_shuffle_ps(minDual, minDual, 0x1);
            __m128 min = _mm_min_ss(lo, hi);
            float tmp = _mm_cvtss_f32(min);
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
