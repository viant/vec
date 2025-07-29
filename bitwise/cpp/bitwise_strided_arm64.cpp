#include <arm_neon.h>
#include <arm_sve.h>
#include <cstdint>

// Feature discriminator identical to non-strided implementation
const uint32_t SVE = 3;

// Helper: test whether all bits are zero in a 128-bit NEON register
inline bool neon_is_zero(uint64x2_t v) {
    uint64x2_t cmp = vceqq_u64(v, vdupq_n_u64(0));
    uint32x2_t narrow = vmovn_u64(cmp);
    return (vget_lane_u32(narrow, 0) & vget_lane_u32(narrow, 1)) == 0xFFFFFFFFu;
}

// ===================== Optimised dispatcher (NEON vs SVE) =====================

void and_strided_optimized(uint64_t* out, const uint64_t* v1, const uint64_t* v2, uint32_t* control) {
    if (control[1] == SVE) {
        and_sve_strided(out, v1, v2, control);
    } else {
        and_neon_strided(reinterpret_cast<uint64x2_t*>(out),
                         reinterpret_cast<const uint64x2_t*>(v1),
                         reinterpret_cast<const uint64x2_t*>(v2),
                         control);
    }
}

// ---- V3 Optimised ----
void and_v3_strided_optimized(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, uint32_t* control) {
    if (control[1] == SVE) {
        and3_sve_strided(out, v1, v2, v3, control);
    } else {
        and3_neon_strided(reinterpret_cast<uint64x2_t*>(out),
                          reinterpret_cast<const uint64x2_t*>(v1),
                          reinterpret_cast<const uint64x2_t*>(v2),
                          reinterpret_cast<const uint64x2_t*>(v3), control);
    }
}

void or_v3_strided_optimized(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, uint32_t* control) {
    if (control[1] == SVE) {
        or3_sve_strided(out, v1, v2, v3, control);
    } else {
        or3_neon_strided(reinterpret_cast<uint64x2_t*>(out),
                         reinterpret_cast<const uint64x2_t*>(v1),
                         reinterpret_cast<const uint64x2_t*>(v2),
                         reinterpret_cast<const uint64x2_t*>(v3), control);
    }
}

// ---- V4 Optimised ----
void and_v4_strided_optimized(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, uint32_t* control) {
    if (control[1] == SVE) {
        and4_sve_strided(out, v1, v2, v3, v4, control);
    } else {
        and4_neon_strided(reinterpret_cast<uint64x2_t*>(out),
                          reinterpret_cast<const uint64x2_t*>(v1),
                          reinterpret_cast<const uint64x2_t*>(v2),
                          reinterpret_cast<const uint64x2_t*>(v3),
                          reinterpret_cast<const uint64x2_t*>(v4), control);
    }
}

void or_v4_strided_optimized(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, uint32_t* control) {
    if (control[1] == SVE) {
        or4_sve_strided(out, v1, v2, v3, v4, control);
    } else {
        or4_neon_strided(reinterpret_cast<uint64x2_t*>(out),
                         reinterpret_cast<const uint64x2_t*>(v1),
                         reinterpret_cast<const uint64x2_t*>(v2),
                         reinterpret_cast<const uint64x2_t*>(v3),
                         reinterpret_cast<const uint64x2_t*>(v4), control);
    }
}

// ---- V5 Optimised ----
void and_v5_strided_optimized(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, const uint64_t* v5, uint32_t* control) {
    if (control[1] == SVE) {
        and5_sve_strided(out, v1, v2, v3, v4, v5, control);
    } else {
        and5_neon_strided(reinterpret_cast<uint64x2_t*>(out),
                          reinterpret_cast<const uint64x2_t*>(v1),
                          reinterpret_cast<const uint64x2_t*>(v2),
                          reinterpret_cast<const uint64x2_t*>(v3),
                          reinterpret_cast<const uint64x2_t*>(v4),
                          reinterpret_cast<const uint64x2_t*>(v5), control);
    }
}

void or_v5_strided_optimized(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, const uint64_t* v5, uint32_t* control) {
    if (control[1] == SVE) {
        or5_sve_strided(out, v1, v2, v3, v4, v5, control);
    } else {
        or5_neon_strided(reinterpret_cast<uint64x2_t*>(out),
                         reinterpret_cast<const uint64x2_t*>(v1),
                         reinterpret_cast<const uint64x2_t*>(v2),
                         reinterpret_cast<const uint64x2_t*>(v3),
                         reinterpret_cast<const uint64x2_t*>(v4),
                         reinterpret_cast<const uint64x2_t*>(v5), control);
    }
}

// ---- V6 Optimised ----
void and_v6_strided_optimized(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, const uint64_t* v5, const uint64_t* v6, uint32_t* control) {
    if (control[1] == SVE) {
        and6_sve_strided(out, v1, v2, v3, v4, v5, v6, control);
    } else {
        and6_neon_strided(reinterpret_cast<uint64x2_t*>(out),
                          reinterpret_cast<const uint64x2_t*>(v1),
                          reinterpret_cast<const uint64x2_t*>(v2),
                          reinterpret_cast<const uint64x2_t*>(v3),
                          reinterpret_cast<const uint64x2_t*>(v4),
                          reinterpret_cast<const uint64x2_t*>(v5),
                          reinterpret_cast<const uint64x2_t*>(v6), control);
    }
}

void or_v6_strided_optimized(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, const uint64_t* v5, const uint64_t* v6, uint32_t* control) {
    if (control[1] == SVE) {
        or6_sve_strided(out, v1, v2, v3, v4, v5, v6, control);
    } else {
        or6_neon_strided(reinterpret_cast<uint64x2_t*>(out),
                         reinterpret_cast<const uint64x2_t*>(v1),
                         reinterpret_cast<const uint64x2_t*>(v2),
                         reinterpret_cast<const uint64x2_t*>(v3),
                         reinterpret_cast<const uint64x2_t*>(v4),
                         reinterpret_cast<const uint64x2_t*>(v5),
                         reinterpret_cast<const uint64x2_t*>(v6), control);
    }
}

void or_strided_optimized(uint64_t* out, const uint64_t* v1, const uint64_t* v2, uint32_t* control) {
    if (control[1] == SVE) {
        or_sve_strided(out, v1, v2, control);
    } else {
        or_neon_strided(reinterpret_cast<uint64x2_t*>(out),
                        reinterpret_cast<const uint64x2_t*>(v1),
                        reinterpret_cast<const uint64x2_t*>(v2),
                        control);
    }
}



// Helper: test whether all bits are one in a 128-bit NEON register
inline bool neon_is_allones(uint64x2_t v) {
    uint64x2_t cmp = vceqq_u64(v, vdupq_n_u64(~0ULL));
    uint32x2_t narrow = vmovn_u64(cmp);
    return (vget_lane_u32(narrow, 0) & vget_lane_u32(narrow, 1)) == 0xFFFFFFFFu;
}

// ==== SVE helpers ====

// Returns true if every 64-bit element of an SVE vector is zero.
static inline bool sve_is_zero(svuint64_t v) {
    // svptest_any returns true when *any* lane has 1-bits.
    // Therefore, !svptest_any means the entire vector is zero.
    return !svptest_any(svptrue_b64(), v);
}

// Returns true if every bit of an SVE vector is one (0xFFFF_FFFF_FFFF_FFFF).
static inline bool sve_is_allones(svuint64_t v) {
    svbool_t all_pred = svcmpeq_n_u64(svptrue_b64(), v, ~0ULL);
    return svptest_all(svptrue_b64(), all_pred);
}

// Fallback scalar OR for remaining elements (no strides)
inline void or_scalar_tail(uint64_t* out, const uint64_t* v1, const uint64_t* v2,
                           uint32_t start, uint32_t size) {
    for (uint32_t i = start; i < size; ++i) {
        out[i] = v1[i] | v2[i];
    }
}

// Main strided vectorized OR with NEON. The stride is widened when the OR
// result becomes fully saturated (all bits set).
inline void or_neon_strided(uint64x2_t* out, const uint64x2_t* v1, const uint64x2_t* v2,
                            uint32_t* control) {
    uint32_t size = control[0];    // total number of 64-bit words
    uint32_t feature = control[1]; // cpu feature (unused yet)
    uint32_t i = control[2];       // starting word index
    uint32_t j = 2;                // stride index

    while (i + 1 < size && j + 1 < size) {
        // Compute OR on 2 × 64-bit packed into one 128-bit vector.
        uint64x2_t result = vorrq_u64(v1[i / 2], v2[i / 2]);
        out[i / 2] = result;

        // If fully saturated – widen next stride.
        if (neon_is_allones(result)) {
            control[j + 1]++;
        }

        j++;
        i += control[j];  // jump by (possibly widened) stride
    }

    // Scalar cleanup
    or_scalar_tail(reinterpret_cast<uint64_t*>(out),
                   reinterpret_cast<const uint64_t*>(v1),
                   reinterpret_cast<const uint64_t*>(v2),
                   i, size);
}


// ==== V3..V6 STRIDED AND IMPLEMENTATIONS ====

inline void and3_neon_strided(uint64x2_t* out, const uint64x2_t* v1, const uint64x2_t* v2, const uint64x2_t* v3,
                              uint32_t* control) {
    uint32_t size = control[0];
    uint32_t i = control[2];
    uint32_t j = 2;

    while (i + 1 < size && j + 1 < size) {
        uint64x2_t res = vandq_u64(v1[i/2], v2[i/2]);
        res = vandq_u64(res, v3[i/2]);
        out[i/2] = res;
        if (neon_is_zero(res)) {
            control[j+1]++;
        }
        j++; i += control[j];
    }
    and3_scalar_tail(reinterpret_cast<uint64_t*>(out),
                     reinterpret_cast<const uint64_t*>(v1),
                     reinterpret_cast<const uint64_t*>(v2),
                     reinterpret_cast<const uint64_t*>(v3),
                     i, size);
}


// ===================== V3..V6 STRIDED AND WITH SVE =====================

inline void and3_sve_strided(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, uint32_t* control) {
    uint32_t size = control[0];
    uint32_t i = control[2];
    uint32_t j = 2;
    while (i < size && j + 1 < size) {
        svbool_t pred = svwhilelt_b64(i, size);
        svuint64_t vec1 = svld1_u64(pred, &v1[i]);
        svuint64_t vec2 = svld1_u64(pred, &v2[i]);
        svuint64_t vec3 = svld1_u64(pred, &v3[i]);
        svuint64_t res  = svand_u64_x(pred, vec1, vec2);
        res = svand_u64_x(pred, res, vec3);
        svst1_u64(pred, &out[i], res);
        if (sve_is_zero(res)) control[j+1]++;
        j++; i += control[j];
    }
    and3_scalar_tail(out, v1, v2, v3, i, size);
}

inline void and4_sve_strided(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, uint32_t* control) {
    uint32_t size = control[0]; uint32_t i = control[2]; uint32_t j = 2;
    while (i < size && j + 1 < size) {
        svbool_t pred = svwhilelt_b64(i, size);
        svuint64_t r = svand_u64_x(pred, svld1_u64(pred, &v1[i]), svld1_u64(pred, &v2[i]));
        r = svand_u64_x(pred, r, svld1_u64(pred, &v3[i]));
        r = svand_u64_x(pred, r, svld1_u64(pred, &v4[i]));
        svst1_u64(pred, &out[i], r);
        if (sve_is_zero(r)) control[j+1]++;
        j++; i += control[j];
    }
    and4_scalar_tail(out, v1, v2, v3, v4, i, size);
}

inline void and5_sve_strided(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, const uint64_t* v5, uint32_t* control) {
    uint32_t size = control[0]; uint32_t i = control[2]; uint32_t j = 2;
    while (i < size && j + 1 < size) {
        svbool_t pred = svwhilelt_b64(i, size);
        svuint64_t r = svand_u64_x(pred, svld1_u64(pred, &v1[i]), svld1_u64(pred, &v2[i]));
        r = svand_u64_x(pred, r, svld1_u64(pred, &v3[i]));
        r = svand_u64_x(pred, r, svld1_u64(pred, &v4[i]));
        r = svand_u64_x(pred, r, svld1_u64(pred, &v5[i]));
        svst1_u64(pred, &out[i], r);
        if (sve_is_zero(r)) control[j+1]++;
        j++; i += control[j];
    }
    and5_scalar_tail(out, v1, v2, v3, v4, v5, i, size);
}

inline void and6_sve_strided(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, const uint64_t* v5, const uint64_t* v6, uint32_t* control) {
    uint32_t size = control[0]; uint32_t i = control[2]; uint32_t j = 2;
    while (i < size && j + 1 < size) {
        svbool_t pred = svwhilelt_b64(i, size);
        svuint64_t r = svand_u64_x(pred, svld1_u64(pred, &v1[i]), svld1_u64(pred, &v2[i]));
        r = svand_u64_x(pred, r, svld1_u64(pred, &v3[i]));
        r = svand_u64_x(pred, r, svld1_u64(pred, &v4[i]));
        r = svand_u64_x(pred, r, svld1_u64(pred, &v5[i]));
        r = svand_u64_x(pred, r, svld1_u64(pred, &v6[i]));
        svst1_u64(pred, &out[i], r);
        if (sve_is_zero(r)) control[j+1]++;
        j++; i += control[j];
    }
    and6_scalar_tail(out, v1, v2, v3, v4, v5, v6, i, size);
}

// ===================== V3..V6 STRIDED OR WITH SVE =====================

inline void or3_sve_strided(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, uint32_t* control) {
    uint32_t size = control[0]; uint32_t i = control[2]; uint32_t j = 2;
    while (i < size && j + 1 < size) {
        svbool_t pred = svwhilelt_b64(i, size);
        svuint64_t r = svorr_u64_x(pred, svld1_u64(pred, &v1[i]), svld1_u64(pred, &v2[i]));
        r = svorr_u64_x(pred, r, svld1_u64(pred, &v3[i]));
        svst1_u64(pred, &out[i], r);
        if (sve_is_allones(r)) control[j+1]++;
        j++; i += control[j];
    }
    or3_scalar_tail(out, v1, v2, v3, i, size);
}

inline void or4_sve_strided(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, uint32_t* control) {
    uint32_t size = control[0]; uint32_t i = control[2]; uint32_t j = 2;
    while (i < size && j + 1 < size) {
        svbool_t pred = svwhilelt_b64(i, size);
        svuint64_t r = svorr_u64_x(pred, svld1_u64(pred, &v1[i]), svld1_u64(pred, &v2[i]));
        r = svorr_u64_x(pred, r, svld1_u64(pred, &v3[i]));
        r = svorr_u64_x(pred, r, svld1_u64(pred, &v4[i]));
        svst1_u64(pred, &out[i], r);
        if (sve_is_allones(r)) control[j+1]++;
        j++; i += control[j];
    }
    or4_scalar_tail(out, v1, v2, v3, v4, i, size);
}

inline void or5_sve_strided(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, const uint64_t* v5, uint32_t* control) {
    uint32_t size = control[0]; uint32_t i = control[2]; uint32_t j = 2;
    while (i < size && j + 1 < size) {
        svbool_t pred = svwhilelt_b64(i, size);
        svuint64_t r = svorr_u64_x(pred, svld1_u64(pred, &v1[i]), svld1_u64(pred, &v2[i]));
        r = svorr_u64_x(pred, r, svld1_u64(pred, &v3[i]));
        r = svorr_u64_x(pred, r, svld1_u64(pred, &v4[i]));
        r = svorr_u64_x(pred, r, svld1_u64(pred, &v5[i]));
        svst1_u64(pred, &out[i], r);
        if (sve_is_allones(r)) control[j+1]++;
        j++; i += control[j];
    }
    or5_scalar_tail(out, v1, v2, v3, v4, v5, i, size);
}

inline void or6_sve_strided(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, const uint64_t* v5, const uint64_t* v6, uint32_t* control) {
    uint32_t size = control[0]; uint32_t i = control[2]; uint32_t j = 2;
    while (i < size && j + 1 < size) {
        svbool_t pred = svwhilelt_b64(i, size);
        svuint64_t r = svorr_u64_x(pred, svld1_u64(pred, &v1[i]), svld1_u64(pred, &v2[i]));
        r = svorr_u64_x(pred, r, svld1_u64(pred, &v3[i]));
        r = svorr_u64_x(pred, r, svld1_u64(pred, &v4[i]));
        r = svorr_u64_x(pred, r, svld1_u64(pred, &v5[i]));
        r = svorr_u64_x(pred, r, svld1_u64(pred, &v6[i]));
        svst1_u64(pred, &out[i], r);
        if (sve_is_allones(r)) control[j+1]++;
        j++; i += control[j];
    }
    or6_scalar_tail(out, v1, v2, v3, v4, v5, v6, i, size);
}

inline void and4_neon_strided(uint64x2_t* out, const uint64x2_t* v1, const uint64x2_t* v2, const uint64x2_t* v3, const uint64x2_t* v4,
                              uint32_t* control) {
    uint32_t size = control[0];
    uint32_t i = control[2];
    uint32_t j = 2;
    while (i + 1 < size && j + 1 < size) {
        uint64x2_t res = vandq_u64(v1[i/2], v2[i/2]);
        res = vandq_u64(res, v3[i/2]);
        res = vandq_u64(res, v4[i/2]);
        out[i/2] = res;
        if (neon_is_zero(res)) control[j+1]++;
        j++; i += control[j];
    }
    and4_scalar_tail(reinterpret_cast<uint64_t*>(out),
                     reinterpret_cast<const uint64_t*>(v1),
                     reinterpret_cast<const uint64_t*>(v2),
                     reinterpret_cast<const uint64_t*>(v3),
                     reinterpret_cast<const uint64_t*>(v4), i, size);
}

inline void and5_neon_strided(uint64x2_t* out, const uint64x2_t* v1, const uint64x2_t* v2, const uint64x2_t* v3, const uint64x2_t* v4, const uint64x2_t* v5,
                              uint32_t* control) {
    uint32_t size = control[0];
    uint32_t i = control[2];
    uint32_t j = 2;
    while (i + 1 < size && j + 1 < size) {
        uint64x2_t res = vandq_u64(v1[i/2], v2[i/2]);
        res = vandq_u64(res, v3[i/2]);
        res = vandq_u64(res, v4[i/2]);
        res = vandq_u64(res, v5[i/2]);
        out[i/2] = res;
        if (neon_is_zero(res)) control[j+1]++;
        j++; i += control[j];
    }
    and5_scalar_tail(reinterpret_cast<uint64_t*>(out),
                     reinterpret_cast<const uint64_t*>(v1),
                     reinterpret_cast<const uint64_t*>(v2),
                     reinterpret_cast<const uint64_t*>(v3),
                     reinterpret_cast<const uint64_t*>(v4),
                     reinterpret_cast<const uint64_t*>(v5), i, size);
}

inline void and6_neon_strided(uint64x2_t* out, const uint64x2_t* v1, const uint64x2_t* v2, const uint64x2_t* v3, const uint64x2_t* v4, const uint64x2_t* v5, const uint64x2_t* v6,
                              uint32_t* control) {
    uint32_t size = control[0];
    uint32_t i = control[2];
    uint32_t j = 2;
    while (i + 1 < size && j + 1 < size) {
        uint64x2_t res = vandq_u64(v1[i/2], v2[i/2]);
        res = vandq_u64(res, v3[i/2]);
        res = vandq_u64(res, v4[i/2]);
        res = vandq_u64(res, v5[i/2]);
        res = vandq_u64(res, v6[i/2]);
        out[i/2] = res;
        if (neon_is_zero(res)) control[j+1]++;
        j++; i += control[j];
    }
    and6_scalar_tail(reinterpret_cast<uint64_t*>(out),
                     reinterpret_cast<const uint64_t*>(v1),
                     reinterpret_cast<const uint64_t*>(v2),
                     reinterpret_cast<const uint64_t*>(v3),
                     reinterpret_cast<const uint64_t*>(v4),
                     reinterpret_cast<const uint64_t*>(v5),
                     reinterpret_cast<const uint64_t*>(v6), i, size);
}

// ==== V3..V6 STRIDED OR IMPLEMENTATIONS ====

inline void or3_neon_strided(uint64x2_t* out, const uint64x2_t* v1, const uint64x2_t* v2, const uint64x2_t* v3,
                             uint32_t* control) {
    uint32_t size = control[0];
    uint32_t i = control[2];
    uint32_t j = 2;
    while (i + 1 < size && j + 1 < size) {
        uint64x2_t res = vorrq_u64(v1[i/2], v2[i/2]);
        res = vorrq_u64(res, v3[i/2]);
        out[i/2] = res;
        if (neon_is_allones(res)) control[j+1]++;
        j++; i += control[j];
    }
    or3_scalar_tail(reinterpret_cast<uint64_t*>(out),
                    reinterpret_cast<const uint64_t*>(v1),
                    reinterpret_cast<const uint64_t*>(v2),
                    reinterpret_cast<const uint64_t*>(v3), i, size);
}

inline void or4_neon_strided(uint64x2_t* out, const uint64x2_t* v1, const uint64x2_t* v2, const uint64x2_t* v3, const uint64x2_t* v4,
                             uint32_t* control) {
    uint32_t size = control[0];
    uint32_t i = control[2];
    uint32_t j = 2;
    while (i + 1 < size && j + 1 < size) {
        uint64x2_t res = vorrq_u64(v1[i/2], v2[i/2]);
        res = vorrq_u64(res, v3[i/2]);
        res = vorrq_u64(res, v4[i/2]);
        out[i/2] = res;
        if (neon_is_allones(res)) control[j+1]++;
        j++; i += control[j];
    }
    or4_scalar_tail(reinterpret_cast<uint64_t*>(out),
                    reinterpret_cast<const uint64_t*>(v1),
                    reinterpret_cast<const uint64_t*>(v2),
                    reinterpret_cast<const uint64_t*>(v3),
                    reinterpret_cast<const uint64_t*>(v4), i, size);
}

inline void or5_neon_strided(uint64x2_t* out, const uint64x2_t* v1, const uint64x2_t* v2, const uint64x2_t* v3, const uint64x2_t* v4, const uint64x2_t* v5,
                             uint32_t* control) {
    uint32_t size = control[0];
    uint32_t i = control[2];
    uint32_t j = 2;
    while (i + 1 < size && j + 1 < size) {
        uint64x2_t res = vorrq_u64(v1[i/2], v2[i/2]);
        res = vorrq_u64(res, v3[i/2]);
        res = vorrq_u64(res, v4[i/2]);
        res = vorrq_u64(res, v5[i/2]);
        out[i/2] = res;
        if (neon_is_allones(res)) control[j+1]++;
        j++; i += control[j];
    }
    or5_scalar_tail(reinterpret_cast<uint64_t*>(out),
                    reinterpret_cast<const uint64_t*>(v1),
                    reinterpret_cast<const uint64_t*>(v2),
                    reinterpret_cast<const uint64_t*>(v3),
                    reinterpret_cast<const uint64_t*>(v4),
                    reinterpret_cast<const uint64_t*>(v5), i, size);
}

inline void or6_neon_strided(uint64x2_t* out, const uint64x2_t* v1, const uint64x2_t* v2, const uint64x2_t* v3, const uint64x2_t* v4, const uint64x2_t* v5, const uint64x2_t* v6,
                             uint32_t* control) {
    uint32_t size = control[0];
    uint32_t i = control[2];
    uint32_t j = 2;
    while (i + 1 < size && j + 1 < size) {
        uint64x2_t res = vorrq_u64(v1[i/2], v2[i/2]);
        res = vorrq_u64(res, v3[i/2]);
        res = vorrq_u64(res, v4[i/2]);
        res = vorrq_u64(res, v5[i/2]);
        res = vorrq_u64(res, v6[i/2]);
        out[i/2] = res;
        if (neon_is_allones(res)) control[j+1]++;
        j++; i += control[j];
    }
    or6_scalar_tail(reinterpret_cast<uint64_t*>(out),
                    reinterpret_cast<const uint64_t*>(v1),
                    reinterpret_cast<const uint64_t*>(v2),
                    reinterpret_cast<const uint64_t*>(v3),
                    reinterpret_cast<const uint64_t*>(v4),
                    reinterpret_cast<const uint64_t*>(v5),
                    reinterpret_cast<const uint64_t*>(v6), i, size);
}

// Fallback scalar AND for remaining elements (no strides)
inline void and_scalar_tail(uint64_t* out, const uint64_t* v1, const uint64_t* v2,
                            uint32_t start, uint32_t size) {
    for (uint32_t i = start; i < size; ++i) {
        out[i] = v1[i] & v2[i];
    }
}

// Scalar tails for 3–6 operand AND/OR
inline void and3_scalar_tail(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3,
                             uint32_t start, uint32_t size) {
    for (uint32_t i = start; i < size; ++i) {
        out[i] = v1[i] & v2[i] & v3[i];
    }
}

inline void and4_scalar_tail(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4,
                             uint32_t start, uint32_t size) {
    for (uint32_t i = start; i < size; ++i) {
        out[i] = v1[i] & v2[i] & v3[i] & v4[i];
    }
}

inline void and5_scalar_tail(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, const uint64_t* v5,
                             uint32_t start, uint32_t size) {
    for (uint32_t i = start; i < size; ++i) {
        out[i] = v1[i] & v2[i] & v3[i] & v4[i] & v5[i];
    }
}

inline void and6_scalar_tail(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, const uint64_t* v5, const uint64_t* v6,
                             uint32_t start, uint32_t size) {
    for (uint32_t i = start; i < size; ++i) {
        out[i] = v1[i] & v2[i] & v3[i] & v4[i] & v5[i] & v6[i];
    }
}

inline void or3_scalar_tail(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3,
                            uint32_t start, uint32_t size) {
    for (uint32_t i = start; i < size; ++i) {
        out[i] = v1[i] | v2[i] | v3[i];
    }
}

inline void or4_scalar_tail(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4,
                            uint32_t start, uint32_t size) {
    for (uint32_t i = start; i < size; ++i) {
        out[i] = v1[i] | v2[i] | v3[i] | v4[i];
    }
}

inline void or5_scalar_tail(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, const uint64_t* v5,
                            uint32_t start, uint32_t size) {
    for (uint32_t i = start; i < size; ++i) {
        out[i] = v1[i] | v2[i] | v3[i] | v4[i] | v5[i];
    }
}

inline void or6_scalar_tail(uint64_t* out, const uint64_t* v1, const uint64_t* v2, const uint64_t* v3, const uint64_t* v4, const uint64_t* v5, const uint64_t* v6,
                            uint32_t start, uint32_t size) {
    for (uint32_t i = start; i < size; ++i) {
        out[i] = v1[i] | v2[i] | v3[i] | v4[i] | v5[i] | v6[i];
    }
}

// Main strided vectorized AND with NEON.
inline void and_neon_strided(uint64x2_t* out, const uint64x2_t* v1, const uint64x2_t* v2,
                             uint32_t* control) {
    uint32_t size = control[0];    // total number of 64-bit words
    uint32_t feature = control[1]; //cpu feature
    uint32_t i = control[2];       // starting word index
    uint32_t j = 2;                // stride index

    while (i + 1 < size && j + 1 < size) {
        // Vectorized AND: process two 64-bit words at once
        uint64x2_t result = vandq_u64(v1[i / 2], v2[i / 2]);
        out[i / 2] = result;

        // Dynamically widen stride if result becomes zero (all bits cleared)
        if (neon_is_zero(result)) {
            control[j + 1]++;
        }

        j++;
        i += control[j];  // move to next strided word index
    }

    // Scalar fallback for remaining words (no stride logic)
    and_scalar_tail(reinterpret_cast<uint64_t*>(out),
                    reinterpret_cast<const uint64_t*>(v1),
                    reinterpret_cast<const uint64_t*>(v2),
                    i, size);
}


// ===================== SVE STRIDED 2-operand =====================

inline void and_sve_strided(uint64_t* out, const uint64_t* v1, const uint64_t* v2, uint32_t* control) {
    uint32_t size = control[0];
    uint32_t i = control[2];
    uint32_t j = 2;

    while (i < size && j + 1 < size) {
        svbool_t pred = svwhilelt_b64(i, size);
        svuint64_t vec1 = svld1_u64(pred, &v1[i]);
        svuint64_t vec2 = svld1_u64(pred, &v2[i]);
        svuint64_t res  = svand_u64_x(pred, vec1, vec2);
        svst1_u64(pred, &out[i], res);

        if (sve_is_zero(res)) {
            control[j + 1]++;
        }

        j++;
        i += control[j];
    }

    // Scalar tail for any remaining words
    and_scalar_tail(out, v1, v2, i, size);
}

inline void or_sve_strided(uint64_t* out, const uint64_t* v1, const uint64_t* v2, uint32_t* control) {
    uint32_t size = control[0];
    uint32_t i = control[2];
    uint32_t j = 2;

    while (i < size && j + 1 < size) {
        svbool_t pred = svwhilelt_b64(i, size);
        svuint64_t vec1 = svld1_u64(pred, &v1[i]);
        svuint64_t vec2 = svld1_u64(pred, &v2[i]);
        svuint64_t res  = svorr_u64_x(pred, vec1, vec2);
        svst1_u64(pred, &out[i], res);

        if (sve_is_allones(res)) {
            control[j + 1]++;
        }

        j++;
        i += control[j];
    }

    or_scalar_tail(out, v1, v2, i, size);
}
