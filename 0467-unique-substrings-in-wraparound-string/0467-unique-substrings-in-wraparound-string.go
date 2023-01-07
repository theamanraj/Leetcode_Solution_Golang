func findSubstringInWraproundString(p string) int {
	var rst int
	if len(p) == 0 {
		return rst
	}

	rstSlice := [26]int{} // a开头的，b开头的

	for i := 0; i < len(p); {
		index := i + 1
		for index < len(p) && (p[index] == p[index-1]+1 || p[index] == p[index-1]-25) {
			index++
		}
		for j := i; j < index; j++ {
			id := int(p[j]) - int('a')
			rstSlice[id] = int(math.Max(float64(rstSlice[id]), float64(index-j)))
		}
		i = index
	}

	for _, v := range rstSlice {
		rst += v
	}

	return rst
}
