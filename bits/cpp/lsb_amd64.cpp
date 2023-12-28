#include <immintrin.h>
#include <memory.h>

const uint32_t AVX2 = 1;

void lsb_avx2(uint64_t *in, uint64_t info, int64_t *out) asm("lsb_avx2");

void lsb_avx2(uint64_t *in, uint64_t info, int64_t *out) {
   uint32_t *infoPart = (uint32_t *) &info;
    uint32_t len = infoPart[0];
    uint32_t feature = infoPart[1];
    uint32_t parts = 0;
    *out = -1;

    if (feature == AVX2) {
        parts = len / 8;
        if (parts > 0) {
            for (int i = 0; i < parts; i++) {
                __m256i v1 = _mm256_loadu_si256((const __m256i *) &in[8 * i]);
                __m256i v2 = _mm256_loadu_si256((const __m256i *) &(in[8 * i + 4]));
                __m256i or1 = _mm256_or_si256(v1, v2);
                if (!_mm256_testz_si256(or1, or1)) {        // find the first non-zero cache line
                    __m256i v1z = _mm256_cmpeq_epi32(v1, _mm256_setzero_si256());
                    __m256i v2z = _mm256_cmpeq_epi32(v2, _mm256_setzero_si256());
                    uint32_t zero_map = _mm256_movemask_ps(_mm256_castsi256_ps(v1z));
                    zero_map |= _mm256_movemask_ps(_mm256_castsi256_ps(v2z)) << 8;

                    unsigned idx = __builtin_ctz(~zero_map);
                    uint32_t nonzero_chunk;
                    memcpy(&nonzero_chunk, (char *) &in[8 * i] + 4 * idx, sizeof(nonzero_chunk));

                    *out = (64 * i + 4 * idx) * 8 + __builtin_ctz(nonzero_chunk);
                    return;
                }
            }
        }
    }

    // Leftover:
    for (uint32_t i = parts * 8; i < len; i++) {
        if (in[i] != 0) {
            *out = 64 * i + __builtin_ctzll(in[i]);
            return;
        }
    }

    return;
}
