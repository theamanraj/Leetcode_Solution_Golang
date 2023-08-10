func maxOperations(nums []int, k int) int {
    counts := make(map[int]int)
	result := 0
	
	for _, num := range nums {
		if counts[k-num] > 0 {
			counts[k-num]--
			result++
		} else {
			counts[num]++
		}
	}
	
	return result
}