#include <immintrin.h>
#include <math.h>

void cosine_distance_f32_avx2(const float *a, const float *b, uint64_t size, float *output) asm("cosine_distance_f32_avx2");
void cosine_distance_f32_avx2(const float *a, const float *b, uint64_t size, float *output) {
    float ab = 0, a2 = 0, b2 = 0;
    for (int i = 0; i != size; ++i) {
            float ai = a[i];                                                       \
            float bi = b[i];                                                       \
            ab += ai * bi;                                                                                             \
            a2 += ai * ai;                                                                                             \
            b2 += bi * bi;                                                                                             \
    }
    *output = ab != 0 ? (1 - ab / sqrtf(a2) / sqrtf(b2)) : 1;
}

void cosine_distance_with_magnitudes_f32_avx2(const float *a, const float *b, const float *magnitudes,
                                               uint64_t size, float *output) asm("cosine_distance_with_magnitudes_f32_avx2");
void cosine_distance_with_magnitudes_f32_avx2(const float *a, const float *b, const float *magnitudes,
                                               uint64_t size, float *output) {
    float ab = 0, m1 = magnitudes[0], m2 = magnitudes[1];
    for (int i = 0; i != size; ++i) {
        float ai = a[i];
        float bi = b[i];
        ab += ai * bi;
    }
    *output = (ab != 0 && m1 != 0 && m2 != 0) ? (1 - ab / m1 / m2) : 1;
}
