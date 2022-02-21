package lut

func lookup(input, table []byte, output *[]byte) {
	for i, e := range input {
		(*output)[i] = table[e]
	}
	return
}

func lookupSubrange(input, range_upper, table []byte, output *[]byte) {
	for i := 0; i < len(input); i++ {
		n := 0
		for _, e := range range_upper {
			if input[i] > e {
				n += 1
			}
		}
		(*output)[i] = table[n]
	}
}
