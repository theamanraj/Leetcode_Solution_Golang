func maxVowels(s string, k int) int {
	maxVowelCount := 0
	currentVowelCount := 0
	
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			currentVowelCount++
		}
		
		if i >= k && isVowel(s[i-k]) {
			currentVowelCount--
		}
		
		if currentVowelCount > maxVowelCount {
			maxVowelCount = currentVowelCount
		}
	}
	
	return maxVowelCount
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}