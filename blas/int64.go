package blas

func (o Int64s) add(input1, input2 []int64) {
	for i, v := range input1 {
		o[i] = v + input2[i]
	}
}

func (o Int64s) sub(input1, input2 []int64) {
	for i, v := range input1 {
		o[i] = v - input2[i]
	}
}

func (o Int64s) mul(input1, input2 []int64) {
	for i, v := range input1 {
		o[i] = v * input2[i]
	}
}
