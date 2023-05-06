func minAreaRect(points [][]int) int {
	hashX := map[int][]int{}
	for _, p := range points {
		hashX[p[0]] = append(hashX[p[0]], p[1])
	}
	allX := make([]int, 0, len(hashX))
	for k, v := range hashX {
		allX = append(allX, k)
		sort.Ints(v)
	}
	sort.Ints(allX)
	minArea := math.MaxInt32
	for i := 0; i < len(allX); i++ {
		for j := i+1; j < len(allX); j++ {
			if temp := commonShortest(hashX[allX[i]], hashX[allX[j]]); temp != math.MaxInt32 {
				curr := temp * (allX[j]-allX[i])
				if curr < minArea {
					minArea = curr
				}
			}
		}
	}
	if minArea == math.MaxInt32 {
		return 0
	}
	return minArea
}

func commonShortest(s1, s2 []int) int {
	values := []int{}
	for i, j := 0, 0; i < len(s1) && j < len(s2); {
		if s1[i] == s2[j] {
			values = append(values, s1[i])
			i++
			j++
		} else if s1[i] > s2[j] {
			j++
		} else {
			i++
		}
	}
	shortest := math.MaxInt32
	for i := 0; i < len(values)-1; i++ {
		if values[i+1] - values[i] < shortest {
			shortest = values[i+1] - values[i]
		}
	}
	return shortest
}