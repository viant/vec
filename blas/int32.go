package blas

func (o Int32s) incInt32(constant int32) {
	for i, v := range o {
		o[i] = v + constant
	}
}

func (o Int32s) addInt32(input1, input2 []int32) {
	for i, v := range input1 {
		o[i] = v + input2[i]
	}
}

func (o Int32s) subInt32(input1, input2 []int32) {
	for i, v := range input1 {
		o[i] = v - input2[i]
	}
}

func (o Int32s) mulInt32(input1, input2 []int32) {
	for i, v := range input1 {
		o[i] = v * input2[i]
	}
}
