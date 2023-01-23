func shortestCompletingWord(licensePlate string, words []string) string {
	chars := toChars(licensePlate)

	var shortest string
	wordsLoop:
	for _, word := range words {
		wordChars := toChars(word)
		for i := 0; i < 26; i++ {
			if chars[i] > wordChars[i] {
				continue wordsLoop
			}
		}
		if shortest == "" || len(word) < len(shortest) {
			shortest = word
		}
	}
	return shortest
}

func toChars(word string) [26]int {
	var chars [26]int
	for _, r := range word {
		if unicode.IsLetter(r) {
			chars[unicode.ToLower(r)-'a']++
		}
	}
	return chars
}