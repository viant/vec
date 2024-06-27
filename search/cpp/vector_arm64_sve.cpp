#include <arm_sve.h> // For SVE intrinsics
#include <cmath>     // For std::sqrt


void magnitude_f32_sve(const float* input, uint64_t size, float* output) {
    uint64_t i = 0;
    svfloat32_t magnitudeSqVec = svdup_f32(0.0f);
    svbool_t pg;

    int chunks = size / svcntw();
    for (int j = 0; j < chunks; i += svcntw(), j++) {
        pg = svwhilelt_b32(i, size);
        svfloat32_t vec = svld1(pg, input + i);
        magnitudeSqVec = svmla_f32_x(pg, magnitudeSqVec, vec, vec);
    }

    float magnitudeSq = svaddv_f32(svptrue_b32(), magnitudeSqVec);
    // Handle remaining elements
    for (; i < size; ++i) {
        float val = input[i];
        magnitudeSq += val * val;
    }
    *output = std::sqrt(magnitudeSq);
}