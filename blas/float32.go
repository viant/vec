package blas

func (o Float32s) addFloat32(input1, input2 []float32) {
	for i, v := range input1 {
		o[i] = v + input2[i]
	}
}

func (o Float32s) subFloat32(input1, input2 []float32) {
	for i, v := range input1 {
		o[i] = v - input2[i]
	}
}

func (o Float32s) mulFloat32(input1, input2 []float32) {
	for i, v := range input1 {
		o[i] = v * input2[i]
	}
}

func (o Float32s) divFloat32(input1, input2 []float32) {
	for i, v := range input1 {
		o[i] = v / input2[i]
	}
}
