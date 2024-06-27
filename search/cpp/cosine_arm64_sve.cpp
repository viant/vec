#include <arm_sve.h> // For SVE intrinsics
#include <cmath>     // For std::sqrt

void cosine_distance_f32_sve(const float* input1, const float* input2, uint64_t size, float* output) {
    uint64_t i = 0;
    svfloat32_t dotProductVec = svdup_f32(0.0f);
    svbool_t pg;
    svfloat32_t magnitudeSqVec1 = svdup_f32(0.0f);
    svfloat32_t magnitudeSqVec2 = svdup_f32(0.0f);

    int chunks = size / svcntw();
    for (int j = 0; j < chunks; i += svcntw(), j++) {
        pg = svwhilelt_b32(i, size);
        svfloat32_t vec1 = svld1(pg, input1 + i);
        svfloat32_t vec2 = svld1(pg, input2 + i);
        dotProductVec = svmla_f32_x(pg, dotProductVec, vec1, vec2);
        magnitudeSqVec1 = svmla_f32_x(pg, magnitudeSqVec1, vec1, vec1);
        magnitudeSqVec2 = svmla_f32_x(pg, magnitudeSqVec2, vec2, vec2);
    }

    float magnitudeSq1 = svaddv_f32(svptrue_b32(), magnitudeSqVec1);
    float magnitudeSq2 = svaddv_f32(svptrue_b32(), magnitudeSqVec2);
    float dotProduct = svaddv_f32(svptrue_b32(), dotProductVec);

    // Handle remaining elements
    for (; i < size; ++i) {
        float val1 = input1[i];
        float val2 = input2[i];
        dotProduct += val1 * val2;
        magnitudeSq1 += val1 * val1;
        magnitudeSq2 += val2 * val2;
    }
    magnitudeSq1 = std::sqrt(magnitudeSq1);
    magnitudeSq2 = std::sqrt(magnitudeSq2);
    if (dotProduct == 0 || magnitudeSq1 == 0 || magnitudeSq2 == 0) {
          *output = 1.0f;
     } else {
         *output = 1.0f - (dotProduct / (magnitudeSq1 * magnitudeSq2));
     }
}


void cosine_distance_with_magnitudes_f32_sve(const float* input1, const float* input2, const float* magnitudes, uint64_t size, float* output) {
    uint64_t i = 0;
    svfloat32_t dotProductVec = svdup_f32(0.0f);
    svbool_t pg;

    int chunks = size / svcntw();
    for (int j = 0; j < chunks; i += svcntw(), j++) {
        pg = svwhilelt_b32(i, size);
        svfloat32_t vec1 = svld1(pg, input1 + i);
        svfloat32_t vec2 = svld1(pg, input2 + i);
        dotProductVec = svmla_f32_x(pg, dotProductVec, vec1, vec2);
    }

    float dotProduct = svaddv_f32(svptrue_b32(), dotProductVec);

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
        *output = 1.0f - (dotProduct / (magnitude1 * magnitude2));
    }
}

