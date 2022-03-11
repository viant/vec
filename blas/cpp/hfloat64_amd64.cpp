#include <immintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

void hsum_float64(double *input, double *result, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    *result = 0.0;
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 4;
        for (int i = 0; i < chunks; i++) {
            __m256d vInput = _mm256_loadu_pd(input + 4 * i);
            __m128d vlow = _mm256_castpd256_pd128(vInput);
            __m128d vhigh = _mm256_extractf128_pd(vInput, 1);
            vlow = _mm_add_pd(vlow, vhigh);
            __m128d high64 = _mm_unpackhi_pd(vlow, vlow);
            *result += _mm_cvtsd_f64(_mm_add_sd(vlow, high64));
        }
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        *result += input[i];
    }
}

void hmax_float64(double *input, double *result, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    *result = input[0];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 4;
        for (int i = 0; i < chunks; i++) {
            __m256d vInput = _mm256_loadu_pd(input + 4 * i);
            __m128d vlow = _mm256_castpd256_pd128(vInput);
            __m128d vhigh = _mm256_extractf128_pd(vInput, 1);
            vlow = _mm_max_pd(vlow, vhigh);
            __m128d high64 = _mm_unpackhi_pd(vlow, vlow);
            double tmp = _mm_cvtsd_f64(_mm_max_sd(vlow, high64));
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

void hmin_float64(double *input, double *result, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    *result = input[0];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 4;
        for (int i = 0; i < chunks; i++) {
            __m256d vInput = _mm256_loadu_pd(input + 4 * i);
            __m128d vlow = _mm256_castpd256_pd128(vInput);
            __m128d vhigh = _mm256_extractf128_pd(vInput, 1);
            vlow = _mm_min_pd(vlow, vhigh);
            __m128d high64 = _mm_unpackhi_pd(vlow, vlow);
            double tmp = _mm_cvtsd_f64(_mm_min_sd(vlow, high64));
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
