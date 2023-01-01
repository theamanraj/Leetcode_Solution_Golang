func permute(nums []int) [][]int {
	var res [][]int
	permuteRec([]int{}, nums, &res)
	return res
}

func permuteRec(currComb, left []int, res *[][]int) {
	if 0 == len(left) {
		*res = append(*res, currComb)
		return
	}
	
	for idx, l := range left {
		permuteRec(
			append(currComb, l),
			append(append([]int{}, left[:idx]...), left[idx+1:]...),
			res,
		)
	}
}