func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
func getMaxLen(nums []int) int {
	res := 0
	firstZero := -1
	firstNegative := -1
	totalNegative := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			if firstNegative == -1 {
				firstNegative = i
			}
			totalNegative++
		}
		if nums[i] == 0 {
			firstZero = i
			firstNegative = -1
			totalNegative = 0
		}
		if totalNegative%2 == 0 {
			res = max(res, i-firstZero)
		} else {
			res = max(res, i-firstNegative)
		}
	}
	return res
}