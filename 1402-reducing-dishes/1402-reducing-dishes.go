func maxSatisfaction(s []int) (likeTime int) {
	sort.Ints(s)
    

    if s[len(s)-1] <= 0 {
        return 0
    }

	
    intSum := 0
    for i := len(s) - 1; i >= 0; i-- {
        newLT := s[i] + likeTime + intSum
        if newLT < likeTime {
            return likeTime
        }
		
        likeTime = newLT
        intSum += s[i]
    } 
    return
}
