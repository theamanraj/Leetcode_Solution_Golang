func maxVowels(s string, k int) int {
    vowels := "aeiouAEIOU"
	maxVowelCount := 0
	currentVowelCount := 0
	
	for i := 0; i < len(s); i++ {
		if i >= k && strings.Contains(vowels, string(s[i-k])) {
			currentVowelCount--
		}
		
		if strings.Contains(vowels, string(s[i])) {
			currentVowelCount++
		}
		
		if currentVowelCount > maxVowelCount {
			maxVowelCount = currentVowelCount
		}
	}
	
	return maxVowelCount
}