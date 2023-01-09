func bitwiseComplement(N int) int {
	x := N
	x |= 1
	for k := x; k > 0; k /= 2 {
		x |= k
	}
	return x ^ N
}