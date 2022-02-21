package pack

func (o Int32s) deltas(in []int32) {
	curValue := in[0]
	o[0] = curValue
	for i, v := range in[1:] {
		o[i+1] = v - curValue
		curValue = v
	}
}

func (o Int32s) recoverDeltas(in []int32) {
	curValue := in[0]
	o[0] = curValue
	for i, v := range in[1:] {
		o[i+1] = v + curValue
		curValue = o[i+1]
	}
}
