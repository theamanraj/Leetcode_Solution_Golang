func maximumProduct(nums []int) int {
	n1, n2, n3 := -1<<63, -1<<63, -1<<63
	n4, n5 := 0, 0
	for _, v := range nums {
		if v > n1 {
			n3 = n2
			n2 = n1
			n1 = v
		} else if v > n2 {
			n3 = n2
			n2 = v
		} else if v > n3 {
			n3 = v
		}
		if v < n4 {
			n5 = n4
			n4 = v
		} else if v < n5 {
			n5 = v
		}
	}
    return max(n1*n2*n3, n1*n4*n5)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}