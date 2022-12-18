func checkIfPangram(sentence string) bool {
    if len(sentence) < 26 {
        return false
    }
	m := make(map[rune]struct{})
	for _, v := range sentence {
		m[v] = struct{}{}
	}
	return len(m) == 26
}