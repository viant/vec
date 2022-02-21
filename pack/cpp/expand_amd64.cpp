#include <immintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

void expand_mask_avx(uint64_t *input, __m256i_u *output, uint64_t info) asm("expand_mask_avx");

void expand_mask_avx(uint64_t *input, __m256i_u *output, uint64_t info) {
    if (((uint32_t*)&info)[1] == AVX2) {
        for (int i = 0; i < 2; i++) {
            __m256i vmask = _mm256_set1_epi32(((uint32_t *) input)[i]);
            __m256i shuffle = _mm256_setr_epi64x(0x0000000000000000,
                                                 0x0101010101010101, 0x0202020202020202, 0x0303030303030303);
            vmask = _mm256_shuffle_epi8(vmask, shuffle);
            __m256i bit_mask = _mm256_set1_epi64x(0x7fbfdfeff7fbfdfe);
            vmask = _mm256_or_si256(vmask, bit_mask);
            output[i] = _mm256_cmpeq_epi8(vmask, _mm256_set1_epi64x(-1));
        }
        return;
    }

    // naive:
    for (int i = 0; i < 64; i++) {
        ((uint8_t *) output)[i] = ((*input & 1L << i) != 0) ? 0xff : 0;
    }
}
