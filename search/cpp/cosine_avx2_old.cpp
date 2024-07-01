#include <immintrin.h>
#include <cstdint>
#include <cmath>

typedef float float32_t;

void cosine_distance_f32_avx2(const float32_t* input1, const float32_t* input2, uint64_t size, float32_t* output) asm("cosine_distance_f32_avx2");
void cosine_distance_f32_avx2(const float32_t* input1, const float32_t* input2, uint64_t size, float32_t* output) {
    __m256 dotProductVec = _mm256_setzero_ps();
    __m256 magnitude1SqVec = _mm256_setzero_ps();
    __m256 magnitude2SqVec = _mm256_setzero_ps();
    uint64_t i = 0;

    int chunks = size / 8;
    for (int j = 0; j < chunks; i += 8, j++) {
        __m256 vec1 = _mm256_loadu_ps(input1 + i);
        __m256 vec2 = _mm256_loadu_ps(input2 + i);
        dotProductVec = _mm256_fmadd_ps(vec1, vec2, dotProductVec);
        magnitude1SqVec = _mm256_fmadd_ps(vec1, vec1, magnitude1SqVec);
        magnitude2SqVec = _mm256_fmadd_ps(vec2, vec2, magnitude2SqVec);
    }

    // Horizontal reduction to sum up all elements in the vectors
    dotProductVec = _mm256_add_ps(dotProductVec, _mm256_permute2f128_ps(dotProductVec, dotProductVec, 1));
    dotProductVec = _mm256_hadd_ps(dotProductVec, dotProductVec);
    dotProductVec = _mm256_hadd_ps(dotProductVec, dotProductVec);
    float dotProduct = ((float*)&dotProductVec)[0] + ((float*)&dotProductVec)[4];

    magnitude1SqVec = _mm256_add_ps(magnitude1SqVec, _mm256_permute2f128_ps(magnitude1SqVec, magnitude1SqVec, 1));
    magnitude1SqVec = _mm256_hadd_ps(magnitude1SqVec, magnitude1SqVec);
    magnitude1SqVec = _mm256_hadd_ps(magnitude1SqVec, magnitude1SqVec);
    float magnitude1Sq = ((float*)&magnitude1SqVec)[0] + ((float*)&magnitude1SqVec)[4];

    magnitude2SqVec = _mm256_add_ps(magnitude2SqVec, _mm256_permute2f128_ps(magnitude2SqVec, magnitude2SqVec, 1));
    magnitude2SqVec = _mm256_hadd_ps(magnitude2SqVec, magnitude2SqVec);
    magnitude2SqVec = _mm256_hadd_ps(magnitude2SqVec, magnitude2SqVec);
    float magnitude2Sq = ((float*)&magnitude2SqVec)[0] + ((float*)&magnitude2SqVec)[4];

    // Handle remaining elements
    for (; i < size; ++i) {
        float val1 = input1[i];
        float val2 = input2[i];
        dotProduct += val1 * val2;
        magnitude1Sq += val1 * val1;
        magnitude2Sq += val2 * val2;
    }

    // Calculate reciprocal square roots using Newton-Raphson iterations
    __m128 magnitudesSq = _mm_set_ps(0.0f, 0.0f, magnitude2Sq, magnitude1Sq);
    __m128 rsqrtApprox = _mm_rsqrt_ps(magnitudesSq);

    // Refine the approximation
    rsqrtApprox = _mm_mul_ps(rsqrtApprox, _mm_rsqrt_ps(_mm_mul_ps(magnitudesSq, rsqrtApprox)));
    rsqrtApprox = _mm_mul_ps(rsqrtApprox, _mm_rsqrt_ps(_mm_mul_ps(magnitudesSq, rsqrtApprox)));

    // Compute the square roots using the refined reciprocal square roots
    __m128 sqrtApprox = _mm_mul_ps(magnitudesSq, rsqrtApprox);

    float magnitudes[4];
    _mm_storeu_ps(magnitudes, sqrtApprox);

    float magnitude1 = magnitudes[0];
    float magnitude2 = magnitudes[1];

    if (dotProduct == 0 || magnitude1 == 0 || magnitude2 == 0) {
        *output = 1.0f;
    } else {
        *output = dotProduct / (magnitude1 * magnitude2);
    }
}


void cosine_distance_with_magnitudes_f32_avx2(const float* input1, const float* input2, const float* magnitudes, uint64_t size, float* output) asm("cosine_distance_with_magnitudes_f32_avx2");
void cosine_distance_with_magnitudes_f32_avx2(const float* input1, const float* input2, const float* magnitudes, uint64_t size, float* output) {
    __m256 dotProductVec = _mm256_setzero_ps();
    uint64_t i = 0;
    int chunks = size / 8;
    for (int j = 0; j < chunks; i += 8, j++) {
        __m256 vec1 = _mm256_loadu_ps(input1 + i);
        __m256 vec2 = _mm256_loadu_ps(input2 + i);
        dotProductVec = _mm256_fmadd_ps(vec1, vec2, dotProductVec);
    }
    float dotProductArray[8];
    _mm256_storeu_ps(dotProductArray, dotProductVec);
    float dotProduct = 0.0f;
    for (int k = 0; k < 8; ++k) {
        dotProduct += dotProductArray[k];
    }

    // Handle remaining elements
    for (; i < size; ++i) {
        float val1 = input1[i];
        float val2 = input2[i];
        dotProduct += val1 * val2;
    }
    float magnitude1 = magnitudes[0];
    float magnitude2 = magnitudes[1];
    if (dotProduct == 0 || magnitude1 == 0 || magnitude2 == 0) {
        *output = 1.0f;
    } else {
        *output = dotProduct / (magnitude1 * magnitude2);
    }
}


