package blas

func hsumInt32(input []int32) int32 {
	sum := int32(0)
	for _, v := range input {
		sum += v
	}
	return sum
}

func hmaxInt32(input []int32) int32 {
	max := input[0]
	for _, v := range input[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func hminInt32(input []int32) int32 {
	min := input[0]
	for _, v := range input[1:] {
		if v < min {
			min = v
		}
	}
	return min
}
