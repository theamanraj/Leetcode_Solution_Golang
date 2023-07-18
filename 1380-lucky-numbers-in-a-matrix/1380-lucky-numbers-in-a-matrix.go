func luckyNumbers(matrix [][]int) []int {
	var arr []int
	for _, val1 := range matrix {
		var m1 = make(map[int]int)
		x := MinInRow(val1)
		for i2, val2 := range val1 {
			m1[val2] = i2
		}
		z := m1[x]
		max := 0
		for i3 := 0; i3 < len(matrix); i3++ {
			if matrix[i3][z] > max {
				max = matrix[i3][z]
			}
		}
		if x == max {
			arr = append(arr, x)
		}
	}
	return arr
}
func MinInRow(a []int) int {
	var b []int
	b = append(b, a...)
	sort.Ints(b)
	return b[0]
}