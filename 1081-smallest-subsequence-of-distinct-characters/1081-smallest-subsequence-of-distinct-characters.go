func smallestSubsequence(s string) string {
	count := [26]int{}
	vis := [26]bool{}
	var ans string
	for _, ch := range s {
		count[ch-'a']++
	}
	for _, ch := range s {
		count[ch-'a']--
		if !vis[ch-'a'] {
			for len(ans) != 0 && ans[len(ans)-1] > byte(ch) && count[ans[len(ans)-1]-'a'] > 0 {
				vis[ans[len(ans)-1]-'a'] = false
				ans = ans[0 : len(ans)-1]
			}
			ans += string(ch)
			vis[ch-'a'] = true
		}
	}
	return ans
}