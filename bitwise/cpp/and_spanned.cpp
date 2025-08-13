#include <arm_neon.h>

static inline int any_nonzero_u128(uint64x2_t v) {
    // AArch64: horizontal max of bytes; non-zero if any byte non-zero
    return vmaxvq_u8(vreinterpretq_u8_u64(v)) != 0;
}

// -----------------------------
// 1) Build span mask for a single vector V (N 128-bit words)
//    Caller must set mask[0] = N and ensure capacity >= 3 + 2*N
// -----------------------------
void find_spans_neon(const uint64x2_t* V, uint32_t* mask) {
    const uint32_t N = mask[0];
    uint32_t* out = &mask[3];
    uint32_t num_spans = 0;

    uint32_t i = 0;
    while (i < N) {
        // Skip zeros
        while (i < N) {
            uint64x2_t v = vld1q_u64((const uint64_t*)&V[i]);
            if (any_nonzero_u128(v)) break;
            ++i;
        }
        if (i == N) break;

        const uint32_t start = i;
        // Extend run
        do {
            ++i;
        } while (i < N && any_nonzero_u128(vld1q_u64((const uint64_t*)&V[i])));

        out[2*num_spans + 0] = start;
        out[2*num_spans + 1] = i - start;
        ++num_spans;
    }
    mask[2] = num_spans;
}

// -----------------------------
// 2) AND using only one mask (no unrolling)
//    Computes C = A & B for indices covered by maskA's spans.
//    Regions outside those spans are left untouched in C.
// -----------------------------
void and_span_neon(const uint64x2_t* A,
                         const uint64x2_t* B,
                         uint64x2_t* C,
                         const uint32_t* maskA) {
    const uint32_t num_spans_A = maskA[2];
    const uint32_t* spansA = &maskA[3];

    for (uint32_t s = 0; s < num_spans_A; ++s) {
        const uint32_t start = spansA[2*s + 0];
        const uint32_t len   = spansA[2*s + 1];
        const uint32_t end   = start + len;

        for (uint32_t i = start; i < end; ++i) {
            uint64x2_t a = vld1q_u64((const uint64_t*)&A[i]);
            uint64x2_t b = vld1q_u64((const uint64_t*)&B[i]);
            vst1q_u64((uint64_t*)&C[i], vandq_u64(a, b));
        }
    }
}

