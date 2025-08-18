#include <arm_neon.h>
#include <arm_sve.h>


const uint32_t SVE = 3;

// -----------------------------
// 64-bit span mask builder
// Caller sets mask[0] = N (number of 64-bit words) and ensures capacity >= 3 + 2*N
// Produces: stride[2] = num_strides; stride[3..] = {start0,len0,start1,len1,...} in 64-bit units.

void set_and_hot_blocks(const uint64_t *__restrict V,
                        uint32_t *__restrict mask) {
    const uint32_t N = mask[0];   // number of 64-bit words
    uint32_t *out = &mask[3];
    uint32_t num_spans = 0;

    uint32_t i = 0;
    while (i < N) {
        while (i < N && V[i] == 0) { ++i; }
        if (i == N) break;

        const uint32_t start = i;
        do { ++i; } while (i < N && V[i] != 0);

        out[2 * num_spans + 0] = start;      // start in 64-bit words
        out[2 * num_spans + 1] = i - start;  // length in 64-bit words
        ++num_spans;
    }
    mask[2] = num_spans;
}

void set_or_hot_blocks(const uint64_t *__restrict V,
                       uint32_t *__restrict mask) {
    const uint32_t N = mask[0];   // number of 64-bit words
    uint32_t *out = &mask[3];
    uint32_t num_spans = 0;

    const uint64_t ALL_ONES = ~UINT64_C(0);

    uint32_t i = 0;
    while (i < N) {
        // Skip words that ARE all ones
        while (i < N && V[i] == ALL_ONES) { ++i; }
        if (i == N) break;

        const uint32_t start = i;

        do { ++i; } while (i < N && V[i] != ALL_ONES);

        out[2 * num_spans + 0] = start;      // start index in 64-bit words
        out[2 * num_spans + 1] = i - start;  // length in 64-bit words
        ++num_spans;
    }

    mask[2] = num_spans;
}


// inline functions
// -----------------------------
// AND over 64-bit spans using NEON for the aligned middle and scalar head/tail.
// Spans in mask are in 64-bit words. Unaligned addresses are OK on AArch64.
// -----------------------------
inline void and_strided_neon(const uint64_t *__restrict A,
                             const uint64_t *__restrict B,
                             uint64_t *__restrict C,
                             const uint32_t *mask) {
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint32_t start = spans[2 * s + 0];   // in 64-bit words
        uint32_t len = spans[2 * s + 1];   // in 64-bit words
        if (len == 0) continue;

        // Head: handle odd starting index
        if (start & 1u) {
            C[start] = A[start] & B[start];
            ++start;
            --len;
            if (len == 0) continue;
        }

        // Middle: process pairs of 64-bit words as one 128-bit vector
        uint32_t nvec = len >> 1; // number of 128-bit vectors
        for (uint32_t v = 0; v < nvec; ++v) {
            uint64x2_t va = vld1q_u64(A + start + (v << 1));
            uint64x2_t vb = vld1q_u64(B + start + (v << 1));
            vst1q_u64(C + start + (v << 1), vandq_u64(va, vb));
        }

        // Tail: one leftover 64-bit word if len was odd
        if (len & 1u) {
            uint32_t k = start + (nvec << 1);
            C[k] = A[k] & B[k];
        }
    }
}

inline void or_strided_neon(const uint64_t *__restrict A,
                            const uint64_t *__restrict B,
                            uint64_t *__restrict C,
                            const uint32_t *mask) {
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint32_t start = spans[2 * s + 0];   // in 64-bit words
        uint32_t len = spans[2 * s + 1];   // in 64-bit words
        if (len == 0) continue;

        // Head: handle odd starting index
        if (start & 1u) {
            C[start] = A[start] | B[start];
            ++start;
            --len;
            if (len == 0) continue;
        }

        // Middle: process pairs of 64-bit words as one 128-bit vector
        uint32_t nvec = len >> 1; // number of 128-bit vectors
        for (uint32_t v = 0; v < nvec; ++v) {
            uint64x2_t va = vld1q_u64(A + start + (v << 1));
            uint64x2_t vb = vld1q_u64(B + start + (v << 1));
            vst1q_u64(C + start + (v << 1), vorrq_u64(va, vb));
        }

        // Tail: one leftover 64-bit word if len was odd
        if (len & 1u) {
            uint32_t k = start + (nvec << 1);
            C[k] = A[k] | B[k];
        }
    }
}

inline void and3_strided_neon(const uint64_t *__restrict A,
                              const uint64_t *__restrict B,
                              const uint64_t *__restrict D,  // third input
                              uint64_t *__restrict R,        // output
                              const uint32_t *mask) {
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint32_t start = spans[2 * s + 0];   // in 64-bit words
        uint32_t len = spans[2 * s + 1];   // in 64-bit words
        if (len == 0) continue;

        // Head: handle odd starting index
        if (start & 1u) {
            R[start] = A[start] & B[start] & D[start];
            ++start;
            --len;
            if (len == 0) continue;
        }

        // Middle: process pairs of 64-bit words as one 128-bit vector
        uint32_t nvec = len >> 1; // number of 128-bit vectors
        for (uint32_t v = 0; v < nvec; ++v) {
            const uint32_t k = start + (v << 1);
            uint64x2_t va = vld1q_u64(A + k);
            uint64x2_t vb = vld1q_u64(B + k);
            uint64x2_t vd = vld1q_u64(D + k);
            uint64x2_t t = vandq_u64(va, vb);
            uint64x2_t r = vandq_u64(t, vd);
            vst1q_u64(R + k, r);
        }

        // Tail: one leftover 64-bit word if len was odd
        if (len & 1u) {
            const uint32_t k = start + (nvec << 1);
            R[k] = A[k] & B[k] & D[k];
        }
    }
}

inline void or3_strided_neon(const uint64_t *__restrict A,
                             const uint64_t *__restrict B,
                             const uint64_t *__restrict D,  // third input
                             uint64_t *__restrict R,        // output
                             const uint32_t *mask) {
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint32_t start = spans[2 * s + 0];   // in 64-bit words
        uint32_t len = spans[2 * s + 1];   // in 64-bit words
        if (len == 0) continue;

        // Head: handle odd starting index
        if (start & 1u) {
            R[start] = A[start] | B[start] | D[start];
            ++start;
            --len;
            if (len == 0) continue;
        }

        // Middle: process pairs of 64-bit words as one 128-bit vector
        const uint32_t nvec = len >> 1; // number of 128-bit vectors
        for (uint32_t v = 0; v < nvec; ++v) {
            const uint32_t k = start + (v << 1);
            uint64x2_t va = vld1q_u64(A + k);
            uint64x2_t vb = vld1q_u64(B + k);
            uint64x2_t vd = vld1q_u64(D + k);
            uint64x2_t t = vorrq_u64(va, vb);
            uint64x2_t r = vorrq_u64(t, vd);
            vst1q_u64(R + k, r);
        }

        // Tail: one leftover 64-bit word if len was odd
        if (len & 1u) {
            const uint32_t k = start + (nvec << 1);
            R[k] = A[k] | B[k] | D[k];
        }
    }
}

inline void and4_strided_neon(const uint64_t *__restrict A,
                              const uint64_t *__restrict B,
                              const uint64_t *__restrict C,  // new third input
                              const uint64_t *__restrict D,  // fourth input
                              uint64_t *__restrict R,        // output
                              const uint32_t *mask) {
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint32_t start = spans[2 * s + 0]; // in 64-bit words
        uint32_t len   = spans[2 * s + 1]; // in 64-bit words
        if (len == 0) continue;

        // Head: handle odd starting index to keep 128-bit pairs aligned
        if (start & 1u) {
            R[start] = A[start] & B[start] & C[start] & D[start];
            ++start;
            if (--len == 0) continue;
        }

        // Middle: process pairs of 64-bit words as one 128-bit vector
        const uint32_t nvec = len >> 1; // number of 128-bit vectors
        for (uint32_t v = 0; v < nvec; ++v) {
            const uint32_t k = start + (v << 1);
            uint64x2_t va = vld1q_u64(A + k);
            uint64x2_t vb = vld1q_u64(B + k);
            uint64x2_t vc = vld1q_u64(C + k);
            uint64x2_t vd = vld1q_u64(D + k);

            uint64x2_t t  = vandq_u64(va, vb);
            t             = vandq_u64(t, vc);
            uint64x2_t r  = vandq_u64(t, vd);

            vst1q_u64(R + k, r);
        }

        // Tail: one leftover 64-bit word if len was odd
        if (len & 1u) {
            const uint32_t k = start + (nvec << 1);
            R[k] = A[k] & B[k] & C[k] & D[k];
        }
    }
}

inline void or4_strided_neon(const uint64_t *__restrict A,
                             const uint64_t *__restrict B,
                             const uint64_t *__restrict C,  // new third input
                             const uint64_t *__restrict D,  // fourth input
                             uint64_t *__restrict R,        // output
                             const uint32_t *mask) {
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint32_t start = spans[2 * s + 0]; // in 64-bit words
        uint32_t len   = spans[2 * s + 1]; // in 64-bit words
        if (len == 0) continue;

        // Head: handle odd starting index to keep 128-bit pairs aligned
        if (start & 1u) {
            R[start] = A[start] | B[start] | C[start] | D[start];
            ++start;
            if (--len == 0) continue;
        }

        // Middle: process pairs of 64-bit words as one 128-bit vector
        const uint32_t nvec = len >> 1; // number of 128-bit vectors
        for (uint32_t v = 0; v < nvec; ++v) {
            const uint32_t k = start + (v << 1);
            uint64x2_t va = vld1q_u64(A + k);
            uint64x2_t vb = vld1q_u64(B + k);
            uint64x2_t vc = vld1q_u64(C + k);
            uint64x2_t vd = vld1q_u64(D + k);

            uint64x2_t t = vorrq_u64(va, vb);
            t            = vorrq_u64(t, vc);
            uint64x2_t r = vorrq_u64(t, vd);

            vst1q_u64(R + k, r);
        }

        // Tail: one leftover 64-bit word if len was odd
        if (len & 1u) {
            const uint32_t k = start + (nvec << 1);
            R[k] = A[k] | B[k] | C[k] | D[k];
        }
    }
}

inline void and5_strided_neon(const uint64_t *__restrict A,
                              const uint64_t *__restrict B,
                              const uint64_t *__restrict C,
                              const uint64_t *__restrict D,
                              const uint64_t *__restrict E,
                              uint64_t *__restrict R,
                              const uint32_t *mask) {
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint32_t start = spans[2 * s + 0]; // in 64-bit words
        uint32_t len   = spans[2 * s + 1]; // in 64-bit words
        if (len == 0) continue;

        // Head
        if (start & 1u) {
            R[start] = A[start] & B[start] & C[start] & D[start] & E[start];
            ++start;
            if (--len == 0) continue;
        }

        // Middle
        const uint32_t nvec = len >> 1; // number of 128-bit vectors
        for (uint32_t v = 0; v < nvec; ++v) {
            const uint32_t k = start + (v << 1);
            uint64x2_t va = vld1q_u64(A + k);
            uint64x2_t vb = vld1q_u64(B + k);
            uint64x2_t vc = vld1q_u64(C + k);
            uint64x2_t vd = vld1q_u64(D + k);
            uint64x2_t ve = vld1q_u64(E + k);

            uint64x2_t t = vandq_u64(va, vb);
            t = vandq_u64(t, vc);
            t = vandq_u64(t, vd);
            uint64x2_t r = vandq_u64(t, ve);

            vst1q_u64(R + k, r);
        }

        // Tail
        if (len & 1u) {
            const uint32_t k = start + (nvec << 1);
            R[k] = A[k] & B[k] & C[k] & D[k] & E[k];
        }
    }
}

inline void or5_strided_neon(const uint64_t *__restrict A,
                             const uint64_t *__restrict B,
                             const uint64_t *__restrict C,
                             const uint64_t *__restrict D,
                             const uint64_t *__restrict E,
                             uint64_t *__restrict R,
                             const uint32_t *mask) {
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint32_t start = spans[2 * s + 0]; // in 64-bit words
        uint32_t len   = spans[2 * s + 1]; // in 64-bit words
        if (len == 0) continue;

        // Head
        if (start & 1u) {
            R[start] = A[start] | B[start] | C[start] | D[start] | E[start];
            ++start;
            if (--len == 0) continue;
        }

        // Middle
        const uint32_t nvec = len >> 1;
        for (uint32_t v = 0; v < nvec; ++v) {
            const uint32_t k = start + (v << 1);
            uint64x2_t va = vld1q_u64(A + k);
            uint64x2_t vb = vld1q_u64(B + k);
            uint64x2_t vc = vld1q_u64(C + k);
            uint64x2_t vd = vld1q_u64(D + k);
            uint64x2_t ve = vld1q_u64(E + k);

            uint64x2_t t = vorrq_u64(va, vb);
            t = vorrq_u64(t, vc);
            t = vorrq_u64(t, vd);
            uint64x2_t r = vorrq_u64(t, ve);

            vst1q_u64(R + k, r);
        }

        // Tail
        if (len & 1u) {
            const uint32_t k = start + (nvec << 1);
            R[k] = A[k] | B[k] | C[k] | D[k] | E[k];
        }
    }
}

inline void and6_strided_neon(const uint64_t *__restrict A,
                              const uint64_t *__restrict B,
                              const uint64_t *__restrict C,
                              const uint64_t *__restrict D,
                              const uint64_t *__restrict E,
                              const uint64_t *__restrict F,
                              uint64_t *__restrict R,
                              const uint32_t *mask) {
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint32_t start = spans[2 * s + 0]; // in 64-bit words
        uint32_t len   = spans[2 * s + 1]; // in 64-bit words
        if (len == 0) continue;

        // Head
        if (start & 1u) {
            R[start] = A[start] & B[start] & C[start] & D[start] & E[start] & F[start];
            ++start;
            if (--len == 0) continue;
        }

        // Middle
        const uint32_t nvec = len >> 1;
        for (uint32_t v = 0; v < nvec; ++v) {
            const uint32_t k = start + (v << 1);
            uint64x2_t va = vld1q_u64(A + k);
            uint64x2_t vb = vld1q_u64(B + k);
            uint64x2_t vc = vld1q_u64(C + k);
            uint64x2_t vd = vld1q_u64(D + k);
            uint64x2_t ve = vld1q_u64(E + k);
            uint64x2_t vf = vld1q_u64(F + k);

            uint64x2_t t = vandq_u64(va, vb);
            t = vandq_u64(t, vc);
            t = vandq_u64(t, vd);
            t = vandq_u64(t, ve);
            uint64x2_t r = vandq_u64(t, vf);

            vst1q_u64(R + k, r);
        }

        // Tail
        if (len & 1u) {
            const uint32_t k = start + (nvec << 1);
            R[k] = A[k] & B[k] & C[k] & D[k] & E[k] & F[k];
        }
    }
}


// OR inline
inline void or6_strided_neon(const uint64_t *__restrict A,
                             const uint64_t *__restrict B,
                             const uint64_t *__restrict C,
                             const uint64_t *__restrict D,
                             const uint64_t *__restrict E,
                             const uint64_t *__restrict F,
                             uint64_t *__restrict R,
                             const uint32_t *mask) {
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint32_t start = spans[2 * s + 0]; // in 64-bit words
        uint32_t len   = spans[2 * s + 1]; // in 64-bit words
        if (len == 0) continue;

        // Head
        if (start & 1u) {
            R[start] = A[start] | B[start] | C[start] | D[start] | E[start] | F[start];
            ++start;
            if (--len == 0) continue;
        }

        // Middle
        const uint32_t nvec = len >> 1;
        for (uint32_t v = 0; v < nvec; ++v) {
            const uint32_t k = start + (v << 1);
            uint64x2_t va = vld1q_u64(A + k);
            uint64x2_t vb = vld1q_u64(B + k);
            uint64x2_t vc = vld1q_u64(C + k);
            uint64x2_t vd = vld1q_u64(D + k);
            uint64x2_t ve = vld1q_u64(E + k);
            uint64x2_t vf = vld1q_u64(F + k);

            uint64x2_t t = vorrq_u64(va, vb);
            t = vorrq_u64(t, vc);
            t = vorrq_u64(t, vd);
            t = vorrq_u64(t, ve);
            uint64x2_t r = vorrq_u64(t, vf);

            vst1q_u64(R + k, r);
        }

        // Tail
        if (len & 1u) {
            const uint32_t k = start + (nvec << 1);
            R[k] = A[k] | B[k] | C[k] | D[k] | E[k] | F[k];
        }
    }
}

/// SVE implementation
inline void and_strided_sve(const uint64_t *__restrict A,
                             const uint64_t *__restrict B,
                             uint64_t *__restrict Out,
                             const uint32_t *mask)
{
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint64_t start = spans[2*s+0];
        uint64_t len   = spans[2*s+1];
        if (!len) continue;

        const uint64_t *a = A + start;
        const uint64_t *b = B + start;
        uint64_t *o       = Out + start;

        uint64_t i = 0;
        while (i < len) {
            svbool_t pg = svwhilelt_b64(i, len);
            svuint64_t va = svld1(pg, a + i);
            svuint64_t vb = svld1(pg, b + i);
            svuint64_t vr = svand_u64_x(pg, va, vb);
            svst1(pg, o + i, vr);
            i += svcntd();
        }
    }
}

inline void and3_strided_sve(const uint64_t *__restrict A,
                             const uint64_t *__restrict B,
                             const uint64_t *__restrict C,
                             uint64_t *__restrict Out,
                             const uint32_t *mask)
{
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint64_t start = spans[2*s+0];
        uint64_t len   = spans[2*s+1];
        if (!len) continue;

        const uint64_t *a = A + start, *b = B + start, *c = C + start;
        uint64_t *o = Out + start;

        uint64_t i = 0;
        while (i < len) {
            svbool_t pg = svwhilelt_b64(i, len);
            svuint64_t va = svld1(pg, a + i);
            svuint64_t vb = svld1(pg, b + i);
            svuint64_t vc = svld1(pg, c + i);
            svuint64_t vr = svand_u64_x(pg, svand_u64_x(pg, va, vb), vc);
            svst1(pg, o + i, vr);
            i += svcntd();
        }
    }
}

inline void and4_strided_sve(const uint64_t *__restrict A,
                             const uint64_t *__restrict B,
                             const uint64_t *__restrict C,
                             const uint64_t *__restrict D,
                             uint64_t *__restrict Out,
                             const uint32_t *mask)
{
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint64_t start = spans[2*s+0];
        uint64_t len   = spans[2*s+1];
        if (!len) continue;

        const uint64_t *a=A+start,*b=B+start,*c=C+start,*d=D+start;
        uint64_t *o = Out + start;

        uint64_t i=0;
        while (i < len) {
            svbool_t pg = svwhilelt_b64(i, len);
            svuint64_t va=svld1(pg,a+i), vb=svld1(pg,b+i);
            svuint64_t vc=svld1(pg,c+i), vd=svld1(pg,d+i);
            svuint64_t vr = svand_u64_x(pg, svand_u64_x(pg, va, vb),
                                            svand_u64_x(pg, vc, vd));
            svst1(pg, o+i, vr);
            i += svcntd();
        }
    }
}

inline void and5_strided_sve(const uint64_t *__restrict A,
                             const uint64_t *__restrict B,
                             const uint64_t *__restrict C,
                             const uint64_t *__restrict D,
                             const uint64_t *__restrict E,
                             uint64_t *__restrict Out,
                             const uint32_t *mask)
{
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s=0; s<num_spans; ++s) {
        uint64_t start=spans[2*s+0], len=spans[2*s+1];
        if (!len) continue;

        const uint64_t *a=A+start,*b=B+start,*c=C+start,*d=D+start,*e=E+start;
        uint64_t *o=Out+start;

        uint64_t i=0;
        while (i<len) {
            svbool_t pg=svwhilelt_b64(i,len);
            svuint64_t va=svld1(pg,a+i), vb=svld1(pg,b+i);
            svuint64_t vc=svld1(pg,c+i), vd=svld1(pg,d+i), ve=svld1(pg,e+i);
            svuint64_t vr = svand_u64_x(pg,
                                svand_u64_x(pg, svand_u64_x(pg, va, vb),
                                                svand_u64_x(pg, vc, vd)),
                                ve);
            svst1(pg,o+i,vr);
            i+=svcntd();
        }
    }
}

inline void and6_strided_sve(const uint64_t *__restrict A,
                             const uint64_t *__restrict B,
                             const uint64_t *__restrict C,
                             const uint64_t *__restrict D,
                             const uint64_t *__restrict E,
                             const uint64_t *__restrict F,
                             uint64_t *__restrict Out,
                             const uint32_t *mask)
{
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s=0; s<num_spans; ++s) {
        uint64_t start=spans[2*s+0], len=spans[2*s+1];
        if (!len) continue;

        const uint64_t *a=A+start,*b=B+start,*c=C+start,*d=D+start,*e=E+start,*f=F+start;
        uint64_t *o=Out+start;

        uint64_t i=0;
        while (i<len) {
            svbool_t pg=svwhilelt_b64(i,len);
            svuint64_t va=svld1(pg,a+i), vb=svld1(pg,b+i);
            svuint64_t vc=svld1(pg,c+i), vd=svld1(pg,d+i);
            svuint64_t ve=svld1(pg,e+i), vf=svld1(pg,f+i);
            // ((A&B)&(C&D))&(E&F)
            svuint64_t t0 = svand_u64_x(pg, va, vb);
            svuint64_t t1 = svand_u64_x(pg, vc, vd);
            svuint64_t t2 = svand_u64_x(pg, ve, vf);
            svuint64_t vr = svand_u64_x(pg, svand_u64_x(pg, t0, t1), t2);
            svst1(pg,o+i,vr);
            i+=svcntd();
        }
    }
}

// ---------- OR over 2..6 inputs ----------

inline void or_strided_sve(const uint64_t *__restrict A,
                            const uint64_t *__restrict B,
                            uint64_t *__restrict Out,
                            const uint32_t *mask)
{
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s=0; s<num_spans; ++s) {
        uint64_t start=spans[2*s+0], len=spans[2*s+1];
        if (!len) continue;

        const uint64_t *a=A+start,*b=B+start; uint64_t *o=Out+start;
        uint64_t i=0;
        while (i<len) {
            svbool_t pg=svwhilelt_b64(i,len);
            svuint64_t va=svld1(pg,a+i), vb=svld1(pg,b+i);
            svuint64_t vr=svorr_u64_x(pg,va,vb);
            svst1(pg,o+i,vr);
            i+=svcntd();
        }
    }
}

inline void or3_strided_sve(const uint64_t *__restrict A,
                            const uint64_t *__restrict B,
                            const uint64_t *__restrict C,
                            uint64_t *__restrict Out,
                            const uint32_t *mask)
{
    const uint32_t num_spans=mask[2];
    const uint32_t *spans=&mask[3];

    for (uint32_t s=0; s<num_spans; ++s) {
        uint64_t start=spans[2*s+0], len=spans[2*s+1];
        if (!len) continue;

        const uint64_t *a=A+start,*b=B+start,*c=C+start; uint64_t *o=Out+start;
        uint64_t i=0;
        while (i<len) {
            svbool_t pg=svwhilelt_b64(i,len);
            svuint64_t va=svld1(pg,a+i), vb=svld1(pg,b+i), vc=svld1(pg,c+i);
            svuint64_t vr=svorr_u64_x(pg, svorr_u64_x(pg,va,vb), vc);
            svst1(pg,o+i,vr);
            i+=svcntd();
        }
    }
}

inline void or4_strided_sve(const uint64_t *__restrict A,
                            const uint64_t *__restrict B,
                            const uint64_t *__restrict C,
                            const uint64_t *__restrict D,
                            uint64_t *__restrict Out,
                            const uint32_t *mask)
{
    const uint32_t num_spans=mask[2];
    const uint32_t *spans=&mask[3];

    for (uint32_t s=0; s<num_spans; ++s) {
        uint64_t start=spans[2*s+0], len=spans[2*s+1];
        if (!len) continue;

        const uint64_t *a=A+start,*b=B+start,*c=C+start,*d=D+start; uint64_t *o=Out+start;
        uint64_t i=0;
        while (i<len) {
            svbool_t pg=svwhilelt_b64(i,len);
            svuint64_t va=svld1(pg,a+i), vb=svld1(pg,b+i);
            svuint64_t vc=svld1(pg,c+i), vd=svld1(pg,d+i);
            svuint64_t vr=svorr_u64_x(pg, svorr_u64_x(pg,va,vb),
                                          svorr_u64_x(pg,vc,vd));
            svst1(pg,o+i,vr);
            i+=svcntd();
        }
    }
}

inline void or5_strided_sve(const uint64_t *__restrict A,
                            const uint64_t *__restrict B,
                            const uint64_t *__restrict C,
                            const uint64_t *__restrict D,
                            const uint64_t *__restrict E,
                            uint64_t *__restrict Out,
                            const uint32_t *mask)
{
    const uint32_t num_spans=mask[2];
    const uint32_t *spans=&mask[3];

    for (uint32_t s=0; s<num_spans; ++s) {
        uint64_t start=spans[2*s+0], len=spans[2*s+1];
        if (!len) continue;

        const uint64_t *a=A+start,*b=B+start,*c=C+start,*d=D+start,*e=E+start; uint64_t *o=Out+start;
        uint64_t i=0;
        while (i<len) {
            svbool_t pg=svwhilelt_b64(i,len);
            svuint64_t va=svld1(pg,a+i), vb=svld1(pg,b+i);
            svuint64_t vc=svld1(pg,c+i), vd=svld1(pg,d+i), ve=svld1(pg,e+i);
            svuint64_t vr=svorr_u64_x(pg,
                                svorr_u64_x(pg, svorr_u64_x(pg,va,vb),
                                                svorr_u64_x(pg,vc,vd)),
                                ve);
            svst1(pg,o+i,vr);
            i+=svcntd();
        }
    }
}

inline void or6_strided_sve(const uint64_t *__restrict A,
                            const uint64_t *__restrict B,
                            const uint64_t *__restrict C,
                            const uint64_t *__restrict D,
                            const uint64_t *__restrict E,
                            const uint64_t *__restrict F,
                            uint64_t *__restrict Out,
                            const uint32_t *mask)
{
    const uint32_t num_spans=mask[2];
    const uint32_t *spans=&mask[3];

    for (uint32_t s=0; s<num_spans; ++s) {
        uint64_t start=spans[2*s+0], len=spans[2*s+1];
        if (!len) continue;

        const uint64_t *a=A+start,*b=B+start,*c=C+start,*d=D+start,*e=E+start,*f=F+start;
        uint64_t *o=Out+start;

        uint64_t i=0;
        while (i<len) {
            svbool_t pg=svwhilelt_b64(i,len);
            svuint64_t va=svld1(pg,a+i), vb=svld1(pg,b+i);
            svuint64_t vc=svld1(pg,c+i), vd=svld1(pg,d+i);
            svuint64_t ve=svld1(pg,e+i), vf=svld1(pg,f+i);
            // ((A|B)|(C|D))|(E|F)
            svuint64_t t0 = svorr_u64_x(pg, va, vb);
            svuint64_t t1 = svorr_u64_x(pg, vc, vd);
            svuint64_t t2 = svorr_u64_x(pg, ve, vf);
            svuint64_t vr = svorr_u64_x(pg, svorr_u64_x(pg, t0, t1), t2);
            svst1(pg,o+i,vr);
            i+=svcntd();
        }
    }
}


/// Exposed functions:

/// AND
void and_strided(const uint64_t *__restrict A,
                 const uint64_t *__restrict B,
                 uint64_t *__restrict C,
                 const uint32_t *mask) {
    if (mask[1] == SVE) {
        and_strided_sve(A, B, C, mask);
    } else {
        and_strided_neon(A, B, C, mask);
    }
}

void and3_strided(const uint64_t *__restrict A,
                  const uint64_t *__restrict B,
                  const uint64_t *__restrict D,  // third input
                  uint64_t *__restrict R,        // output
                  const uint32_t *mask) {
    if (mask[1] == SVE) {
        and3_strided_sve(A, B, D, R, mask);
    } else {
        and3_strided_neon(A, B, D, R, mask);
    }
}

void and4_strided(const uint64_t *__restrict A,
                  const uint64_t *__restrict B,
                  const uint64_t *__restrict C,
                  const uint64_t *__restrict D,
                  uint64_t *__restrict R,        // output
                  const uint32_t *mask) {
    if (mask[1] == SVE) {
        and4_strided_sve(A, B, C, D, R, mask);
    } else {
        and4_strided_neon(A, B, C, D, R, mask);
    }
}

void and5_strided(const uint64_t *__restrict A,
                  const uint64_t *__restrict B,
                  const uint64_t *__restrict C,
                  const uint64_t *__restrict D,
                  const uint64_t *__restrict E,
                  uint64_t *__restrict R,        // output
                  const uint32_t *mask) {
    and5_strided_neon(A, B, C, D, E, R, mask);
    if (mask[1] == SVE) {
        and5_strided_sve(A, B, C, D, E, R, mask);
    } else {
        and5_strided_neon(A, B, C, D, E, R, mask);
    }
}

void and6_strided(const uint64_t *__restrict A,
                  const uint64_t *__restrict B,
                  const uint64_t *__restrict C,
                  const uint64_t *__restrict D,
                  const uint64_t *__restrict E,
                  const uint64_t *__restrict F,
                  uint64_t *__restrict R,        // output
                  const uint32_t *mask) {
    if (mask[1] == SVE) {
       and6_strided_neon(A, B, C, D, E, F, R, mask);
    } else {
       and6_strided_neon(A, B, C, D, E, F, R, mask);
    }
}


/// OR
void or_strided(const uint64_t *__restrict A,
                 const uint64_t *__restrict B,
                 uint64_t *__restrict C,
                 const uint32_t *mask) {
    if (mask[1] == SVE) {
        or_strided_sve(A, B, C, mask);
    } else {
        or_strided_neon(A, B, C, mask);
    }
}

void or3_strided(const uint64_t *__restrict A,
                  const uint64_t *__restrict B,
                  const uint64_t *__restrict D,  // third input
                  uint64_t *__restrict R,        // output
                  const uint32_t *mask) {
    if (mask[1] == SVE) {
        or3_strided_sve(A, B, D, R, mask);
    } else {
        or3_strided_neon(A, B, D, R, mask);
    }
}

void or4_strided(const uint64_t *__restrict A,
                  const uint64_t *__restrict B,
                  const uint64_t *__restrict C,
                  const uint64_t *__restrict D,
                  uint64_t *__restrict R,        // output
                  const uint32_t *mask) {
    if (mask[1] == SVE) {
        or4_strided_sve(A, B, C, D, R, mask);
    } else {
        or4_strided_neon(A, B, C, D, R, mask);
    }
}

void or5_strided(const uint64_t *__restrict A,
                  const uint64_t *__restrict B,
                  const uint64_t *__restrict C,
                  const uint64_t *__restrict D,
                  const uint64_t *__restrict E,
                  uint64_t *__restrict R,        // output
                  const uint32_t *mask) {
    if (mask[1] == SVE) {
        or5_strided_sve(A, B, C, D, E, R, mask);
    } else {
        or5_strided_neon(A, B, C, D, E, R, mask);
    }
}

void or6_strided(const uint64_t *__restrict A,
                  const uint64_t *__restrict B,
                  const uint64_t *__restrict C,
                  const uint64_t *__restrict D,
                  const uint64_t *__restrict E,
                  const uint64_t *__restrict F,
                  uint64_t *__restrict R,        // output
                  const uint32_t *mask) {
    if (mask[1] == SVE) {
       or6_strided_sve(A, B, C, D, E, F, R, mask);
    } else {
       or6_strided_neon(A, B, C, D, E, F, R, mask);
    }
}


