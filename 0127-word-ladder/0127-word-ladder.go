func ladderLength(beginWord string, endWord string, wordList []string) int {
	runesMap := buildRunesMap(append(wordList, beginWord))

	wordMap := map[string]int{}
	for i, w := range wordList {
		wordMap[w] = i
	}
	queue := []string{beginWord}
	depth := 0

	for len(queue) > 0 {
		depth++
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			word := queue[0]
			queue = queue[1:]
			candidates := findCandidates(word, wordMap, runesMap)
			for _, c := range candidates {
				if c == endWord {
					return depth + 1
				} else {
					delete(wordMap, c)
					queue = append(queue, c)
				}
			}
		}
	}

	return 0
}

type RuneMap = map[string][]rune

func findCandidates(word string, words map[string]int, runesMap RuneMap) []string {
	var candidates []string
	for k, _ := range words {
		if isCandidate(word, k, runesMap) {
			candidates = append(candidates, k)
		}
	}
	return candidates
}

func isCandidate(s1 string, s2 string, runeMap RuneMap) bool {
	var cnt int
	s2runes := runeMap[s2]
	for i, c := range runeMap[s1] {
		if s2runes[i] != c {
			cnt++
			if cnt > 1 {
				return false
			}
		}
	}
	return true
}

func buildRunesMap(wordList []string) RuneMap {
	runesMap := RuneMap{}
	for _, w := range wordList {
		runesMap[w] = []rune(w)
	}
	return runesMap
}