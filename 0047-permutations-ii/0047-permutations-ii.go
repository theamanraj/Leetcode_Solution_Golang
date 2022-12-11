func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	preRet := make([][]int, 0)

	ls := permuteUnique(nums[1:])
	for index := range ls {
		preRet = append(preRet, addOnePermute(nums[0], ls[index])...)
	}
	ret := make([][]int, 0)
	unique := make(map[string]struct{})
	for _, v := range preRet {
		s := lsToStr(v)
		if _, ok := unique[s]; ok {
			continue
		}
		unique[s] = struct{}{}
		ret = append(ret, v)
	}
	return ret
}

func addOnePermute(k int, ls []int) [][]int {
	ret := make([][]int, len(ls)+1)
	for i := range ret {
		ret[i] = make([]int, len(ls)+1)
		for index := 0; index < len(ret[i]); index++ {
			if index < i {
				ret[i][index] = ls[index]
				continue
			}
			if index == i {
				ret[i][index] = k
				continue
			}
			if index > i {
				ret[i][index] = ls[index-1]
				continue
			}
		}
	}
	return ret
}

func lsToStr(ls []int) string {
	return fmt.Sprintf("%v", ls)
}