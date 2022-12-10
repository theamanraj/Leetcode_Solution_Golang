func findPeakElement(nums []int) int {
	old := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] > old && (i+2 > len(nums) || i+1 <= len(nums)-1 && nums[i] > nums[i+1]) {
			return i
		}
		old = nums[i]
	}
	return 0
}
