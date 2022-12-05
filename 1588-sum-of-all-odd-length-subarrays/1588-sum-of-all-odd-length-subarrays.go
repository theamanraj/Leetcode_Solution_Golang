func sumOddLengthSubarrays(a []int) int {
	s := 0
	for k := 1; k <= len(a); k += 2 {
		for i := 0; (i + k) <= len(a); i++ {
			for _, v := range a[i : i+k] {
				s += v
			}
		}
	}
	return s
}