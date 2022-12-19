func findAllConcatenatedWordsInADict(words []string) []string {
    wordMap := getWordMap(words)
	// sort the word array by word length
    sort.Slice(words, func(i, j int)bool{
        return len(words[i]) < len(words[j])
    })

	// first element in the array would be the minimum length one
    minWordLen := len(words[0])
    ans := []string{}

    for i := 1;i < len(words);i++{
        word := words[i]
		// If current word's lenght is less than 2*minWordLen, it's impossible to be made from any two words in the word array
        if len(word) < 2*minWordLen{
            continue
        }
        
        // Check if current word could be successfully break into many different words
        if wordBreak(word, wordMap, minWordLen) {
            ans = append(ans, word)
        }
    }

    return ans
}

func wordBreak(word string, wordMap map[string]bool, minWordLen int)bool{
    dp := make([]bool, len(word)+1)
    for j := minWordLen; j <= len(word); j++{
        if wordMap[word[:j]] && j != len(word){
            dp[j] = true
        }
        for k := minWordLen; k< j;k++{
            if !dp[k]{
                continue
            }
            if wordMap[word[k:j]]{
                dp[j] = true
                break
            }
        }
    }
    
    return dp[len(word)]
}
func getWordMap(words []string)map[string]bool{
    wMap := map[string]bool{}
    
    for _, word := range words{
		// there is a corner test case that input words array contains empty string.
		// It would result in minimum word length as zero
        if len(word) < 1{
            continue
        }
        wMap[word] = true
    }
    
    return wMap
}