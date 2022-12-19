func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	charCount := make(map[rune]int)
	for _, c := range s {
		cnt, ok := charCount[c]
		if ok {
			charCount[c] = cnt + 1
		} else {
			charCount[c] = 1
		}
	}
	minCount := len(s)
	var minCountChar rune
	for k, v := range charCount {
		if v < minCount {
			minCountChar = k
			minCount = v
		}
	}
	if minCount >= k {
		return len(s)
	}
	subStrs := strings.Split(s, string(minCountChar))
	maxSubStrCount := 0
	for _, subStr := range subStrs {
		subStrCount := longestSubstring(subStr, k)
		if subStrCount > maxSubStrCount {
			maxSubStrCount = subStrCount
		}
	}
	return maxSubStrCount
}
