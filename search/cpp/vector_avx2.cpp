#include <immintrin.h>
#include <cstdint>
#include <cmath>


void magnitude_f32_avx2(const float* input, uint64_t size, float* output) {
    __m256 magnitudeSqVec = _mm256_setzero_ps();
    uint64_t i = 0;
    int chunks = size / 8;
    for (int j = 0; j < chunks; i += 8, j++) {
        __m256 vec = _mm256_loadu_ps(input + i);
        magnitudeSqVec = _mm256_fmadd_ps(vec, vec, magnitudeSqVec);
    }
    float magnitudeSqArray[8];
    _mm256_storeu_ps(magnitudeSqArray, magnitudeSqVec);
    float magnitudeSq = 0.0f;
    for (int k = 0; k < 8; ++k) {
        magnitudeSq += magnitudeSqArray[k];
    }
    // Handle remaining elements
    for (; i < size; ++i) {
        float val = input[i];
        magnitudeSq += val * val;
    }
    *output = std::sqrt(magnitudeSq);
}
