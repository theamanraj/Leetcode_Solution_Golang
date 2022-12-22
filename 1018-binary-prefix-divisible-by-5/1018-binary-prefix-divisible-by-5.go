func prefixesDivBy5(A []int) []bool {
	n := len(A)
	r, acc := make([]bool, n), 0
	for i, v := range A {
		acc = acc<<1 + v
		r[i] = (acc%5 == 0)
		acc %= 10
	}
	return r
}
