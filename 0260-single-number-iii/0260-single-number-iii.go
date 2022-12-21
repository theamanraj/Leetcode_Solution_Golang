func singleNumber(nums []int) []int {
	mapp := make(map[int]bool)
	result := make([]int, 0)

    for _, i := range nums {
		if mapp[i] == true{
           mapp[i] = false
        } else {
           mapp[i] = true
        }
	}

    
    for k, v := range mapp {
		if v  {
			result = append(result, k)
		}
	}
    
	return result
}