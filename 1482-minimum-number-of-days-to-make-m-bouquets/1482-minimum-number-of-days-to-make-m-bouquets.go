func minDays(bd []int, mm int, k int) int {
	if (mm * k) > len(bd) {
		return -1
	}
	l := 0
	r := 0
	for i := range bd {
		if bd[i] > l {
			l = bd[i]
		}
	}

	for l > r {
		m := (l + r) / 2
		vv := 0
		c := 0
		for i := range bd {
			if bd[i] <= m {
				vv++
			} else {
				if vv >= k {
					c += vv / k
				}
				vv = 0
			}
		}
		if vv >= k {
			c += vv / k
		}
		if c >= mm {
			l = m
		} else if c < mm {
			r = m + 1
		}
	}
	return l
}
