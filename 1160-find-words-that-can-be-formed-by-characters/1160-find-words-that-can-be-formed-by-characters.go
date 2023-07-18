func countCharacters(words []string, chars string) int {
	charMap := make([]int, 26)
	charArr := []rune(chars)
	for _, val := range charArr {
		charMap[int(val)-97]++
	}

	totalSum := 0
	for _, word := range words {
		wordArr := []rune(word)
		wordDone := true
		localCharMap := make([]int, 26)
		for _, val := range wordArr {
			idx := int(val) - 97
			localCharMap[idx]++
			if localCharMap[idx] > charMap[idx] {
				wordDone = false
				break
			}
		}

		if wordDone {
			totalSum += len(word)
		}
	}

	return totalSum
}