func minimumOperations(nums []int) int {
	currentNumber := 1
	getArrayLens := len(nums)
	var outputNums int
	sort.Ints(nums)

	for count := 0; count < getArrayLens; count++ {

		for _, v := range nums {
			if v > 0 {
				currentNumber = v
				outputNums++
				break
			}
		}

		for k, v := range nums {
			if v > 0 {
				nums[k] -= currentNumber
			}
		}
	}

	return outputNums
}