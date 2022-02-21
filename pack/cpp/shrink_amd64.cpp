#include <immintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

void packed_mask_avx(__m256i_u *input, uint64_t *output, uint64_t size) asm("packed_mask_avx");
void packed_mask2_avx(__m256i_u *input, uint64_t *output, uint64_t value, uint64_t size) asm("packed_mask2_avx");

__attribute__((noinline))
void packed_mask_avx(__m256i_u *input, uint64_t *output, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    size = (size > 8) ? 8 : size;

    if (feature == AVX2) {
        uint64_t chunks = size / 4;
        uint64_t remainder = size % 4;

        switch (chunks) {
            case 1:
                *output = (uint32_t) _mm256_movemask_epi8(input[0]);
                break;
            case 2:
                *output = (uint32_t) _mm256_movemask_epi8(input[0]);
                uint64_t high = (uint32_t) _mm256_movemask_epi8(input[1]);
                *output |= high << 32;
                return;
        }
        if (remainder != 0) {
            // remainder:
            __m256i padding = {0, 0, 0, 0};
            for (int i = 0; i < remainder; i++) {
                padding[i] = input[chunks][i];
            }
            uint64_t remainder_bits = (uint32_t) _mm256_movemask_epi8(padding);
            if (chunks == 0) {
                *output = remainder_bits;
            } else {
                *output |= remainder_bits << 32;
            }
        }
        return;
    }

    // naive:
    *output = 0;
    for (int i = 0; i < 8 * size; i++) {
        uint8_t x = ((uint8_t *) input)[i];
        if ((((uint8_t *) input)[i] & 0x80) != 0) {
            *output |= 1L << i;
        }
    }
}

__attribute__((noinline))
void packed_mask2_avx(__m256i_u *input, uint64_t *output, uint64_t value, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    size = (size > 8) ? 8 : size;

    if (feature == AVX2) {
        uint64_t chunks = size / 4;
        uint64_t remainder = size % 4;

        switch (chunks) {
            case 1:
                *output = (uint32_t) _mm256_movemask_epi8(_mm256_cmpeq_epi8(input[0], _mm256_set1_epi8(value)));
                break;
            case 2:
                *output = (uint32_t) _mm256_movemask_epi8(_mm256_cmpeq_epi8(input[0], _mm256_set1_epi8(value)));
                uint64_t high = (uint32_t) _mm256_movemask_epi8(_mm256_cmpeq_epi8(input[1], _mm256_set1_epi8(value)));
                *output |= high << 32;
                return;
        }
        if (remainder != 0) {
            // remainder:
            __m256i padding = {0, 0, 0, 0};
            for (int i = 0; i < remainder; i++) {
                padding[i] = input[chunks][i];
            }
            uint64_t remainder_bits = (uint32_t) _mm256_movemask_epi8(
                    _mm256_cmpeq_epi8(padding, _mm256_set1_epi8(value)));
            if (chunks == 0) {
                *output = remainder_bits;
            } else {
                *output |= remainder_bits << 32;
            }
        }
        return;
    }

    // naive:
    *output = 0;
    for (int i = 0; i < 8 * size; i++) {
        if (((uint8_t *) input)[i] == value) {
            *output |= 1L << i;
        }
    }
}
