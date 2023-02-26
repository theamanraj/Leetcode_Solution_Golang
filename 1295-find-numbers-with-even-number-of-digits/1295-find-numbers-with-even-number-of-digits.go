func findNumbers(nums []int) int {
	var result int
	for _, n := range nums {
		var count int
		num := float32(n)
		for num >= 10.0 {
			num /= 10.0
			count++
		}

		if (count+1)%2 == 0 {
			result++
		}
	}
	return result
}