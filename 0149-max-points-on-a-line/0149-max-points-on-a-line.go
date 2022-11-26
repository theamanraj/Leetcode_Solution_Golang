func maxPoints(points [][]int) int {
	resPoints := 0
	for i := 0; i < len(points); i++ {
		slopeHash := make(map[string]int)
		infinitySlope := 1
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			if points[i][1] == points[j][1] {
				infinitySlope++
				continue
			}
			x := float64(points[i][0] - points[j][0])
			y := float64(points[i][1] - points[j][1])
			slopeInString := strconv.FormatFloat(x / y, 'f', 8, 64)
			if slopeHash[slopeInString] == 0 {
				slopeHash[slopeInString] = 1
			}
			slopeHash[slopeInString]++
		}
		curPoints := infinitySlope
		for _, v := range slopeHash {
			if v > curPoints {
				curPoints = v
			}
		}
		if curPoints > resPoints {
			resPoints = curPoints
		}
	}
	return resPoints
}