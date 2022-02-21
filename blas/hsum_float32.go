package blas

func hsumFloat32(input []float32) float32 {
	sum := float32(0)
	for _, v := range input {
		sum += v
	}
	return sum
}
