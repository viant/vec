package blas

func (o Int32s) addScalar(input []int32, value int32) {
	for i, v := range input {
		o[i] = v + value
	}
}
