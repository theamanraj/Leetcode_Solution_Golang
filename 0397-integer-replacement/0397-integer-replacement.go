func min(vals ...int) int {
	min_v := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] < min_v {
			min_v = vals[i]
		}
	}
	return min_v
}
func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 1 {
		return min(
			integerReplacement((n+1)/2)+1,
			integerReplacement((n-1)/2)+1) + 1
	} else {
		return integerReplacement(n/2) + 1
	}
}