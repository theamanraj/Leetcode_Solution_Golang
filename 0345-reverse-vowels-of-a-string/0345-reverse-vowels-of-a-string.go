func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	sBytes := []byte(s)
	
	i, j := 0, len(s)-1
	
	for i < j {
		for i < j && !strings.Contains(vowels, string(sBytes[i])) {
			i++
            fmt.Println(i,"ith Attempt")
		}
		
		for i < j && !strings.Contains(vowels, string(sBytes[j])) {
			j--
            fmt.Println(j,"jth Attempt")
		} 
		
		sBytes[i], sBytes[j] = sBytes[j], sBytes[i]
        fmt.Println("nth Attempt")
		i++
		j--
	}
	
	return string(sBytes)
}