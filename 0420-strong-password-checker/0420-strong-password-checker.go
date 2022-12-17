func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func strongPasswordChecker(s string) int {
	// #1
	needLower := true
	needUpper := true
	needDigit := true
	neededTypeChange := 3
	// a -> 1, aa -> 20, aaa -> 300, aaaa -> 4000
	repeats := make([]int, len(s))

	// #2
	// Look for needed type change and counting repeats
	for i := 0; i < len(s); {
		c := s[i]
		if needLower && isLower(c) {
			needLower = false
			neededTypeChange--
		}
		if needUpper && isUpper(c) {
			needUpper = false
			neededTypeChange--
		}
		if needDigit && isDigit(c) {
			needDigit = false
			neededTypeChange--
		}

		f := i
		for i < len(s) && s[i] == c {
			i++
		}
		repeats[f] = i - f
	}

	// #3
	// Too short, appending and/or needed type change
	if len(s) < 6 {
		return max(neededTypeChange, 6-len(s))
	}

	// #4

	neededChanges := 0
	forDelete := max(len(s)-20, 0)
	replaces := 0
	// chars over 20 must be deleted, but keep in mind: already deleted should not be replaced (see later on)
	neededChanges += forDelete

	// Skip replacing of dangling chars, which already deleted first
	// r=1: [aaa, ...]aaa -> [aaa, ...]aa
	// r=2: [aaa, ...]aaaa -> [aaa, ...]aa
	// the [aaa, ...] will be solved by replacing each 3rd (if not deleted at all)
	for r := 1; r < 3; r++ {
		for i := 0; i < len(repeats) && forDelete > 0; i++ {
			if repeats[i] < 3 || repeats[i]%3 != (r-1) {
				continue
			}
			repeats[i] -= min(forDelete, r) // skip replace if there is enough delete (can be 1, if r=2)
			forDelete -= r                  // decrease replace plan
		}
	}

	// count replacing
	for i := 0; i < len(repeats); i++ {
		if repeats[i] >= 3 && forDelete > 0 { // skip deleted chars from replacing
			need := repeats[i] - 2  // [aaa, ...] are already deleted, until forDelete; dangling aa can be remained
			repeats[i] -= forDelete // can be negative, don't worry
			forDelete -= need       // can be negative, don't worry
		}

		// we deleted repeates until len(s)>20
		// the remained aaa chunks will be replaced to aaT
		if repeats[i] >= 3 {
			replaces += repeats[i] / 3
		}
	}

	// at least needed type change must be done
	neededChanges += max(neededTypeChange, replaces)

	return neededChanges
}