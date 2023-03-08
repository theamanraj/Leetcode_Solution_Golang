func oddString(words []string) string {
	mapping := make(map[string][]int, 2)
	for idx, w := range words {
		array := make([]string, 0, len(w)-1)
		for i := 0; i < len(w)-1; i++ {
			array = append(array, strconv.Itoa(int(w[i+1]-w[i])))
		}
		marker := strings.Join(array, ",")
		mapping[marker] = append(mapping[marker], idx)
	}
	for _, v := range mapping {
		if len(v) == 1 {
			return words[v[0]]
		}
	}
	return ""
}