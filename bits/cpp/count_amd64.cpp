
#include <iostream>

#include <immintrin.h>

const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;

/*
This implementation uses method described in  https://arxiv.org/pdf/1611.07612.pdf
*/

// Function below implement vectorized bit count
inline void CSA(__m256i *h, __m256i *l, __m256i a, __m256i b, __m256i c) {
    __m256i u = _mm256_xor_si256(a, b);
    *h = _mm256_or_si256(_mm256_and_si256(a, b), _mm256_and_si256(u, c));
    *l = _mm256_xor_si256(u, c);
}

inline __m256i count(__m256i v) {
    __m256i lookup = _mm256_setr_epi8(0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3,
                                      3, 4, 0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4);
    __m256i low_mask = _mm256_set1_epi8(0x0f);
    __m256i lo = _mm256_and_si256(v, low_mask);
    __m256i hi = _mm256_and_si256(_mm256_srli_epi32(v, 4), low_mask);
    __m256i popcnt1 = _mm256_shuffle_epi8(lookup, lo);
    __m256i popcnt2 = _mm256_shuffle_epi8(lookup, hi);
    __m256i total = _mm256_add_epi8(popcnt1, popcnt2);
    return _mm256_sad_epu8(total, _mm256_setzero_si256());

}

inline void count_sisd(const uint64_t *data, uint32_t offset, uint32_t size, uint64_t *out) {
    for (int i = offset; i < size; i++) {
        uint64_t x = data[i] - ((data[i] >> 1) & 0x5555555555555555);
        x = (x & 0x3333333333333333) + (x >> 2 & 0x3333333333333333);
        *out += ((x + (x >> 4) & 0xF0F0F0F0F0F0F0F) * 0x101010101010101) >> 56;
    }
}

void count_optimized(__m256_u *d, uint64_t info, uint64_t *result) asm("count_optimized");

void count_optimized(__m256_u *d, uint64_t info, uint64_t *result) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    uint32_t parts = 0;
    *result = 0;

    if (feature == AVX2) {
        parts = size / 64;
        if (parts > 0) {
            __m256i total = _mm256_setzero_si256();
            __m256i ones = _mm256_setzero_si256();
            __m256i twos = _mm256_setzero_si256();
            __m256i fours = _mm256_setzero_si256();
            __m256i eights = _mm256_setzero_si256();
            __m256i sixteens = _mm256_setzero_si256();
            __m256i twosA, twosB, foursA, foursB, eightsA, eightsB;
            for (int i = 0; i < parts * 16; i += 16) {
                CSA(&twosA, &ones, ones, d[i], d[i + 1]);
                CSA(&twosB, &ones, ones, d[i + 2], d[i + 3]);
                CSA(&foursA, &twos, twos, twosA, twosB);
                CSA(&twosA, &ones, ones, d[i + 4], d[i + 5]);
                CSA(&twosB, &ones, ones, d[i + 6], d[i + 7]);
                CSA(&foursB, &twos, twos, twosA, twosB);
                CSA(&eightsA, &fours, fours, foursA, foursB);
                CSA(&twosA, &ones, ones, d[i + 8], d[i + 9]);
                CSA(&twosB, &ones, ones, d[i + 10], d[i + 11]);
                CSA(&foursA, &twos, twos, twosA, twosB);
                CSA(&twosA, &ones, ones, d[i + 12], d[i + 13]);
                CSA(&twosB, &ones, ones, d[i + 14], d[i + 15]);
                CSA(&foursB, &twos, twos, twosA, twosB);
                CSA(&eightsB, &fours, fours, foursA, foursB);
                CSA(&sixteens, &eights, eights, eightsA, eightsB);
                total = _mm256_add_epi64(total, count(sixteens));
            }
            total = _mm256_slli_epi64(total, 4);
            total = _mm256_add_epi64(total, _mm256_slli_epi64(count(eights), 3));
            total = _mm256_add_epi64(total, _mm256_slli_epi64(count(fours), 2));
            total = _mm256_add_epi64(total, _mm256_slli_epi64(count(twos), 1));
            total = _mm256_add_epi64(total, count(ones));
            *result += _mm256_extract_epi64(total, 0) + _mm256_extract_epi64(total, 1)
                       + _mm256_extract_epi64(total, 2) + _mm256_extract_epi64(total, 3);
        }
    }

    // Leftover:
    count_sisd((uint64_t *) d, parts * 64, size, result);
//    for (int i = parts * 64; i < size; i++) {
//        *result += _mm_popcnt_u64(((uint64_t *) d)[i]);
//    }
}
