func minEatingSpeed(piles []int, h int) int {
    sort.Ints(piles)

    left := 1
    right := piles[len(piles)-1]
    
    for left < right {
        mid := left + (right - left)/2
        
        // calculate time spent on piles
        timeSpent := 0
        for _, pile := range piles {
            timeSpent += (pile/mid) 
            if pile % mid != 0 {
                timeSpent++
            }
        }

        if timeSpent <= h {
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return right // right is workable solution
}