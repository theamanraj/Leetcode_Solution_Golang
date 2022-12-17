func longestPalindrome(words []string) int {
	palindromeSize := 0
	hash := make(map[string]int)
	for _, word := range words {
		hash[word]++
	}
	middle := 0
	for word, count := range hash {
		pair := getPair(word)
		if word == pair {
			if count % 2 != 0 {
				middle = 2
			}
			palindromeSize += count / 2 * 4
			continue
		}
		pairCount := hash[pair]
		costCount := min(count, pairCount)
		if costCount == 0 {
			continue
		}
		hash[word] -= costCount
		hash[pair] -= costCount
		palindromeSize += costCount * 4
	}
	return palindromeSize + middle
}

func getPair(word string) string {
	return string(word[1]) + string(word[0])
}

func min(values ...int) int {
	minValue := math.MaxInt64
	for _, v := range values {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}