package puzzle

func NewValueSet(valueCount uint8) (vs []bool) {
	vs = make([]bool, valueCount)
	for i := range vs {
		vs[i] = true
	}

	return
}
