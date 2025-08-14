#include <arm_neon.h>
#include <arm_sve.h>


const uint32_t SVE = 3;

// -----------------------------
// 64-bit span mask builder
// Caller sets mask[0] = N (number of 64-bit words) and ensures capacity >= 3 + 2*N
// Produces: mask[2] = num_spans; mask[3..] = {start0,len0,start1,len1,...} in 64-bit units.

void set_and_hot_blocks(const uint64_t * __restrict V,
                         uint32_t * __restrict mask)
{
    const uint32_t N = mask[0];   // number of 64-bit words
    uint32_t *out = &mask[3];
    uint32_t num_spans = 0;

    uint32_t i = 0;
    while (i < N) {
        while (i < N && V[i] == 0) { ++i; }
        if (i == N) break;

        const uint32_t start = i;
        do { ++i; } while (i < N && V[i] != 0);

        out[2*num_spans + 0] = start;      // start in 64-bit words
        out[2*num_spans + 1] = i - start;  // length in 64-bit words
        ++num_spans;
    }
    mask[2] = num_spans;
}

// -----------------------------
// AND over 64-bit spans using NEON for the aligned middle and scalar head/tail.
// Spans in mask are in 64-bit words. Unaligned addresses are OK on AArch64.
// -----------------------------
inline void and_strided_neon(const uint64_t *__restrict A,
                             const uint64_t *__restrict B,
                             uint64_t *__restrict C,
                             const uint32_t *mask)
{
    const uint32_t num_spans = mask[2];
    const uint32_t *spans = &mask[3];

    for (uint32_t s = 0; s < num_spans; ++s) {
        uint32_t start = spans[2*s + 0];   // in 64-bit words
        uint32_t len   = spans[2*s + 1];   // in 64-bit words
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

void and_strided(const uint64_t *__restrict A,
                                             const uint64_t *__restrict B,
                                             uint64_t *__restrict C,
                                             const uint32_t *mask)
{
    if (mask[1] == SVE) {
        and_strided_neon(A, B, C, mask);
    } else {
        and_strided_neon(A, B, C, mask);
    }
}
