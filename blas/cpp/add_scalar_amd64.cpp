#include <immintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

void add_scalar(__m256i_u *input, __m256i_u *output, int32_t value, uint64_t info) asm("add_scalar");
void add_scalar(__m256i_u *input, __m256i_u *output, int32_t value, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 8;
        __m256i vValue = _mm256_set1_epi32(value);
        for (int i = 0; i < chunks; i++) {
            _mm256_storeu_si256(output + i, _mm256_add_epi32(_mm256_loadu_si256(input + i), vValue));
        }
    }

    // Leftovers
    for (int i = chunks * 8; i < size; i++) {
        ((uint32_t *) output)[i] = ((uint32_t *) input)[i] + value;
    }
}
