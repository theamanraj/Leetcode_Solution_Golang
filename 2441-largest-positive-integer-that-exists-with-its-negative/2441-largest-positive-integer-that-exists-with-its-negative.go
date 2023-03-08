func findMaxK(nums []int) int {
	numsMap := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if num < 0 {
			numsMap[num] = struct{}{}
		}
	}

	result := -1
	for _, num := range nums {
		if num > 0 {
			if _, ok := numsMap[-num]; ok {
				if num > result {
					result = num
				}
			}
		}
	}

	return result
}