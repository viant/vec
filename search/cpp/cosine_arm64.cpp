#include <arm_neon.h>
#include <stdint.h>
#include <cmath>

void cosine_distance_f32_neon(const float32_t* input1, const float32_t* input2, uint64_t size, float32_t* output) {
    float32x4_t dotProductVec = vdupq_n_f32(0.0f);
    float32x4_t magnitude1SqVec = vdupq_n_f32(0.0f);
    float32x4_t magnitude2SqVec = vdupq_n_f32(0.0f);
    uint64_t i = 0;

    int chunks = size / 4;
    for (int j = 0; j < chunks; i += 4, j++) {
        float32x4_t vec1 = vld1q_f32(input1 + i);
        float32x4_t vec2 = vld1q_f32(input2 + i);
        dotProductVec = vfmaq_f32(dotProductVec, vec1, vec2);
        magnitude1SqVec = vfmaq_f32(magnitude1SqVec, vec1, vec1);
        magnitude2SqVec = vfmaq_f32(magnitude2SqVec, vec2, vec2);
    }
    float dotProduct = vaddvq_f32(dotProductVec);
    float magnitude1Sq = vaddvq_f32(magnitude1SqVec);
    float magnitude2Sq = vaddvq_f32(magnitude2SqVec);

    // Handle remaining elements
    for (; i < size; ++i) {
        float val1 = input1[i];
        float val2 = input2[i];
        dotProduct += val1 * val2;
        magnitude1Sq += val1 * val1;
        magnitude2Sq += val2 * val2;
    }

    float magnitude1 = std::sqrt(magnitude1Sq);
    float magnitude2 = std::sqrt(magnitude2Sq);

    if (dotProduct == 0 || magnitude1 == 0 || magnitude2 == 0) {
        *output = 1.0f;
    } else {
        *output = 1.0f - dotProduct / (magnitude1 * magnitude2);
    }
    if (dotProduct == 0 || magnitude1 == 0 || magnitude2 == 0) {
        *output = 1.0f;
    } else {
        *output = 1.0f - ( dotProduct / (magnitude1 * magnitude2));
    }
}

void cosine_distance_with_magnitudes_f32_neon(const float32_t* input1, const float32_t* input2, const float32_t* magnitudes, uint64_t size, float32_t* output) {
    float32x4_t dotProductVec = vdupq_n_f32(0.0f);
    uint64_t i = 0;
    int chunks = size / 4;
    for (int j = 0; j < chunks; i += 4, j++) {
        float32x4_t vec1 = vld1q_f32(input1 + i);
        float32x4_t vec2 = vld1q_f32(input2 + i);
        dotProductVec = vfmaq_f32(dotProductVec, vec1, vec2);
    }
    float dotProduct = vaddvq_f32(dotProductVec);

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


