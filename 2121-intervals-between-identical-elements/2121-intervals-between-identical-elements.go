func getDistances(arr []int) []int64 {
	indexes := make([][]int, 100001)
	distances := make([]int64, len(arr))
	for i, v := range arr {
		indexes[v] = append(indexes[v], i)
	}
	for _, valueIndexes := range indexes {
		if len(valueIndexes) == 0 {
			continue
		}
		if len(valueIndexes) == 1 {
			distances[valueIndexes[0]] = 0
			continue
		}
		ltSummary, geSummary := 0, sum(valueIndexes)
		ltCount, geCount := 0, len(valueIndexes)
		for _, index := range valueIndexes {
			distances[index] = int64(geSummary - geCount*index + ltCount*index - ltSummary)
			ltCount++
			geCount--
			ltSummary += index
			geSummary -= index
		}
	}
	return distances
}

func sum(values []int) int {
	summary := 0
	for _, v := range values {
		summary += v
	}
	return summary
}