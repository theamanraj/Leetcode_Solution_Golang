func findRepeatedDnaSequences(s string) []string {
    counts := map[string]int{}
	output := []string{}
	for i := 0; i + 10 <= len(s); i++ {
		current := string(s[i: i + 10])
		v, k := counts[current]
		if k {
			if v == 1 {
				output = append(output, current)
			}
			counts[current]++
		} else {
			counts[current] = 1
		}
	}
	return output
}