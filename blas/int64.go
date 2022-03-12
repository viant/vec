package blas

func (o Int64s) addInt64(input1, input2 []int64) {
	for i, v := range input1 {
		o[i] = v + input2[i]
	}
}

func (o Int64s) subInt64(input1, input2 []int64) {
	for i, v := range input1 {
		o[i] = v - input2[i]
	}
}

func (o Int64s) mulInt64(input1, input2 []int64) {
	for i, v := range input1 {
		o[i] = v * input2[i]
	}
}
