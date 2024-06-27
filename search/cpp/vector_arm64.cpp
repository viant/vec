#include <arm_neon.h>
#include <cmath>


void magnitude_f32_neon(const float32_t* input, uint64_t size, float32_t* output) {
    float32x4_t magnitudeSqVec = vdupq_n_f32(0.0f);
    uint64_t i = 0;
    int chunks = size / 4;
    for (int j = 0; j < chunks; i += 4, j++) {
        float32x4_t vec = vld1q_f32(input + i);
        magnitudeSqVec = vfmaq_f32(magnitudeSqVec, vec, vec);
    }
    // Handle remaining elements
    for (; i < size; ++i) {
        float val = input[i];
        magnitudeSqVec += val * val;
    }
    float magnitudeSq = vaddvq_f32(magnitudeSqVec);
    *output =  std::sqrtf(magnitudeSq);
}

