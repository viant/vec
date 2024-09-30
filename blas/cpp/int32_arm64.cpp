#include <arm_neon.h>
#include <arm_sve.h>
#include <stdint.h>

void inc_int32(int32x4_t* vec, int32_t constant,  uint64_t size) {
    int chunks = size / 4;
    int32x4_t constant_vec = vdupq_n_s32(constant);

    for (int i = 0; i < chunks; i++) {
        vec[i] = vaddq_s32(vec[i], constant_vec);
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        ((int32_t *) vec)[i] = ((int32_t *) vec)[i] + constant;
    }
}

void add_int32(int32x4_t *input1, int32x4_t *input2, int32x4_t *output, uint64_t size) {
    int chunks = size / 4;

    for (int i = 0; i < chunks; i++) {
        output[i] = vaddq_s32(input1[i], input2[i]);
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        ((int32_t *) output)[i] = ((int32_t *) input1)[i] + ((int32_t *) input2)[i];
    }
}

void sub_int32(int32x4_t *input1, int32x4_t *input2, int32x4_t *output, uint64_t size) {
    int chunks = size / 4;

    for (int i = 0; i < chunks; i++) {
        output[i] = vsubq_s32(input1[i], input2[i]);
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        ((int32_t *) output)[i] = ((int32_t *) input1)[i] - ((int32_t *) input2)[i];
    }
}

void scalar_mul_int32(int32x4_t* vec, int32_t constant,  uint64_t size) {
    int chunks = size / 4;
    int32x4_t constant_vec = vdupq_n_s32(constant);

    for (int i = 0; i < chunks; i++) {
        vec[i] = vmulq_s32(vec[i], constant_vec);
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        ((int32_t *) vec)[i] = ((int32_t *) vec)[i] * constant;
    }
}

void mul_int32(int32x4_t *input1, int32x4_t *input2, int32x4_t *output, uint64_t size) {
    int chunks = size / 4;

    for (int i = 0; i < chunks; i++) {
        output[i] = vmulq_s32(input1[i], input2[i]);
    }
    // Leftovers
    for (int i = chunks * 4; i < size; i++) {
        ((int32_t *) output)[i] = ((int32_t *) input1)[i] * ((int32_t *) input2)[i];
    }
}


void inc_int32_sve(int32_t* vec, int32_t constant,  int size) {
    int i = 0;
    svint32_t constant_vec = svdup_s32(constant);

    while (i < size) {
        svbool_t pg = svwhilelt_b32(i, size);
        svst1(pg, &vec[i], svadd_s32_m(pg, svld1(pg, &vec[i]), constant_vec));
        i += svcntw();
    }
}

void scalar_mul_int32_sve(int32_t* vec, int32_t constant, int size) {
    int i = 0;
    svint32_t constant_vec = svdup_s32(constant);

    while (i < size) {
        svbool_t pg = svwhilelt_b32(i, size);
        svst1(pg, &vec[i], svmul_s32_m(pg, svld1(pg, &vec[i]), constant_vec));
        i += svcntw();
    }
}
