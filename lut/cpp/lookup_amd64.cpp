#include <immintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

void lookup_512vbmi(const __m512i_u *input, __m512i_u *output, const __m512i_u *lut,
                       uint64_t info)  asm("lookup_512vbmi");

void lookup_512vbmi(const __m512i_u *input, __m512i_u *output, const __m512i_u *lut, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    uint32_t chunks = 0;

    if (feature == AVX512) {
        chunks = size / 64;
        for (int i = 0; i < chunks; i++) {
            const __mmask64 mask = _mm512_movepi8_mask(input[i]);
            output[i] = _mm512_or_epi64(_mm512_maskz_permutex2var_epi8(~mask, lut[0], input[i], lut[1]),
                                        _mm512_maskz_permutex2var_epi8(mask, lut[2], input[i], lut[3]));
        }
    }
    // leftovers
    for (int i = chunks * 16; i < size; i++) {
        ((uint8_t *) output)[i] = ((uint8_t *) lut)[((uint8_t *) input)[i]];
    }
}
