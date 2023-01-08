func numEquivDominoPairs(dominoes [][]int) int {
	var res int
	hash := make(map[[2]int]int)
	for _, v := range dominoes {
		if v[0] > v[1] {
			v[0], v[1] = v[1], v[0]
		}
		tmp := [2]int{v[0], v[1]}
		if v, _ := hash[tmp]; v>0 {
			res += v
		}
		hash[tmp]++
	}
	return res
}