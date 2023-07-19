func eraseOverlapIntervals(intervals [][]int) int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	count := 0
	left, right := 0, 1

	for right < len(intervals) {

		// non overlapping case
		if intervals[left][1] <= intervals[right][0] {
            left = right
			right += 1
			// overlapping case 1
		} else if intervals[left][1] <= intervals[right][1] {
			count += 1
			right += 1
			// overlapping case 2
		} else if intervals[left][1] > intervals[right][1] {
			left = right
			count += 1
			right += 1
		}

	}
	return count
}