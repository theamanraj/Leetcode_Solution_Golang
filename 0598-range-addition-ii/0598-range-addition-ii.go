func maxCount(m int, n int, ops [][]int) int {

    if len(ops) == 0 {
        return m * n
    }
	a := []int{}
	b := []int{}

	for _, v := range ops {
		a = append(a, v[0])
		b = append(b, v[1])
	}
    
	sort.Ints(a)
	sort.Ints(b)
    
	return a[0] * b[0]
    
}