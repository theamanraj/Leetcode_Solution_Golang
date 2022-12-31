func findLongestWord(s string, d []string) string {
	bytes := []byte(s)
	longest := ""
	for _, word := range d {
		if len(word) < len(longest) {
			continue
		}
		i := 0
		for c := 0; c < len(bytes); c++ {
			if bytes[c] == word[i] {
				i++
			}
			if i == len(word) {
				if i > len(longest) || word < longest {
					longest = word
				}
				break
			}
		}
	}
	return longest
}