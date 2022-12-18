func binaryGap(N int) int {
    count := 0
	max := 0
	isPreviousOne := false
	for N > 0 {
		// fmt.Println(N % 2)
		if N%2 == 1 {
			if count > max {
				max = count
			}
			count = 1
			isPreviousOne = true
		} else {
			if isPreviousOne {
				count++
			}
		}
		N = N / 2
	}
	// fmt.Println(max)
	return max
}