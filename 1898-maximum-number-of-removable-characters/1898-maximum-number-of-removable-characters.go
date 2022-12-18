func maximumRemovals(s string, p string, removable []int) int {
	if len(removable) > len(s)-len(p) {
		removable = removable[:len(s)-len(p)]
	}
	l, r := 0, len(removable)-1
	for l <= r {
		m := (l+r)>>1
		hash := make(map[int]bool)
		for i := 0; i <= m; i++ {
			hash[removable[i]] = true
		}
		if isSubsequence(s, p, hash) {
			l = m+1
		} else {
			r = m-1
		}
	}
	return r+1
}

func isSubsequence(s string, p string, hash map[int]bool) bool {
	res := len(s)
	for i := 0; i < len(p); {
		for hash[len(s)-res] == true {
			res--
		}
		if res < len(p)-i {
			return false
		}
		if p[i] == s[len(s)-res] {
			i++
		}
		res--
	}
	return true
}