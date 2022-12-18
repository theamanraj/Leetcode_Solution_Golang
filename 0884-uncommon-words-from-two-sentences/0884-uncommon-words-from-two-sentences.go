func uncommonFromSentences(s1 string, s2 string) []string {

	var wordMap = make(map[string]int)
	for _, word := range strings.Split(s1, " ") {
		wordMap[word]++
	}

	for _, word := range strings.Split(s2, " ") {
		wordMap[word]++
	}

	var uncommonWords []string

	for word, count := range wordMap {
		if count == 1 {
			uncommonWords = append(uncommonWords, word)
		}
	}

	return uncommonWords
}