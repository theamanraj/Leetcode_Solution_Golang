func shortestToChar(s string, c byte) []int {
	// get all c index in s
	targetIndexes := []int{}
	for i, char := range s {
		if byte(char) == c {
			targetIndexes = append(targetIndexes, i)
		}
	}

	result := []int{}
	for i, char := range s {
		// if char is target, append 0 index
		if byte(char) == c {
			result = append(result, 0)
		} else {
			// check every minimum abs difference
			minimumAbsDiff := math.MaxInt
			for _, targetIndex := range targetIndexes {
				abdDiff := int(math.Abs(float64(i) - float64(targetIndex)))
				if abdDiff < minimumAbsDiff {
					minimumAbsDiff = abdDiff
				}
			}
			result = append(result, minimumAbsDiff)
		}
	}

	return result
}