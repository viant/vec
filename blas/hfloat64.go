package blas

func hsumFloat64(input []float64) float64 {
	sum := float64(0)
	for _, v := range input {
		sum += v
	}
	return sum
}

func hmaxFloat64(input []float64) float64 {
	max := input[0]
	for _, v := range input[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func hminFloat64(input []float64) float64 {
	min := input[0]
	for _, v := range input[1:] {
		if v < min {
			min = v
		}
	}
	return min
}
