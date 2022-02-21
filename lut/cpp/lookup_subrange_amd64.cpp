#include <smmintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

#define mm_cmpgt_epu8(a, b) _mm_cmpgt_epi8( _mm_xor_si128(a, _mm_set1_epi8(-128)), _mm_xor_si128(b, _mm_set1_epi8(-128)))

void lookup_subrange_avx(__m128i_u *input, __m128i_u *output, char *range_upper, __m128i_u *lut, int ranges,
                         uint64_t info)asm("lookup_subrange_avx");

void
lookup_subrange_avx(__m128i_u *input, __m128i_u *output, char *range_upper, __m128i_u *lut, int ranges, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    int chunks = 0;

    if (feature == AVX2) {
        chunks = size / 16;
        for (int i = 0; i < chunks; i++) {
            __m128i vInput = _mm_loadu_si128(input + i);
            __m128i vLut = _mm_loadu_si128(lut);
            __m128i index = {0, 0};
            for (int j = 0; j < ranges; j++) {
                index = _mm_sub_epi8(index, mm_cmpgt_epu8(vInput, _mm_set1_epi8(range_upper[j])));
            }
            _mm_storeu_si128(output + i, _mm_shuffle_epi8(vLut, index));
        }
    }

    // Leftovers
    for (int i = chunks * 16; i < size; i++) {
        uint8_t n = 0;
        for (int j = 0; j < ranges; j++) {
            if (((uint8_t *) input)[i] > range_upper[j]) {
                n += 1;
            }
        }
        ((uint8_t *) output)[i] = ((uint8_t *) lut)[n];
    }
}
