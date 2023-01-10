func longestNiceSubstring(s string) string {
	if len(s) < 2 {
		return ""
	}

	chars := make(map[rune]struct{}, 0)
	for _, r := range s {
		chars[r] = struct{}{}
	}

	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		_, u := chars[unicode.ToUpper(r)]
		_, l := chars[unicode.ToLower(r)]
		if u && l {
			continue
		}
		left := longestNiceSubstring(s[:i])
		right := longestNiceSubstring(s[i+1:])
		if len(left) >= len(right) {
			return left
		} else {
			return right
		}
	}
	return s
}