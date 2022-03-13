package blas

func hsumInt64(input []int64) int64 {
	sum := int64(0)
	for _, v := range input {
		sum += v
	}
	return sum
}

func hmaxInt64(input []int64) int64 {
	max := input[0]
	for _, v := range input[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func hminInt64(input []int64) int64 {
	min := input[0]
	for _, v := range input[1:] {
		if v < min {
			min = v
		}
	}
	return min
}
