func dfs(wordIndex int, words []string, letterCounts []int, score []int) int {
	if wordIndex == len(words) {
		return 0
	}

	nextCounts := make([]int, len(letterCounts))
	for i := 0; i < len(letterCounts); i += 1 {
		nextCounts[i] = letterCounts[i]
	}

	currentWordScore := 0
	for i := 0; i < len(words[wordIndex]); i += 1 {
		letterCode := words[wordIndex][i] - 97
		if nextCounts[letterCode]-1 >= 0 {
			nextCounts[letterCode] -= 1
			currentWordScore += score[letterCode]
		} else {
			currentWordScore = 0
			break
		}
	}
    // ----------
	includeWordScore := currentWordScore + dfs(wordIndex+1, words, nextCounts, score)
	excludeWordScore := dfs(wordIndex+1, words, letterCounts, score)

	return max(includeWordScore, excludeWordScore)
}

func maxScoreWords(words []string, letters []byte, score []int) int {
    wordIndex := 0
    letterCounts := readLetterCounts(letters)
    return dfs(wordIndex, words, letterCounts, score)
}

// Helper functions
// ----------
func readLetterCounts(letters []byte) []int {
    letterCounts := make([]int, 26)
    for _, letter := range letters {
        letterCounts[letter - 97] += 1
    }
    return letterCounts
}

func max(includeWordScore int, excludeWordScore int) int {
    if includeWordScore > excludeWordScore {
        return includeWordScore
    }
    return excludeWordScore
}