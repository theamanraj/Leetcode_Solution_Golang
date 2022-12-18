func charsExistAndFrequencies(s string) (exist int32, frequencies [26]int) {
	for i := 0; i < len(s); i++ {
		cIdx := s[i] - 'a'
		exist |= 1 << cIdx
		frequencies[cIdx]++
	}
	return
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	e1, f1 := charsExistAndFrequencies(word1)
	e2, f2 := charsExistAndFrequencies(word2)
	if e1 != e2 {
		return false
	}
	sort.Ints(f1[:])
	sort.Ints(f2[:])
	return f1 == f2
}