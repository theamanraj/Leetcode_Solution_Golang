func candy(ratings []int) int {
	candyToGive := make([]int, len(ratings))
	current := make([]int, 0)
	for i, v := range ratings {
		if i-1 >= 0 && v > ratings[i-1] {
			continue
		}
		if i+1 < len(ratings) && v > ratings[i+1] {
			continue
		}
		// child-i not higher than left & right
		// 1 candy is OK, silly child, :)
		candyToGive[i] = 1
		current = append(current, i)
	}
	for len(current) > 0 {
		// rank from low->high
		next := make([]int, 0)
		for _, index := range current {
			if index-1 >= 0 && ratings[index-1] > ratings[index] {
				if candyToGive[index-1] <= candyToGive[index] {
					candyToGive[index-1] = candyToGive[index] + 1
				}
				next = append(next, index-1)
			}
			if index+1 < len(ratings) && ratings[index+1] > ratings[index] {
				if candyToGive[index+1] <= candyToGive[index] {
					candyToGive[index+1] = candyToGive[index] + 1
				}
				next = append(next, index+1)
			}
		}
		current = next
	}
	totalCandies := 0
	for _, v := range candyToGive {
		totalCandies += v
	}
	return totalCandies
}