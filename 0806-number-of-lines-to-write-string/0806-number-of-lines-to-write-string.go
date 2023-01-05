func numberOfLines(widths []int, S string) []int {

	h := make(map[rune]int)
	tmp := "abcdefghijklmnopqrstuvwxyz"
    lines := 1
    n := 0

	for i, v := range tmp {
		h[v] = widths[i]
	}

	for _, v := range S {
		n = n + h[v]
		if n > 100 {
			lines++
			n = h[v]
		}
	}

	return []int{lines, n}

}