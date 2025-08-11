#include <arm_sve.h>

void find_nonzero_64bit_words_sve(const uint64_t* V, uint64_t* mask_buf) {
    const uint64_t N = mask_buf[0];
    uint64_t* out = &mask_buf[3];
    uint64_t count = 0;

    uint64_t i = 0;
    while (i < N) {
        svbool_t pg = svwhilelt_b64(i, N);
        svuint64_t vec = svld1(pg, &V[i]);

        uint64_t tmp[svcntd()];
        svst1(pg, tmp, vec);

        for (uint64_t lane = 0; lane < svcntd(); ++lane) {
            if ((i + lane) >= N) break;
            if (tmp[lane] != 0) {
                out[count++] = i + lane;
            }
        }

        i += svcntd();
    }

    mask_buf[2] = count;
}

void and_masked_64bit_sve(const uint64_t* A, const uint64_t* B, uint64_t* C, const uint64_t* mask_buf) {
    const uint64_t count = mask_buf[2];
    const uint64_t* indices = &mask_buf[3];

    uint64_t i = 0;
    while (i < count) {
        svbool_t pg = svwhilelt_b64(i, count);
        svuint64_t idx = svld1(pg, &indices[i]);

        svuint64_t a = svld1_gather_index(pg, A, idx);
        svuint64_t b = svld1_gather_index(pg, B, idx);
        svuint64_t result = svand_z(pg, a, b);

        svst1_scatter_index(pg, C, idx, result);
        i += svcntd();
    }
}
