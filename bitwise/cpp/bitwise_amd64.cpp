#include <immintrin.h>
#include <cstdint>


const uint32_t AVX2 = 1;
const uint32_t AVX512 = 2;


//and_avx2 compute vectorized AND
void and_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, uint32_t size) asm("and_avx2");

//and_avx512 compute vectorized AND
void and_avx512(__m512i *out, const __m512i *v1, const __m512i *v2, uint32_t size)   asm("and_avx512");

//and_optimized compute optimized AND based on available CPU info
void and_optimized(__m256i *out, const __m256i *v1, const __m256i *v2, uint64_t info)  asm("and_optimized");

//and_v3_avx2 compute vectorized AND
void and_v3_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint32_t size)  asm("and_v3_avx2");

//and_v3_avx512 compute vectorized AND
void and_v3_avx512(__m256i *out, const __m256i *v1, const  __m256i *v2, const __m256i *v3, uint32_t size)  asm("and_v3_avx512");

//and_v3_optimized compute optimized 3 arguments AND  based on available CPU info
void and_v3_optimized(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint64_t info)  asm("and_v3_optimized");


inline void and_sisd(uint64_t *out, const uint64_t *v1, const uint64_t *v2, uint32_t offset, uint32_t size) {
    //fallback to naive implmentation
    for (int i = offset; i < size; ++i) {
        out[i] = v1[i] & v2[i];
    }
}

inline void _and_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, uint32_t size) {
    uint32_t chunks = size / 4;
    uint32_t offset = chunks * 4;
    if (offset == size) {
        for (int i = 0; i < chunks; i++) {
            __m256i res = _mm256_and_si256(v1[i], v2[i]);
            _mm256_store_si256(&out[i], res);
        }
        return;
    }
    for (int i = 0; i < chunks; i++) {
        __m256i vec1 = _mm256_loadu_si256(v1+i);
        __m256i vec2 = _mm256_loadu_si256(v2+i);
        __m256i res = _mm256_and_si256(vec1, vec2);
        _mm256_storeu_si256(&out[i], res);
    }
    //handle remainder
    and_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, offset, size);
}

inline void _and_avx512(__m512i *out, const __m512i *v1, const __m512i *v2, uint32_t size)  {
    uint32_t chunks = size / 8;
    uint32_t offset = chunks * 8;
    if (offset == size) {
        for (int i = 0; i < size; i++) {
            __m512i res = _mm512_and_si512(v1[i], v2[i]);
            _mm512_store_si512(&out[i], res);
        }
        return;
    }
    for (int i = 0; i < chunks; i++) {
        __m512i vec1 = _mm512_loadu_si512(v1+i);
        __m512i vec2 = _mm512_loadu_si512(v2+i);
        __m512i res = _mm512_and_si512(vec1, vec2);
        _mm512_storeu_si512(&out[i], res);
    }
    //handle remainder
    and_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, offset, size);
}

void and_avx512(__m512i *out, const __m512i *v1, const __m512i *v2, uint32_t size)  {
    _and_avx512(out, v1, v2, size);
}

void and_optimized(__m256i *out, const __m256i *v1, const __m256i *v2, uint64_t info) {
    uint32_t *infoPart = (uint32_t*)&info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature)  {
        case AVX2:
            _and_avx2(out, v1, v2, size);
            return;
        case AVX512:
            _and_avx512((__m512i*)out, (__m512i*)v1, (__m512i*)v2, size);
            return;
    }
    and_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, 0, size);
}

void and_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, uint32_t size) {
    _and_avx2(out, v1, v2, size);
}

inline void and_v3_sisd(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, uint32_t offset, uint32_t size) {
    //fallback to naive implmentation
    for (int i = offset; i < size; ++i) {
        out[i] = v1[i] & v2[i] & v3[i];
    }
}

inline void _and_v3_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint32_t size) {
    int chunks = size / 4;
    int offset = chunks * 4;
    if (offset == size) {
        for (int i = 0; i < chunks; i++) {
            __m256i res = _mm256_and_si256(v1[i], v2[i]);
            res = _mm256_and_si256(res, v3[i]);
            _mm256_store_si256(&out[i], res);
        }
        return;
    }
    for (int i = 0; i < chunks; i++) {
        __m256i vec1 = _mm256_loadu_si256(v1+i);
        __m256i vec2 = _mm256_loadu_si256(v2+i);
        __m256i vec3 = _mm256_loadu_si256(v3+i);
        __m256i res = _mm256_and_si256(vec1, vec2);
        res = _mm256_and_si256(res, vec3);
        _mm256_storeu_si256(&out[i], res);
    }
    //handle remainder
    and_v3_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, (uint64_t*)v3,  offset, size);
}

inline void _and_v3_avx512(__m512i *out,  const __m512i *v1, const  __m512i *v2, const __m512i *v3, uint32_t size) {
    uint32_t chunks = size / 8;
    uint32_t offset = chunks * 8;
    if (offset == size) {
        for (int i = 0; i < size; i++) {
            __m512i res = _mm512_and_si512(v1[i], v2[i]);
            res = _mm512_and_si512(res, v3[i]);
            _mm512_store_si512(&out[i], res);
        }
        return;
    }
    for (int i = 0; i < chunks; i++) {
        __m512i vec1 = _mm512_loadu_si512(v1+i);
        __m512i vec2 = _mm512_loadu_si512(v2+i);
        __m512i vec3 = _mm512_loadu_si512(v3+i);
        __m512i res = _mm512_and_si512(vec1, vec2);
        res = _mm512_and_si512(res, vec3);
        _mm512_storeu_si512(&out[i], res);
    }
    //handle remainder
    and_v3_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2,  (uint64_t*)v3, offset, size);
}

void and_v3_avx512(__m256i *out, const __m256i *v1, const  __m256i *v2, const __m256i *v3, uint32_t size) {
    _and_v3_avx512((__m512i*)out, (__m512i*)v1, (__m512i*)v2, (__m512i*)v3, size);
}

void and_v3_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint32_t size) {
    _and_v3_avx2(out, v1, v2, v3, size);
}

void and_v3_optimized(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint64_t info) {
    uint32_t *infoPart = (uint32_t*)&info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature)  {
        case AVX2:
            _and_v3_avx2(out, v1, v2, v3, size);
            return;
         case AVX512:
            _and_v3_avx512((__m512i*)out, (__m512i*)v1, (__m512i*)v2, (__m512i*)v3, size);
            return;
    }
    and_v3_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, (uint64_t*)v3, 0, size);
}



//or_avx2 compute vectorized OR
void or_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, uint32_t size)  asm("or_avx2");

//or_avx512 compute vectorized OR
void or_avx512(__m512i *out, const __m512i *v1, const __m512i *v2, uint32_t size)   asm("or_avx512");

//or_optimized compute optimized OR based on available CPU info
void or_optimized(__m256i *out, const __m256i *v1, const __m256i *v2, uint64_t info)  asm("or_optimized");

//or_v3_avx2 compute vectorized OR
void or_v3_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint32_t size)  asm("or_v3_avx2");

//or_v3_avx512 compute vectorized OR
void or_v3_avx512(__m256i *out, const __m256i *v1, const  __m256i *v2, const __m256i *v3, uint32_t size)  asm("or_v3_avx512");

//or_v3_optimized compute optimized 3 arguments OR  based on available CPU info
void or_v3_optimized(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint64_t info)  asm("or_v3_optimized");

inline void or_sisd(uint64_t *out, const uint64_t *v1, const uint64_t *v2, uint32_t offset, uint32_t size) {
    //fallback to naive implmentation
    for (int i = offset; i < size; ++i) {
        out[i] = v1[i] | v2[i];
    }
}

inline void _or_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, uint32_t size) {
    uint32_t chunks = size / 4;
    uint32_t offset = chunks * 4;
    if (offset == size) {
        for (int i = 0; i < chunks; i++) {
            __m256i res = _mm256_or_si256(v1[i], v2[i]);
            _mm256_store_si256(&out[i], res);
        }
        return;
    }
    for (int i = 0; i < chunks; i++) {
        __m256i vec1 = _mm256_loadu_si256(v1+i);
        __m256i vec2 = _mm256_loadu_si256(v2+i);
        __m256i res = _mm256_or_si256(vec1, vec2);
        _mm256_storeu_si256(&out[i], res);
    }
    //horle remainder
    or_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, offset, size);
}

inline void _or_avx512(__m512i *out, const __m512i *v1, const __m512i *v2, uint32_t size)  {
    uint32_t chunks = size / 8;
    uint32_t offset = chunks * 8;
    if (offset == size) {
        for (int i = 0; i < size; i++) {
            __m512i res = _mm512_or_si512(v1[i], v2[i]);
            _mm512_store_si512(&out[i], res);
        }
        return;
    }
    for (int i = 0; i < chunks; i++) {
        __m512i vec1 = _mm512_loadu_si512(v1+i);
        __m512i vec2 = _mm512_loadu_si512(v2+i);
        __m512i res = _mm512_or_si512(vec1, vec2);
        _mm512_storeu_si512(&out[i], res);
    }
    //horle remainder
    or_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, offset, size);
}

void or_optimized(__m256i *out, const __m256i *v1, const __m256i *v2, uint64_t info) {
    uint32_t *infoPart = (uint32_t*)&info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature)  {
        case AVX2:
            _or_avx2(out, v1, v2, size);
            return;
        case AVX512:
            _or_avx512((__m512i*)out, (__m512i*)v1, (__m512i*)v2, size);
            return;
    }
    or_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, 0, size);
}

void or_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, uint32_t size) {
    _or_avx2(out, v1, v2, size);
}

inline void or_v3_sisd(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, uint32_t offset, uint32_t size) {
    //fallback to naive implmentation
    for (int i = offset; i < size; ++i) {
        out[i] = v1[i] | v2[i] | v3[i];
    }
}

inline void _or_v3_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint32_t size) {
    int chunks = size / 4;
    int offset = chunks * 4;
    if (offset == size) {
        for (int i = 0; i < chunks; i++) {
            __m256i res = _mm256_or_si256(v1[i], v2[i]);
            res = _mm256_or_si256(res, v3[i]);
            _mm256_store_si256(&out[i], res);
        }
        return;
    }
    for (int i = 0; i < chunks; i++) {
        __m256i vec1 = _mm256_loadu_si256(v1+i);
        __m256i vec2 = _mm256_loadu_si256(v2+i);
        __m256i vec3 = _mm256_loadu_si256(v3+i);
        __m256i res = _mm256_or_si256(vec1, vec2);
        res = _mm256_or_si256(res, vec3);
        _mm256_storeu_si256(&out[i], res);
    }
    //horle remainder
    or_v3_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, (uint64_t*)v3,  offset, size);
}

inline void _or_v3_avx512(__m512i *out,  const __m512i *v1, const  __m512i *v2, const __m512i *v3, uint32_t size) {
    uint32_t chunks = size / 8;
    uint32_t offset = chunks * 8;
    if (offset == size) {
        for (int i = 0; i < size; i++) {
            __m512i res = _mm512_or_si512(v1[i], v2[i]);
            res = _mm512_or_si512(res, v3[i]);
            _mm512_store_si512(&out[i], res);
        }
        return;
    }
    for (int i = 0; i < chunks; i++) {
        __m512i vec1 = _mm512_loadu_si512(v1+i);
        __m512i vec2 = _mm512_loadu_si512(v2+i);
        __m512i vec3 = _mm512_loadu_si512(v3+i);
        __m512i res = _mm512_or_si512(vec1, vec2);
        res = _mm512_or_si512(res, vec3);
        _mm512_storeu_si512(&out[i], res);
    }
    //horle remainder
    or_v3_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2,  (uint64_t*)v3, offset, size);
}

void or_v3_avx512(__m512i *out,  const __m512i *v1, const  __m512i *v2, const __m512i *v3, uint32_t size) {
    _or_v3_avx512((__m512i*)out, (__m512i*)v1, (__m512i*)v2, (__m512i*)v3, size);
}

void or_v3_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint32_t size) {
    _or_v3_avx2(out, v1, v2, v3, size);
}

void or_v3_optimized(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint64_t info) {
    uint32_t *infoPart = (uint32_t*)&info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature)  {
        case AVX2:
            _or_v3_avx2(out, v1, v2, v3, size);
            return;
         case AVX512:
            _or_v3_avx512((__m512i*)out, (__m512i*)v1, (__m512i*)v2, (__m512i*)v3, size);
            return;
    }
    or_v3_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, (uint64_t*)v3, 0, size);
}






//xor_avx2 compute vectorized XOR
void xor_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, uint32_t size)  asm("xor_avx2");

//xor_avx512 compute vectorized XOR
void xor_avx512(__m512i *out, const __m512i *v1, const __m512i *v2, uint32_t size)   asm("xor_avx512");

//xor_optimized compute optimized XOR based on available CPU info
void xor_optimized(__m256i *out, const __m256i *v1, const __m256i *v2, uint64_t info)  asm("xor_optimized");

//xor_v3_avx2 compute vectorized XOR
void xor_v3_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint32_t size)  asm("xor_v3_avx2");

//xor_v3_avx512 compute vectorized XOR
void xor_v3_avx512(__m256i *out, const __m256i *v1, const  __m256i *v2, const __m256i *v3, uint32_t size)  asm("xor_v3_avx512");

//xor_v3_optimized compute optimized 3 arguments XOR  based on available CPU info
void xor_v3_optimized(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint64_t info)  asm("xor_v3_optimized");

inline void xor_sisd(uint64_t *out, const uint64_t *v1, const uint64_t *v2, uint32_t offset, uint32_t size) {
    //fallback to naive implmentation
    for (int i = offset; i < size; ++i) {
        out[i] = v1[i] ^ v2[i];
    }
}

inline void _xor_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, uint32_t size) {
    uint32_t chunks = size / 4;
    uint32_t offset = chunks * 4;
    if (offset == size) {
        for (int i = 0; i < chunks; i++) {
            __m256i res = _mm256_xor_si256(v1[i], v2[i]);
            _mm256_store_si256(&out[i], res);
        }
        return;
    }
    for (int i = 0; i < chunks; i++) {
        __m256i vec1 = _mm256_loadu_si256(v1+i);
        __m256i vec2 = _mm256_loadu_si256(v2+i);
        __m256i res = _mm256_xor_si256(vec1, vec2);
        _mm256_storeu_si256(&out[i], res);
    }
    //hxorle remainder
    xor_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, offset, size);
}

inline void _xor_avx512(__m512i *out, const __m512i *v1, const __m512i *v2, uint32_t size)  {
    uint32_t chunks = size / 8;
    uint32_t offset = chunks * 8;
    if (offset == size) {
        for (int i = 0; i < size; i++) {
            __m512i res = _mm512_xor_si512(v1[i], v2[i]);
            _mm512_store_si512(&out[i], res);
        }
        return;
    }
    for (int i = 0; i < chunks; i++) {
        __m512i vec1 = _mm512_loadu_si512(v1+i);
        __m512i vec2 = _mm512_loadu_si512(v2+i);
        __m512i res = _mm512_xor_si512(vec1, vec2);
        _mm512_storeu_si512(&out[i], res);
    }
    //hxorle remainder
    xor_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, offset, size);
}

void xor_optimized(__m256i *out, const __m256i *v1, const __m256i *v2, uint64_t info) {
    uint32_t *infoPart = (uint32_t*)&info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature)  {
        case AVX2:
            _xor_avx2(out, v1, v2, size);
            return;
        case AVX512:
            _xor_avx512((__m512i*)out, (__m512i*)v1, (__m512i*)v2, size);
            return;
    }
    xor_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, 0, size);
}

void xor_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, uint32_t size) {
    _xor_avx2(out, v1, v2, size);
}

inline void xor_v3_sisd(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, uint32_t offset, uint32_t size) {
    //fallback to naive implmentation
    for (int i = offset; i < size; ++i) {
        out[i] = v1[i] ^ v2[i] ^ v3[i];
    }
}

inline void _xor_v3_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint32_t size) {
    int chunks = size / 4;
    int offset = chunks * 4;
    if (offset == size) {
        for (int i = 0; i < chunks; i++) {
            __m256i res = _mm256_xor_si256(v1[i], v2[i]);
            res = _mm256_xor_si256(res, v3[i]);
            _mm256_store_si256(&out[i], res);
        }
        return;
    }
    for (int i = 0; i < chunks; i++) {
        __m256i vec1 = _mm256_loadu_si256(v1+i);
        __m256i vec2 = _mm256_loadu_si256(v2+i);
        __m256i vec3 = _mm256_loadu_si256(v3+i);
        __m256i res = _mm256_xor_si256(vec1, vec2);
        res = _mm256_xor_si256(res, vec3);
        _mm256_storeu_si256(&out[i], res);
    }
    //hxorle remainder
    xor_v3_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, (uint64_t*)v3,  offset, size);
}

inline void _xor_v3_avx512(__m512i *out,  const __m512i *v1, const  __m512i *v2, const __m512i *v3, uint32_t size) {
    uint32_t chunks = size / 8;
    uint32_t offset = chunks * 8;
    if (offset == size) {
        for (int i = 0; i < size; i++) {
            __m512i res = _mm512_xor_si512(v1[i], v2[i]);
            res = _mm512_xor_si512(res, v3[i]);
            _mm512_store_si512(&out[i], res);
        }
        return;
    }
    for (int i = 0; i < chunks; i++) {
        __m512i vec1 = _mm512_loadu_si512(v1+i);
        __m512i vec2 = _mm512_loadu_si512(v2+i);
        __m512i vec3 = _mm512_loadu_si512(v3+i);
        __m512i res = _mm512_xor_si512(vec1, vec2);
        res = _mm512_xor_si512(res, vec3);
        _mm512_storeu_si512(&out[i], res);
    }
    //hxorle remainder
    xor_v3_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2,  (uint64_t*)v3, offset, size);
}

void xor_v3_avx512(__m512i *out,  const __m512i *v1, const  __m512i *v2, const __m512i *v3, uint32_t size) {
    _xor_v3_avx512((__m512i*)out, (__m512i*)v1, (__m512i*)v2, (__m512i*)v3, size);
}

void xor_v3_avx2(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint32_t size) {
    _xor_v3_avx2(out, v1, v2, v3, size);
}

void xor_v3_optimized(__m256i *out, const __m256i *v1, const __m256i *v2, const __m256i *v3, uint64_t info) {
    uint32_t *infoPart = (uint32_t*)&info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature)  {
        case AVX2:
            _xor_v3_avx2(out, v1, v2, v3, size);
            return;
         case AVX512:
            _xor_v3_avx512((__m512i*)out, (__m512i*)v1, (__m512i*)v2, (__m512i*)v3, size);
            return;
    }
    xor_v3_sisd((uint64_t*)out, (uint64_t*)v1, (uint64_t*)v2, (uint64_t*)v3, 0, size);
}
