func waysToSplit(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}

	counter := 0
	for rightArrEnd := n - 1; rightArrEnd >= 2; rightArrEnd-- {
		midArrEnd := rightArrEnd - 1
		right := nums[n-1] - nums[midArrEnd]

		lowestLeftEnd := sort.Search(midArrEnd, func(leftArrEnd int) bool {
			middle := nums[midArrEnd] - nums[leftArrEnd]

			return right >= middle
		})

		highestLeftEnd := sort.Search(midArrEnd, func(leftArrEnd int) bool {
			middle, left := nums[midArrEnd]-nums[leftArrEnd], nums[leftArrEnd]

			return middle < left
		})

		highestLeftEnd--

		if highestLeftEnd >= lowestLeftEnd {
			counter += highestLeftEnd - lowestLeftEnd + 1
		}
	}

	return counter % 1_000_000_007
}