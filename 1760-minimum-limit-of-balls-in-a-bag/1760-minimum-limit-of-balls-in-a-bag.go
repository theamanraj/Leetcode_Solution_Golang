func minimumSize(nums []int, maxOperations int) int {
    // decrease sort
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	leftBound, rightBound := 1, nums[0]
    // binary search
	for leftBound <= rightBound {
		m := (leftBound+rightBound)>>1
		if possible(nums, maxOperations, m) {
			rightBound = m - 1
		} else {
			leftBound = m + 1
		}
	}
	return rightBound+1
}

func possible(nums []int, maxOperations int, maxSize int) bool {
	for _, v := range nums {
		if v <= maxSize {
			return true
		}
		cost := v / maxSize
		if v % maxSize == 0 {
			cost--
		}
		maxOperations -= cost
		if maxOperations < 0 {
			return false
		}
	}
	return true
}