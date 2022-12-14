func chalkReplacer(chalk []int, k int) int {
	s := 0
	for i := range chalk {
		s += chalk[i]
	}
	x := k / s
	k -= (x * s)
	for i := range chalk {
		if chalk[i] > k {
			return i
		}
		k -= chalk[i]
	}
	return 0
}