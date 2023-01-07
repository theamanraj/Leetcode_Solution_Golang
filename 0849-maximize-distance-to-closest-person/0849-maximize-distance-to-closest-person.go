package main

func maxDistToClosest(seats []int) int {
	left := -1
	distance := 0
	for i := 0; i < len(seats); i++ {
		if seats[i] != 0 {
			if left == -1 {
				distance = i
			} else {
				distance = max(distance, (i-left)/2)
			}
			left = i
		}
	}
	if seats[len(seats)-1] == 0 {
		distance = max(len(seats)-left-1, distance)
	}
	return distance
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}