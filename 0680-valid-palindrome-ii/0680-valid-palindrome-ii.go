func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isValidPalindrome(l+1, r, s) || isValidPalindrome(l, r-1, s)
		}
		l++
		r--
	}
	return true
}

func isValidPalindrome(l, r int, s string) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
