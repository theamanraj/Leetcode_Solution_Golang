func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	sBytes := []byte(s)
	
	i, j := 0, len(s)-1
	
	for i < j {
		for i < j && !strings.Contains(vowels, string(sBytes[i])) {
			i++
		}
		
		for i < j && !strings.Contains(vowels, string(sBytes[j])) {
			j--
		}
		
		sBytes[i], sBytes[j] = sBytes[j], sBytes[i]
		i++
		j--
	}
	
	return string(sBytes)
}