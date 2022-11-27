func nearestPalindromic(n string) string {
	nValue, _ := strconv.Atoi(n)
	if len(n) == 1 {
		return strconv.Itoa(nValue-1)
	}
	// 1000...001, len(n)+1
	nearestValue, _ := strconv.Atoi("1"+strings.Repeat("0", len(n)-1)+"1")
	nearestDiff := nearestValue - nValue
	if len(n) > 1 {
		// 99...99, len(n)+1
		currValue, _ := strconv.Atoi(strings.Repeat("9", len(n)-1))
		currDiff := nValue - currValue
		if currDiff <= nearestDiff {
			nearestValue = currValue
			nearestDiff = currDiff
		}
	}
	magic := 0
	if len(n) % 2 == 1 {
		// odd
		magic = 1
	}
	// even
	leftPart := n[:len(n)/2+magic]
	leftPartValue, _ := strconv.Atoi(leftPart)
	leftPartRuneSlice := []rune(leftPart)
	reverse(leftPartRuneSlice)
	currValue, _ := strconv.Atoi(leftPart + string(leftPartRuneSlice[magic:]))
	currDiff := abs(nValue - currValue)
	if currValue != nValue {
		if currDiff < nearestDiff {
			nearestDiff = currDiff
			nearestValue = currValue
		} else if currDiff == nearestDiff && currValue < nearestValue {
			nearestValue = currValue
		}
	}
	leftPartMoreValue := strconv.Itoa(leftPartValue+1)
	leftPartMoreSlice := []rune(leftPartMoreValue)
	reverse(leftPartMoreSlice)
	currValue, _ = strconv.Atoi(leftPartMoreValue + string(leftPartMoreSlice[magic:]))
	currDiff = abs(nValue - currValue)
	if currDiff < nearestDiff {
		nearestDiff = currDiff
		nearestValue = currValue
	} else if currDiff == nearestDiff && currValue < nearestValue {
		nearestValue = currValue
	}
	leftPartLessValue := strconv.Itoa(leftPartValue-1)
	leftPartLessSlice := []rune(leftPartLessValue)
	reverse(leftPartLessSlice)
	currValue, _ = strconv.Atoi(leftPartLessValue + string(leftPartLessSlice[magic:]))
	currDiff = abs(nValue - currValue)
	if currDiff < nearestDiff {
		nearestDiff = currDiff
		nearestValue = currValue
	} else if currDiff == nearestDiff && currValue < nearestValue {
		nearestValue = currValue
	}
	return strconv.Itoa(nearestValue)
}

func abs(v int) int {
	if v > 0 {
		return v
	}
	return v * -1
}

func reverse(A []rune) {
	for i := 0; i < len(A)/2; i++ {
		A[i], A[len(A)-1-i] = A[len(A)-1-i], A[i]
	}
}