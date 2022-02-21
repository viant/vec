#include <arm_sve.h>

void lookup_sve_vla(uint8_t *input, uint8_t *output, uint8_t *table, uint64_t size) {
    uint64_t i = 0;
    svuint8_t vInput, vOutput, vLut;
    svbool_t pred = svwhilelt_b8(i, size);

    while (svptest_first(svptrue_b8(), pred)) {
        vInput = svld1(pred, &input[i]);
        uint64_t j = 0;
        svbool_t pred_lut = svwhilelt_b8(j, (uint64_t) 256);
        vOutput = svdup_u8(0);
        while (svptest_first(svptrue_b8(), pred_lut)) {
            vLut = svld1(pred_lut, &table[j]);
            vOutput = svorr_u8_x(pred, vOutput, svtbl(vLut, sveor_n_u8_x(pred, vInput, (uint8_t) j)));
            j += svcntb();
            pred_lut = svwhilelt_b8(j, (uint64_t) 256);
        }
        svst1(pred, &output[i], vOutput);
        i += svcntb();
        pred = svwhilelt_b8(i, size);
    }
}

void lookup_sve_2048(uint8_t *input, uint8_t *output, uint8_t *table, uint64_t size) {
    uint64_t i = 0;
    svuint8_t vInput, vOutput, vLut;
    svbool_t pred = svwhilelt_b8(i, size);
    vLut = svld1(svptrue_b8(), &table[i]);

    while (svptest_first(svptrue_b8(), pred)) {
        vInput = svld1(pred, &input[i]);
        vOutput = svtbl(vLut, vInput);
        svst1(pred, &output[i], vOutput);
        i += svcntb();
        pred = svwhilelt_b8(i, size);
    }
}
