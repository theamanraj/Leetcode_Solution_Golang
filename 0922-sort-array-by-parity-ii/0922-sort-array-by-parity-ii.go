func sortArrayByParityII(slice []int) []int {
	evenIndex := 0
	oddIndex := 1
	ret := make([]int, len(slice))
	for _, ele := range slice {
		if ele%2 == 0 {
			ret[evenIndex] = ele
			evenIndex += 2
		} else {
			ret[oddIndex] = ele
			oddIndex += 2
		}
	}
	return ret
}