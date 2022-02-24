package blas

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
