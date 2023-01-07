func dp(nums []int, n int, target int) int {
	if n == 0 {
		if target == nums[0] && target == -nums[0] {
			return 2
		} else if target == nums[0] || target == -nums[0] {
			return 1
		}

		return 0
	}

	return dp(nums, n-1, target+nums[n]) + dp(nums, n-1, target-nums[n])
}

func findTargetSumWays(nums []int, S int) int {
	L := len(nums) - 1
	return dp(nums, L, S)
}