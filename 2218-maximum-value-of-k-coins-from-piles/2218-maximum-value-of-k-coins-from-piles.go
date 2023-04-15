type node struct {
	i, remain int
}

func maxValueOfCoins(piles [][]int, k int) int {

	var dp func(i, remain int) int

	memo := make(map[node]int, k)

	dp = func(i, remain int) int {
		if i == len(piles) || remain == 0 {
			return 0
		}

		if val, ok := memo[node{i, remain}]; ok {
			return val
		}

		ans := dp(i+1, remain)
		curr := 0
		for j := 0; j < min(len(piles[i]), remain); j++ {
			curr += piles[i][j]
			ans = max(ans, curr+dp(i+1, remain-j-1))
		}

		memo[node{i, remain}] = ans
		
		return ans
	}

	return dp(0, k)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}