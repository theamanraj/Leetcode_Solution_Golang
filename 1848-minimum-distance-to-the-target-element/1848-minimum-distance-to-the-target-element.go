func getMinDistance(nums []int, target int, start int) int {
	indicesOfTarget := []int{}
	minSoFar := math.MaxInt64

	for idx, num := range nums {
		if num == target {
			indicesOfTarget = append(indicesOfTarget, idx)
			deviation := int(math.Abs(float64(idx - start)))
			
			if deviation < minSoFar {
				minSoFar = deviation
			}
		}
	}
	return minSoFar
}