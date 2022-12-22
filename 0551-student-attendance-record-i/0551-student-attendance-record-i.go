func checkRecord(s string) bool {
    	var cntL, cntA int
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == 'L' {
			cntL++
		} else {
			cntL = 0
		}
		if cntL == 3 {
			return false
		}
		if s[i] == 'A' {
			cntA++
		}
		if cntA == 2 {
			return false
		}
	}
	return true
}