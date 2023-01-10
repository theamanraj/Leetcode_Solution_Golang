func minOperations(s string) int {
	var s1, s2 int
	for i := 0; i < len(s); i++ {
		if s[i] == byte(i%2)+'0' {
			s2++
		} else {
			s1++
		}
	}
	if s1 < s2 {
		return s1
	}
	return s2
}