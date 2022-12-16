func maxDistance(arr []int, m int) int {
    n := len(arr)
	sort.Ints(arr)

	low, high := 0, arr[n-1]
	
	// we place the first ball at position 0
	prev := arr[0]

	mid := -1
	res := 0
	for low <= high {
	
	    // we calculate mid , and try to place all balls (m-1 , as we have already placed the first ball at position 0) at distance of >= mid with each other
		mid = low + (high-low)/2
		k := m - 1
		
		// this is the position of the latest previously placed ball (initially it has position of 1st ball that we placed at arr[0])
		prev = arr[0]

		// here we place balls at m distance with each other 
		for t := 1; t < n; t++ {
		    // when we are able to place any ball at say position arr[x] , this becomes the prev now(to place next balls)
			if arr[t]-prev >= mid {
				k--
				prev = arr[t]
			}
		}

		// if : we were able to place m or greater balls (at distance mid from each other , it means there might be a possiblity that we can place them at a higher distance that each other, so we try in right window of binary search)
		// else : we were not able to place all balls , then we have to shorten our distance b/w balls
		if k <= 0 {
			low = mid + 1
			res = mid
		} else {
			high = mid - 1
		}

	}

    return res
}