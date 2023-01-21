func findShortestSubArray(nums []int) int {

	tmp := make(map[int][]int)
	degree := 0
	subArr := 50000

	for i, v := range nums {
		tmp[v] = append(tmp[v], i)
		if len(tmp[v]) > degree {
			degree = len(tmp[v])
		}
	}
	for _, v := range tmp {
		if len(v) == degree {
			l := v[len(v)-1] - v[0] + 1
			if l < subArr {
				subArr = l
			}
		}
	}

	return subArr

}