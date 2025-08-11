#include <arm_neon.h>

void find_nonzero_words(const uint64_t* V, uint32_t* mask_buf) {
    const uint32_t N = mask_buf[0];
    const uint32_t cpu_type = mask_buf[1];
    uint32_t* out = &mask_buf[3];

    uint32_t count = 0;
    uint32_t i = 0;

    for (; i + 1 < N; i += 2) {
        uint64x2_t v = vld1q_u64(&V[i]);
        if (vgetq_lane_u64(v, 0) != 0) out[count++] = i;
        if (vgetq_lane_u64(v, 1) != 0) out[count++] = i + 1;
    }

    if (i < N) {
        if (V[i] != 0) out[count++] = i;
    }

    mask_buf[2] = count;
}

void and_masked_naive(const uint64_t* A, const uint64_t* B, uint64_t* C, const uint32_t* mask_buf) {
    const uint32_t N = mask_buf[0];
    const uint32_t count = mask_buf[2];
    const uint32_t* indices = &mask_buf[3];

    for (uint32_t i = 0; i < count; ++i) {
        uint32_t idx = indices[i];
        if (idx >= N) continue;
        C[idx] = A[idx] & B[idx];
    }
}


void and_masked_neon(const uint64_t* A, const uint64_t* B, uint64_t* C, const uint32_t* mask_buf) {
    const uint32_t N = mask_buf[0];
    const uint32_t count = mask_buf[2];
    const uint32_t* indices = &mask_buf[3];

    uint32_t i = 0;

    // Process 2 indices at a time
    for (; i + 1 < count; i += 2) {
        uint32_t idx0 = indices[i];
        uint32_t idx1 = indices[i + 1];

        if (idx0 < N) {
            uint64x1_t a0 = vld1_u64(&A[idx0]);
            uint64x1_t b0 = vld1_u64(&B[idx0]);
            uint64x1_t c0 = vand_u64(a0, b0);
            vst1_u64(&C[idx0], c0);
        }

        if (idx1 < N) {
            uint64x1_t a1 = vld1_u64(&A[idx1]);
            uint64x1_t b1 = vld1_u64(&B[idx1]);
            uint64x1_t c1 = vand_u64(a1, b1);
            vst1_u64(&C[idx1], c1);
        }
    }

    // Handle leftover
    for (; i < count; ++i) {
        uint32_t idx = indices[i];
        if (idx >= N) continue;
        C[idx] = A[idx] & B[idx];
    }
}

inline void and_scalar_tail(uint64_t* out, const uint64_t* v1, const uint64_t* v2,
                            uint32_t start, uint32_t size) {
    #pragma clang loop vectorize(disable)
    for (uint32_t i = start; i < size; ++i) {
        out[i] = v1[i] & v2[i];
    }
}

void and_neon_strided(uint64x2_t* out, const uint64x2_t* v1, const uint64x2_t* v2,
                             uint32_t* control) {
    uint32_t size = control[0];    // total number of 64-bit words
    uint32_t feature = control[1]; //cpu feature
    uint32_t i = control[2];       // starting word index
    uint32_t j = 2;                // stride index

    while (i + 1 < size && j + 1 < size) {
        // Vectorized AND: process two 64-bit words at once
        uint64x2_t result = vandq_u64(v1[i / 2], v2[i / 2]);
        out[i / 2] = result;

        j++;
        i += control[j];  // move to next strided word index
    }

    // Scalar fallback for remaining words (no stride logic)
    and_scalar_tail(reinterpret_cast<uint64_t*>(out),
                    reinterpret_cast<const uint64_t*>(v1),
                    reinterpret_cast<const uint64_t*>(v2),
                    i, size);
}

