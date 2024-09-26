package blas

func (o Int32s) inc(constant int32) {
	for i, v := range o {
		o[i] = v + constant
	}
}

func (o Int32s) add(input1, input2 []int32) {
	for i, v := range input1 {
		o[i] = v + input2[i]
	}
}

func (o Int32s) sub(input1, input2 []int32) {
	for i, v := range input1 {
		o[i] = v - input2[i]
	}
}

func (o Int32s) mul(input1, input2 []int32) {
	for i, v := range input1 {
		o[i] = v * input2[i]
	}
}
