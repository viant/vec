package blas

func hsumFloat32(input []float32) float32 {
	sum := float32(0)
	for _, v := range input {
		sum += v
	}
	return sum
}

func hmaxFloat32(input []float32) float32 {
	max := input[0]
	for _, v := range input[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func hminFloat32(input []float32) float32 {
	min := input[0]
	for _, v := range input[1:] {
		if v < min {
			min = v
		}
	}
	return min
}
