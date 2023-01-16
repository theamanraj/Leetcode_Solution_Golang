func minMoves2(nums []int) int {
	sort.Ints(nums)
	mid := 0
	if len(nums)%2 == 1 {
		mid = nums[len(nums)/2]
	} else {
		mid = (nums[len(nums)/2] + nums[len(nums)/2-1]) / 2
	}
	ret := 0
	for _, v := range nums {
		if v > mid {
			ret = ret + v - mid
		} else {
			ret = ret + mid - v
		}
	}
	return ret
}