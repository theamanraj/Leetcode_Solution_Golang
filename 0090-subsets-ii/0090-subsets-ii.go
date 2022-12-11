func subsetsWithDup(n []int) [][]int {
	sort.Ints(n)
	rt := [][]int{[]int{}}

	var lastN, lastAdd int
	if len(n) > 0 {
		rt = append(rt, []int{n[0]})
		lastN = n[0]
		lastAdd = 1
	}

	for i := 1; i < len(n); i++ {
		if n[i] != lastN {
			lastAdd = len(rt)
		}
		lastN = n[i]
		rt = addComb(rt, n[i], lastAdd)
	}
	return rt
}

func addComb(n [][]int, j, m int) [][]int {
	//fmt.Println(n)
	//fmt.Println(j, m)
	newComb := make([][]int, m)

	for i := 0; i < m; i++ {
		nc := append([]int{}, n[len(n)-1-i]...)
		//fmt.Println(n[len(n)-i-1], nc)
		newComb[i] = append(nc, j)
	}

	return append(n, newComb...)
}