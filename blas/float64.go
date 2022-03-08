package blas

func (o Float64s) addFloat64(input1, input2 []float64) {
	for i, v := range input1 {
		o[i] = v + input2[i]
	}
}

func (o Float64s) subFloat64(input1, input2 []float64) {
	for i, v := range input1 {
		o[i] = v - input2[i]
	}
}

func (o Float64s) mulFloat64(input1, input2 []float64) {
	for i, v := range input1 {
		o[i] = v * input2[i]
	}
}

func (o Float64s) divFloat64(input1, input2 []float64) {
	for i, v := range input1 {
		o[i] = v / input2[i]
	}
}
