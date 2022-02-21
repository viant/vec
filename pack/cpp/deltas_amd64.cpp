#include <smmintrin.h>

void compute_deltas(const int32_t * input, size_t length, int32_t * output, int32_t starting_point) asm("compute_deltas");
void compute_deltas(const int32_t * input, size_t length, int32_t * output, int32_t starting_point) {
    __m128i prev = _mm_set1_epi32(starting_point);
    size_t i = 0;
    for (; i < length / 4; i++) {
        __m128i curr = _mm_lddqu_si128((const __m128i *) input + i);
        __m128i delta = _mm_sub_epi32(curr,
                                      _mm_alignr_epi8(curr, prev, 12));
        _mm_storeu_si128((__m128i *) output + i, delta);
        prev = curr;
    }
    int32_t lastprev = _mm_extract_epi32(prev, 3);
    for (i = 4 * i; i < length; ++i) {
        int32_t curr = input[i];
        output[i] = curr - lastprev;
        lastprev = curr;
    }
}

void compute_prefix_sum(const int32_t *  input, size_t length, int32_t *  output, int32_t starting_point) asm("compute_prefix_sum");
void compute_prefix_sum(const int32_t *  input, size_t length, int32_t *  output, int32_t starting_point) {
    __m128i prev = _mm_set1_epi32(starting_point);
    size_t i = 0;
    for(; i  < length/4; i++) {
        __m128i curr =  _mm_lddqu_si128 (( const __m128i*) input + i );
        const __m128i _tmp1 = _mm_add_epi32(_mm_slli_si128(curr, 8), curr);
        const __m128i _tmp2 = _mm_add_epi32(_mm_slli_si128(_tmp1, 4), _tmp1);
        prev = _mm_add_epi32(_tmp2, _mm_shuffle_epi32(prev, 0xff));
        _mm_storeu_si128((__m128i*)output + i,prev);
    }
    int32_t lastprev = _mm_extract_epi32(prev,3);
    for(i = 4 * i; i < length; ++i) {
        lastprev = lastprev + input[i];
        output[i] = lastprev;
    }
}
