func findMinDifference(timePoints []string) int {
	mins := timesToInts(timePoints)
	sort.Slice(mins, func(i, j int) bool { return mins[i] < mins[j] })

	minDiff := math.MaxInt32
	for i := 1; i < len(mins); i++ {
		if mins[i]-mins[i-1] < minDiff {
			minDiff = mins[i] - mins[i-1]
		}
	}

	if mins[0]+24*60-mins[len(mins)-1] < minDiff {
		return mins[0] + 24*60 - mins[len(mins)-1]
	}
	return minDiff
}

func timesToInts(timePoints []string) []int {
	resp := make([]int, 0, len(timePoints))
	for _, t := range timePoints {
		hours, _ := strconv.Atoi(t[:2])
		min, _ := strconv.Atoi(t[3:])
		resp = append(resp, hours*60+min)
	}
	return resp
}
