func specialArray(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for x := 0; x < len(nums); x++ {
		if nums[x] == x {
			return -1
		}

		if nums[x] < x+1 {
			return x
		}
	}

	return len(nums)
}