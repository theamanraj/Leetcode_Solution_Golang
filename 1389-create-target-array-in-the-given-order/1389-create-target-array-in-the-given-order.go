func createTargetArray(nums []int, index []int) []int {
	result := make([]int, 0, len(nums))
	for i, v := range index {
		result = append(result[:v+1], result[v:]...)
		result[v] = nums[i]
	}
	return result
}