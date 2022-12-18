func replaceDigits(s string) string {
	for i := 1; i < len(s); i += 2 {
		s = s[:i] + string((s[i - 1] + s[i]) % 'a' + 49) + s[i + 1:]
	}
	return s
}