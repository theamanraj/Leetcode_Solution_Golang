func minimumAverageDifference(nums []int) int {
	n, total := len(nums), 0
	for _, num := range nums {
		total += num
	}
	min, prefix, idx := total, 0, 0
	for i := 0; i < n-1; i++ {
		prefix += nums[i]
		cur := prefix/(i+1) - (total-prefix)/(n-i-1)
		if cur < 0 {
			cur = -cur
		}
		if cur < min {
			min, idx = cur, i
		}
	}
	if total/n < min {
		return n - 1
	}
	return idx
}