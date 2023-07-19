func balancedString(s string) int {
	seen := make(map[byte]int, 4)

	n := len(s)
	res, k := n, n/4
	var start int

	// count frequency of each character in s
	for i := range s {
		seen[s[i]]++
	}

	// sliding window
	for end := range s {
		seen[s[end]]--
		for start < n && seen['Q'] <= k && seen['W'] <= k && seen['E'] <= k && seen['R'] <= k {
			if ws := end - start + 1; ws < res {
				res = ws
			}
			seen[s[start]]++
			start++
		}
	}
	return res
}