func subsets(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	for k := 0; k <= len(nums); k++ {
		generateSubset(nums, k, 0, c, &res)
	}
	return res
}

func generateSubset(nums []int, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	// k-len(c) is the number count that we needed
	// (n - i +1) length from opposite side should greater than or equal requirement
	// index start from 0
	for i := start; i <= len(nums)-1-(k-len(c))+1; i++ {
		c = append(c, nums[i])
		generateSubset(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}
}