func longestWord(words []string) string {
	m := map[string]bool{}
	for _, w := range words {
		m[w] = true
	}
	max := ""
	var findMax func(string)
	findMax = func(s string) {
		if len(s) > len(max) {
			max = s
		}
		for i := 'a'; i <= 'z'; i++ {
			if m[s+string(i)] {
				findMax(s + string(i))
			}
		}
	}
	findMax("")
	return max
}