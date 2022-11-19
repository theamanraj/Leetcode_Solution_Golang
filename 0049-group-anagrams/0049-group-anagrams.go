func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	var sorted string
	var result [][]string

	for _, str := range strs {
		sorted = sortString(str)
		if _, ok := m[sorted]; ok {
			m[sorted] = append(m[sorted], str)
		} else {
			m[sorted] = []string{str}
		}
	}

	for _, ar := range m {
		result = append(result, ar)
	}

	return result
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}