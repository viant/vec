#include <arm_neon.h>
#include <arm_sve.h>

const uint32_t SVE = 3;

//_and_neon computes neon vectorized AND
inline void _and_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint32_t size) {
    int chunks = size / 2;
    for (int i = 0; i < chunks; i++) {
        out[i] = vandq_u64(v1[i], v2[i]);
    }
    int offset = chunks * 2;//handle remainder
    if (offset >= size) {
        return;
    }
    uint64_t *rOut = (uint64_t *) out;
    uint64_t *rV1 = (uint64_t *) v1;
    uint64_t *rV2 = (uint64_t *) v2;
    for (int i = offset; i < size; ++i) {
        rOut[i] = rV1[i] & rV2[i];
    }
}

//and_neon computes neon vectorized AND
void and_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint32_t size) {
    _and_neon(out, v1, v2, size);
}


//and_v3_neon computes neon vectorized AND
inline void _and_v3_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint32_t size) {
    int chunks = size / 2;
    for (int i = 0; i < chunks; i++) {
        out[i] = vandq_u64(v1[i], v2[i]);
        out[i] = vandq_u64(out[i], v3[i]);
    }
    int offset = chunks * 2;//handle remainder
    if (offset >= size) {
        return;
    }
    uint64_t *rOut = (uint64_t *) out;
    uint64_t *rV1 = (uint64_t *) v1;
    uint64_t *rV2 = (uint64_t *) v2;
    uint64_t *rV3 = (uint64_t *) v3;
    for (int i = offset; i < size; ++i) {
        rOut[i] = rV1[i] & rV2[i] & rV3[i];
    }
}

//and_v3_neon computes neon vectorized AND
void and_v3_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint32_t size) {
    _and_v3_neon(out, v1, v2, v3, size);
}



//_and_v4_neon computes neon vectorized AND
inline void _and_v4_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint64x2_t *v4, uint32_t size) {
    int chunks = size / 2;
    for (int i = 0; i < chunks; i++) {
        out[i] = vandq_u64(v1[i], v2[i]);
        out[i] = vandq_u64(out[i], v3[i]);
        out[i] = vandq_u64(out[i], v4[i]);
    }
    int offset = chunks * 2;//handle remainder
    if (offset >= size) {
        return;
    }
    uint64_t *rOut = (uint64_t *) out;
    uint64_t *rV1 = (uint64_t *) v1;
    uint64_t *rV2 = (uint64_t *) v2;
    uint64_t *rV3 = (uint64_t *) v3;
     uint64_t *rV4 = (uint64_t *) v4;
    for (int i = offset; i < size; ++i) {
        rOut[i] = rV1[i] & rV2[i] & rV3[i] & rV4[i];
    }
}

//and_v4_neon computes neon vectorized AND
void and_v4_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint64x2_t *v4, uint32_t size) {
    _and_v4_neon(out, v1, v2, v3, v4, size);
}


//_and_v5_neon computes neon vectorized AND
inline void _and_v5_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint64x2_t *v4, uint64x2_t *v5, uint32_t size) {
    int chunks = size / 2;
    for (int i = 0; i < chunks; i++) {
        out[i] = vandq_u64(v1[i], v2[i]);
        out[i] = vandq_u64(out[i], v3[i]);
        out[i] = vandq_u64(out[i], v4[i]);
        out[i] = vandq_u64(out[i], v5[i]);
    }
    int offset = chunks * 2;//handle remainder
    if (offset >= size) {
        return;
    }
    uint64_t *rOut = (uint64_t *) out;
    uint64_t *rV1 = (uint64_t *) v1;
    uint64_t *rV2 = (uint64_t *) v2;
    uint64_t *rV3 = (uint64_t *) v3;
    uint64_t *rV4 = (uint64_t *) v4;
    uint64_t *rV5 = (uint64_t *) v5;
    for (int i = offset; i < size; ++i) {
        rOut[i] = rV1[i] & rV2[i] & rV3[i] & rV4[i] & rV5[i];
    }
}


//and_v4_neon computes neon vectorized AND
void and_v5_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint64x2_t *v4, uint64x2_t *v5, uint32_t size) {
    _and_v5_neon(out, v1, v2, v3, v4, v5, size);
}


//_and_v6_neon computes neon vectorized AND
inline void _and_v6_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint64x2_t *v4, uint64x2_t *v5, uint64x2_t *v6, uint32_t size) {
    int chunks = size / 2;
    for (int i = 0; i < chunks; i++) {
        out[i] = vandq_u64(v1[i], v2[i]);
        out[i] = vandq_u64(out[i], v3[i]);
        out[i] = vandq_u64(out[i], v4[i]);
        out[i] = vandq_u64(out[i], v5[i]);
        out[i] = vandq_u64(out[i], v6[i]);
    }
    int offset = chunks * 2;//handle remainder
    if (offset >= size) {
        return;
    }
    uint64_t *rOut = (uint64_t *) out;
    uint64_t *rV1 = (uint64_t *) v1;
    uint64_t *rV2 = (uint64_t *) v2;
    uint64_t *rV3 = (uint64_t *) v3;
    uint64_t *rV4 = (uint64_t *) v4;
    uint64_t *rV5 = (uint64_t *) v5;
    uint64_t *rV6 = (uint64_t *) v6;
    for (int i = offset; i < size; ++i) {
        rOut[i] = rV1[i] & rV2[i] & rV3[i] & rV4[i] & rV5[i] & rV6[i];
    }
}


//and_v6_neon computes neon vectorized AND
void and_v6_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint64x2_t *v4, uint64x2_t *v5, uint64x2_t *v6, uint32_t size) {
    _and_v6_neon(out, v1, v2, v3, v4, v5, v6, size);
}



//_and_sve computes SVE vectorized AND
void _and_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t size) {
    uint64_t i = 0;
    svuint64_t vOut, vV1, vV2;
    svbool_t pred = svwhilelt_b64(i, size);
    while (svptest_first(svptrue_b64(), pred)) {
        vV1 = svld1_u64(pred, &v1[i]);
        vV2 = svld1_u64(pred, &v2[i]);
        vOut = svand_u64_x(pred, vV1, vV2);
        svst1_u64(pred, &out[i], vOut);
        i += svcntd();
        pred = svwhilelt_b64(i, size);
    }
}

//and_sve computes SVE vectorized AND
void and_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t size) {
    _and_sve(out, v1, v2, size);
}

//and_v3_sve computes SVE vectorized AND
inline void _and_v3_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t size) {
    uint64_t i = 0;
    svuint64_t vOut, vV1, vV2, vV3;
    svbool_t pred = svwhilelt_b64(i, size);
    while (svptest_first(svptrue_b64(), pred)) {
        vV1 = svld1_u64(pred, &v1[i]);
        vV2 = svld1_u64(pred, &v2[i]);
        vV3 = svld1_u64(pred, &v3[i]);
        vOut = svand_u64_x(pred, vV1, vV2);
        vOut = svand_u64_x(pred, vOut, vV3);
        svst1_u64(pred, &out[i], vOut);
        i += svcntd();
        pred = svwhilelt_b64(i, size);
    }
}


//and_v3_sve computes SVE vectorized AND
void and_v3_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t size) {
    _and_v3_sve(out, v1, v2, v3, size);
}



//_and_v4_sve computes SVE vectorized AND
inline void _and_v4_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t size) {
    uint64_t i = 0;
    svuint64_t vOut, vV1, vV2, vV3, vV4;
    svbool_t pred = svwhilelt_b64(i, size);
    while (svptest_first(svptrue_b64(), pred)) {
        vV1 = svld1_u64(pred, &v1[i]);
        vV2 = svld1_u64(pred, &v2[i]);
        vV3 = svld1_u64(pred, &v3[i]);
        vV4 = svld1_u64(pred, &v4[i]);
        vOut = svand_u64_x(pred, vV1, vV2);
        vOut = svand_u64_x(pred, vOut, vV3);
        vOut = svand_u64_x(pred, vOut, vV4); 
        svst1_u64(pred, &out[i], vOut);
        i += svcntd();
        pred = svwhilelt_b64(i, size);
    }
}


//and_v3_sve computes SVE vectorized AND
void and_v4_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t size) {
    _and_v4_sve(out, v1, v2, v3, v4, size);
}


//_and_v4_sve computes SVE vectorized AND
inline void _and_v5_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t *v5, uint64_t size) {
    uint64_t i = 0;
    svuint64_t vOut, vV1, vV2, vV3, vV4, vV5;
    svbool_t pred = svwhilelt_b64(i, size);
    while (svptest_first(svptrue_b64(), pred)) {
        vV1 = svld1_u64(pred, &v1[i]);
        vV2 = svld1_u64(pred, &v2[i]);
        vV3 = svld1_u64(pred, &v3[i]);
        vV4 = svld1_u64(pred, &v4[i]);
        vV5 = svld1_u64(pred, &v5[i]); 
        vOut = svand_u64_x(pred, vV1, vV2);
        vOut = svand_u64_x(pred, vOut, vV3);
        vOut = svand_u64_x(pred, vOut, vV4); 
        vOut = svand_u64_x(pred, vOut, vV5);  
        svst1_u64(pred, &out[i], vOut);
        i += svcntd();
        pred = svwhilelt_b64(i, size);
    }
}


//and_v5_sve computes SVE vectorized AND
void and_v5_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t *v5, uint64_t size) {
    _and_v5_sve(out, v1, v2, v3, v4, v5, size);
}

//_and_v4_sve computes SVE vectorized AND
inline void _and_v6_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t *v5, uint64_t *v6, uint64_t size) {
    uint64_t i = 0;
    svuint64_t vOut, vV1, vV2, vV3, vV4, vV5, vV6;
    svbool_t pred = svwhilelt_b64(i, size);
    while (svptest_first(svptrue_b64(), pred)) {
        vV1 = svld1_u64(pred, &v1[i]);
        vV2 = svld1_u64(pred, &v2[i]);
        vV3 = svld1_u64(pred, &v3[i]);
        vV4 = svld1_u64(pred, &v4[i]);
        vV5 = svld1_u64(pred, &v5[i]); 
        vV6 = svld1_u64(pred, &v6[i]);  
        vOut = svand_u64_x(pred, vV1, vV2);
        vOut = svand_u64_x(pred, vOut, vV3);
        vOut = svand_u64_x(pred, vOut, vV4); 
        vOut = svand_u64_x(pred, vOut, vV5);  
        vOut = svand_u64_x(pred, vOut, vV6);  
        svst1_u64(pred, &out[i], vOut);
        i += svcntd();
        pred = svwhilelt_b64(i, size);
    }
}


//and_v6_sve computes SVE vectorized AND
void and_v6_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t *v5, uint64_t *v6, uint64_t size) {
    _and_v6_sve(out, v1, v2, v3, v4, v5, v6, size);
}




//and_optimized computes optimized AND
void and_optimized(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature) {
        case SVE:
            _and_sve(out, v1, v2, (uint64_t)size);
            return;
    }
    _and_neon((uint64x2_t *)out, (uint64x2_t *)v1, (uint64x2_t *)v2, size);
}

//and_v3_optimized computes optimized AND
void and_v3_optimized(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature) {
        case SVE:
            _and_v3_sve(out, v1, v2, v3, (uint64_t)size);
            return;
    }
    _and_v3_neon((uint64x2_t *)out, (uint64x2_t *)v1, (uint64x2_t *)v2, (uint64x2_t *)v3, size);
}



//and_v4_optimized computes optimized AND
void and_v4_optimized(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature) {
        case SVE:
            _and_v4_sve(out, v1, v2, v3, v4, (uint64_t)size);
            return;
    }
    _and_v4_neon((uint64x2_t *)out, (uint64x2_t *)v1, (uint64x2_t *)v2, (uint64x2_t *)v3, (uint64x2_t *)v4, size);
}

//and_v5_optimized computes optimized AND
void and_v5_optimized(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t *v5, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature) {
        case SVE:
            _and_v5_sve(out, v1, v2, v3, v4, v5, (uint64_t)size);
            return;
    }
    _and_v5_neon((uint64x2_t *)out, (uint64x2_t *)v1, (uint64x2_t *)v2, (uint64x2_t *)v3, (uint64x2_t *)v4, (uint64x2_t *)v5,  size);
}



//and_v6_optimized computes optimized AND
void and_v6_optimized(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t *v5, uint64_t *v6, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature) {
        case SVE:
            _and_v6_sve(out, v1, v2, v3, v4, v5, v6, (uint64_t)size);
            return;
    }
    _and_v6_neon((uint64x2_t *)out, (uint64x2_t *)v1, (uint64x2_t *)v2, (uint64x2_t *)v3, (uint64x2_t *)v4, (uint64x2_t *)v5, (uint64x2_t *)v6, size);
}


inline void _or_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint32_t size) {
    int chunks = size / 2;
    for (int i = 0; i < chunks; i++) {
        out[i] = vorrq_u64(v1[i], v2[i]);
    }
    int offset = chunks * 2;//handle remainder
    if (offset >= size) {
        return;
    }
    uint64_t *rOut = (uint64_t *) out;
    uint64_t *rV1 = (uint64_t *) v1;
    uint64_t *rV2 = (uint64_t *) v2;
    for (int i = offset; i < size; ++i) {
        rOut[i] = rV1[i] | rV2[i];
    }
}

//or_neon computes neon vectorized OR
void or_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint32_t size) {
    _or_neon(out, v1, v2, size);
}


inline void _or_v3_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint32_t size) {
    int chunks = size / 2;
    for (int i = 0; i < chunks; i++) {
        out[i] = vorrq_u64(v1[i], v2[i]);
        out[i] = vorrq_u64(out[i], v3[i]);
    }
    int offset = chunks * 2;//handle remainder
    if (offset >= size) {
        return;
    }
    uint64_t *rOut = (uint64_t *) out;
    uint64_t *rV1 = (uint64_t *) v1;
    uint64_t *rV2 = (uint64_t *) v2;
    uint64_t *rV3 = (uint64_t *) v3;
    for (int i = offset; i < size; ++i) {
        rOut[i] = rV1[i] | rV2[i] | rV3[i];
    }
}

//or_v3_neon computes neon vectorized OR
void or_v3_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint32_t size) {
    _or_v3_neon(out, v1, v2, v3, size);
}

inline void _or_v4_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint64x2_t *v4, uint32_t size) {
    int chunks = size / 2;
    for (int i = 0; i < chunks; i++) {
        out[i] = vorrq_u64(v1[i], v2[i]);
        out[i] = vorrq_u64(out[i], v3[i]);
        out[i] = vorrq_u64(out[i], v4[i]);
    }
    int offset = chunks * 2;//handle remainder
    if (offset >= size) {
        return;
    }
    uint64_t *rOut = (uint64_t *) out;
    uint64_t *rV1 = (uint64_t *) v1;
    uint64_t *rV2 = (uint64_t *) v2;
    uint64_t *rV3 = (uint64_t *) v3;
    uint64_t *rV4 = (uint64_t *) v4;
    for (int i = offset; i < size; ++i) {
        rOut[i] = rV1[i] | rV2[i] | rV3[i]  | rV4[i];
    }
}


//or_v3_neon computes neon vectorized OR
void or_v4_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint64x2_t *v4, uint32_t size) {
    _or_v4_neon(out, v1, v2, v3, v4,  size);
}


inline void _or_v5_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint64x2_t *v4, uint64x2_t *v5, uint32_t size) {
    int chunks = size / 2;
    for (int i = 0; i < chunks; i++) {
        out[i] = vorrq_u64(v1[i], v2[i]);
        out[i] = vorrq_u64(out[i], v3[i]);
        out[i] = vorrq_u64(out[i], v4[i]);
        out[i] = vorrq_u64(out[i], v5[i]);
    }
    int offset = chunks * 2;//handle remainder
    if (offset >= size) {
        return;
    }
    uint64_t *rOut = (uint64_t *) out;
    uint64_t *rV1 = (uint64_t *) v1;
    uint64_t *rV2 = (uint64_t *) v2;
    uint64_t *rV3 = (uint64_t *) v3;
    uint64_t *rV4 = (uint64_t *) v4;
    uint64_t *rV5 = (uint64_t *) v5;
    for (int i = offset; i < size; ++i) {
        rOut[i] = rV1[i] | rV2[i] | rV3[i]  | rV4[i] | rV5[i];
    }
}

//or_v3_neon computes neon vectorized OR
void or_v5_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint64x2_t *v4, uint64x2_t *v5, uint32_t size) {
    _or_v5_neon(out, v1, v2, v3, v4, v5, size);
}


inline void _or_v6_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint64x2_t *v4, uint64x2_t *v5, uint64x2_t *v6, uint32_t size) {
    int chunks = size / 2;
    for (int i = 0; i < chunks; i++) {
        out[i] = vorrq_u64(v1[i], v2[i]);
        out[i] = vorrq_u64(out[i], v3[i]);
        out[i] = vorrq_u64(out[i], v4[i]);
        out[i] = vorrq_u64(out[i], v5[i]);
        out[i] = vorrq_u64(out[i], v6[i]);
    }
    int offset = chunks * 2;//handle remainder
    if (offset >= size) {
        return;
    }
    uint64_t *rOut = (uint64_t *) out;
    uint64_t *rV1 = (uint64_t *) v1;
    uint64_t *rV2 = (uint64_t *) v2;
    uint64_t *rV3 = (uint64_t *) v3;
    uint64_t *rV4 = (uint64_t *) v4;
    uint64_t *rV5 = (uint64_t *) v5;
    uint64_t *rV6 = (uint64_t *) v6;
    for (int i = offset; i < size; ++i) {
        rOut[i] = rV1[i] | rV2[i] | rV3[i]  | rV4[i] | rV5[i] | rV6[i];
    }
}

//or_v3_neon computes neon vectorized OR
void or_v6_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint64x2_t *v4, uint64x2_t *v5, uint64x2_t *v6, uint32_t size) {
    _or_v6_neon(out, v1, v2, v3, v4, v5, v6, size);
}


//_or_sve computes SVE vectorized OR
inline void _or_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t size) {
    uint64_t i = 0;
    svuint64_t vOut, vV1, vV2;
    svbool_t pred = svwhilelt_b64(i, size);
    while (svptest_first(svptrue_b64(), pred)) {
        vV1 = svld1_u64(pred, &v1[i]);
        vV2 = svld1_u64(pred, &v2[i]);
        vOut = svorr_u64_x(pred, vV1, vV2);
        svst1_u64(pred, &out[i], vOut);
        i += svcntd();
        pred = svwhilelt_b64(i, size);
    }
}

//or_sve computes SVE vectorized OR
void or_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t size) {
    _or_sve(out, v1, v2, size);
}

//or_v3_sve computes SVE vectorized OR
inline void _or_v3_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t size) {
    uint64_t i = 0;
    svuint64_t vOut, vV1, vV2, vV3;
    svbool_t pred = svwhilelt_b64(i, size);
    while (svptest_first(svptrue_b64(), pred)) {
        vV1 = svld1_u64(pred, &v1[i]);
        vV2 = svld1_u64(pred, &v2[i]);
        vV3 = svld1_u64(pred, &v3[i]);
        vOut = svorr_u64_x(pred, vV1, vV2);
        vOut = svorr_u64_x(pred, vOut, vV3);
        svst1_u64(pred, &out[i], vOut);
        i += svcntd();
        pred = svwhilelt_b64(i, size);
    }
}


//or_v3_sve computes SVE vectorized OR
void or_v3_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t size) {
    _or_v3_sve(out, v1, v2, v3, size);
}

//or_v3_sve computes SVE vectorized OR
inline void _or_v4_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t size) {
    uint64_t i = 0;
    svuint64_t vOut, vV1, vV2, vV3, vV4;
    svbool_t pred = svwhilelt_b64(i, size);
    while (svptest_first(svptrue_b64(), pred)) {
        vV1 = svld1_u64(pred, &v1[i]);
        vV2 = svld1_u64(pred, &v2[i]);
        vV3 = svld1_u64(pred, &v3[i]);
        vV4 = svld1_u64(pred, &v4[i]);
        vOut = svorr_u64_x(pred, vV1, vV2);
        vOut = svorr_u64_x(pred, vOut, vV3);
        vOut = svorr_u64_x(pred, vOut, vV4);
        svst1_u64(pred, &out[i], vOut);
        i += svcntd();
        pred = svwhilelt_b64(i, size);
    }
}



//or_v3_sve computes SVE vectorized OR
void or_v4_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t size) {
    _or_v4_sve(out, v1, v2, v3, v4, size);
}


//or_v3_sve computes SVE vectorized OR
inline void _or_v5_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t *v5, uint64_t size) {
    uint64_t i = 0;
    svuint64_t vOut, vV1, vV2, vV3, vV4, vV5;
    svbool_t pred = svwhilelt_b64(i, size);
    while (svptest_first(svptrue_b64(), pred)) {
        vV1 = svld1_u64(pred, &v1[i]);
        vV2 = svld1_u64(pred, &v2[i]);
        vV3 = svld1_u64(pred, &v3[i]);
        vV4 = svld1_u64(pred, &v4[i]);
        vV5 = svld1_u64(pred, &v5[i]);
        vOut = svorr_u64_x(pred, vV1, vV2);
        vOut = svorr_u64_x(pred, vOut, vV3);
        vOut = svorr_u64_x(pred, vOut, vV4);
        vOut = svorr_u64_x(pred, vOut, vV5);
        svst1_u64(pred, &out[i], vOut);
        i += svcntd();
        pred = svwhilelt_b64(i, size);
    }
}


//or_v3_sve computes SVE vectorized OR
void or_v5_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t *v5, uint64_t size) {
    _or_v5_sve(out, v1, v2, v3, v4,v5, size);
}



//or_v3_sve computes SVE vectorized OR
inline void _or_v6_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t *v5, uint64_t *v6, uint64_t size) {
    uint64_t i = 0;
    svuint64_t vOut, vV1, vV2, vV3, vV4, vV5, vV6;
    svbool_t pred = svwhilelt_b64(i, size);
    while (svptest_first(svptrue_b64(), pred)) {
        vV1 = svld1_u64(pred, &v1[i]);
        vV2 = svld1_u64(pred, &v2[i]);
        vV3 = svld1_u64(pred, &v3[i]);
        vV4 = svld1_u64(pred, &v4[i]);
        vV5 = svld1_u64(pred, &v5[i]);
        vV6 = svld1_u64(pred, &v6[i]);
        vOut = svorr_u64_x(pred, vV1, vV2);
        vOut = svorr_u64_x(pred, vOut, vV3);
        vOut = svorr_u64_x(pred, vOut, vV4);
        vOut = svorr_u64_x(pred, vOut, vV5);
        vOut = svorr_u64_x(pred, vOut, vV6);
        svst1_u64(pred, &out[i], vOut);
        i += svcntd();
        pred = svwhilelt_b64(i, size);
    }
}


//or_v3_sve computes SVE vectorized OR
void or_v6_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t *v4, uint64_t *v5, uint64_t *v6, uint64_t size) {
    _or_v6_sve(out, v1, v2, v3, v4,v5,v6, size);
}


//or_optimized computes optimized OR
void or_optimized(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature) {
        case SVE:
            _or_sve(out, v1, v2, (uint64_t)size);
            return;
    }
    _or_neon((uint64x2_t *)out, (uint64x2_t *)v1, (uint64x2_t *)v2, size);
}


//or_v3_optimized computes optimized OR
void or_v3_optimized(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature) {
        case SVE:
            _or_v3_sve(out, v1, v2, v3, (uint64_t)size);
            return;
    }
    _or_v3_neon((uint64x2_t *)out, (uint64x2_t *)v1, (uint64x2_t *)v2, (uint64x2_t *)v3, size);
}

//or_v4_optimized computes optimized OR
void or_v4_optimized(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3,  uint64_t *v4, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature) {
        case SVE:
            _or_v4_sve(out, v1, v2, v3, v4, (uint64_t)size);
            return;
    }
    _or_v4_neon((uint64x2_t *)out, (uint64x2_t *)v1, (uint64x2_t *)v2, (uint64x2_t *)v3,  (uint64x2_t *)v4, size);
}


//or_v5_optimized computes optimized OR
void or_v5_optimized(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3,  uint64_t *v4, uint64_t *v5, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature) {
        case SVE:
            _or_v5_sve(out, v1, v2, v3, v4, v5, (uint64_t)size);
            return;
    }
    _or_v5_neon((uint64x2_t *)out, (uint64x2_t *)v1, (uint64x2_t *)v2, (uint64x2_t *)v3,  (uint64x2_t *)v4,  (uint64x2_t *)v5,  size);
}



//or_v5_optimized computes optimized OR
void or_v6_optimized(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3,  uint64_t *v4, uint64_t *v5, uint64_t *v6, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature) {
        case SVE:
            _or_v6_sve(out, v1, v2, v3, v4, v5, v6, (uint64_t)size);
            return;
    }
    _or_v6_neon((uint64x2_t *)out, (uint64x2_t *)v1, (uint64x2_t *)v2, (uint64x2_t *)v3,  (uint64x2_t *)v4,  (uint64x2_t *)v5,  (uint64x2_t *)v6,  size);
}




//_xor_neon computes neon vectorized XOR
inline void _xor_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint32_t size) {
    int chunks = size / 2;
    for (int i = 0; i < chunks; i++) {
        out[i] = veorq_u64(v1[i], v2[i]);
    }
    int offset = chunks * 2;//handle remainder
    if (offset >= size) {
        return;
    }
    uint64_t *rOut = (uint64_t *) out;
    uint64_t *rV1 = (uint64_t *) v1;
    uint64_t *rV2 = (uint64_t *) v2;
    for (int i = offset; i < size; ++i) {
        (rOut[i]) = (rV1[i]) ^ (rV2[i]);
    }
}


//xor_neon computes neon vectorized XOR
void xor_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint32_t size) {
    _xor_neon(out, v1, v2, size);
}

//xor_v3_neon computes neon vectorized XOR
inline void _xor_v3_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint32_t size) {
    int chunks = size / 2;
    for (int i = 0; i < chunks; i++) {
        out[i] = veorq_u64(v1[i], v2[i]);
        out[i] = veorq_u64(out[i], v3[i]);
    }
    int offset = chunks * 2;//handle remainder
    if (offset >= size) {
        return;
    }
    uint64_t *rOut = (uint64_t *) out;
    uint64_t *rV1 = (uint64_t *) v1;
    uint64_t *rV2 = (uint64_t *) v2;
    uint64_t *rV3 = (uint64_t *) v3;
    for (int i = offset; i < size; ++i) {
        rOut[i] = rV1[i] ^ rV2[i] ^ rV3[i];
    }
}


//xor_v3_neon computes neon vectorized XOR
void xor_v3_neon(uint64x2_t *out, uint64x2_t *v1, uint64x2_t *v2, uint64x2_t *v3, uint32_t size) {
    _xor_v3_neon(out, v1, v2, v3, size);
}


//_xor_sve computes SVE vectorized XOR
inline void _xor_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t size) {
    uint64_t i = 0;
    svuint64_t vOut, vV1, vV2;
    svbool_t pred = svwhilelt_b64(i, size);
    while (svptest_first(svptrue_b64(), pred)) {
        vV1 = svld1_u64(pred, &v1[i]);
        vV2 = svld1_u64(pred, &v2[i]);
        vOut = sveor_u64_x(pred, vV1, vV2);
        svst1_u64(pred, &out[i], vOut);
        i += svcntd();
        pred = svwhilelt_b64(i, size);
    }
}


void xor_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t size) {
    _xor_sve(out, v1, v2, size);
}


//xor_v3_sve computes SVE vectorized XOR
inline void _xor_v3_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t size) {
    uint64_t i = 0;
    svuint64_t vOut, vV1, vV2, vV3;
    svbool_t pred = svwhilelt_b64(i, size);
    while (svptest_first(svptrue_b64(), pred)) {
        vV1 = svld1_u64(pred, &v1[i]);
        vV2 = svld1_u64(pred, &v2[i]);
        vV3 = svld1_u64(pred, &v3[i]);
        vOut = sveor_u64_x(pred, vV1, vV2);
        vOut = sveor_u64_x(pred, vOut, vV3);
        svst1_u64(pred, &out[i], vOut);
        i += svcntd();
        pred = svwhilelt_b64(i, size);
    }
}

void xor_v3_sve(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t size) {
    _xor_v3_sve(out, v1, v2, v3, size);
}


//xor_optimized computes optimized XOR
void xor_optimized(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature) {
        case SVE:
            _xor_sve(out, v1, v2, (uint64_t)size);
            return;
    }
    _xor_neon((uint64x2_t *)out, (uint64x2_t *)v1, (uint64x2_t *)v2, size);
}


//xor_v3_optimized computes optimized XOR
void xor_v3_optimized(uint64_t *out, uint64_t *v1, uint64_t *v2, uint64_t *v3, uint64_t info) {
    uint32_t *infoPart = (uint32_t *) &info;
    uint32_t size = infoPart[0];
    uint32_t feature = infoPart[1];
    switch (feature) {
        case SVE:
            _xor_v3_sve(out, v1, v2, v3, (uint64_t)size);
            return;
    }
    _xor_v3_neon((uint64x2_t *)out, (uint64x2_t *)v1, (uint64x2_t *)v2, (uint64x2_t *)v3, size);
}