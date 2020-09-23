package puzzle

func NewValueSet(valueCount uint) (vs []bool) {
	vs = make([]bool, valueCount)
	for i := range vs {
		vs[i] = true
	}

	return
}
