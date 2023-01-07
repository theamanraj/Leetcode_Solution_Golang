func canIWin(maxChooseInteger, desiredTotal int) bool {
	totalNumbers := 0
	for i := 1; i <= maxChooseInteger; i++ {
		totalNumbers += i
	}
	if totalNumbers < desiredTotal {
		// reach is impossible
		return false
	}
	if desiredTotal <= maxChooseInteger {
		// first pick reach
		return true
	}
	currentNumbers := 0
	for i := 0; i < maxChooseInteger; i++ {
		currentNumbers += 1 << i
	}
	dp := make([]int, currentNumbers+1)
	return canIWinHelper(dp, maxChooseInteger, desiredTotal, currentNumbers)
}

func canIWinHelper(dp []int, maxChooseInteger, desiredTotal, currentNumbers int) bool {
	if dp[currentNumbers] != 0 {
		return dp[currentNumbers] == 1
	}
	for i := maxChooseInteger-1; i >= 0; i-- {
		if currentNumbers&(0x1<<i) != 0 {
			if i+1 >= desiredTotal {
				dp[currentNumbers] = 1
				return true
			}
			break
		}
	}
	currentNumbersBitOp := currentNumbers
	for i := 0; currentNumbersBitOp != 0; i++ {
		if currentNumbersBitOp & 0x1 != 0 {
			if canIWinHelper(dp, maxChooseInteger, desiredTotal-i-1, currentNumbers^(0x1<<i)) == false {
				dp[currentNumbers] = 1
				return true
			}
		}
		currentNumbersBitOp >>= 1
	}
	dp[currentNumbers] = -1
	return false
}