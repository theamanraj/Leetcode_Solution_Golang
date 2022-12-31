func avoidFlood(rains []int) []int {
	lakes := map[int]int{}
	ans := []int{}
	dryable := []int{}

	for i, rain := range rains {
		if rain == 0 {
			dryable = append(dryable, len(ans))
			ans = append(ans, 1)
		} else {
			if lakes[rain] > 0 {
				dryed := -1
				for i, dry := range dryable {
					if dry >= lakes[rain] {
						ans[dry] = rain
						dryed = i
						break
					}
				}

				if dryed == -1 {
					return []int{}
				}

				dryable = remove(dryable, dryed)
			}

			lakes[rain] = i + 1
			ans = append(ans, -1)
		}
	}

	return ans
}

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}