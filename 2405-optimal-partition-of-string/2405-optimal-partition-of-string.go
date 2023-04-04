func partitionString(s string) int {
	result := 1
	running := 0
	for i := 0; i < len(s); i++ {
		c := 1 << (s[i] - 'a')
		if (running & c) > 0 {
			result++
			running = 0
		}
		running |= c
	}
	return result
}