func isAdditiveNumber(num string) bool {

	x := make([]int, 0)
	var bt func(int)
	r := false
	bt = func(i int) {
		if r {
			return
		}
		if i >= len(num)-1 {
			if len(x) >= 3 {
				r = true
			}
			return
		}

		rr := []byte{}
		for ni := i + 1; ni < len(num); ni++ {
			if len(rr) == 0 && num[ni] == 48 {
				x = append(x, 0)
				if len(x) >= 3 {
					if x[len(x)-3]+x[len(x)-2] == x[len(x)-1] {
						bt(ni)
					}
				} else {
					bt(ni)
				}
				x = x[:len(x)-1]
				return
			}

			rr = append(rr, num[ni])
			v, _ := strconv.ParseInt(*(*string)(unsafe.Pointer(&rr)), 10, 64)
			x = append(x, int(v))
			if len(x) >= 3 {
				if x[len(x)-3]+x[len(x)-2] == x[len(x)-1] {
					bt(ni)
				}
			} else {
				bt(ni)
			}
			x = x[:len(x)-1]

		}
	}
	bt(-1)
	return r
}
