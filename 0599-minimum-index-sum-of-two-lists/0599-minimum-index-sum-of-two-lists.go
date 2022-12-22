func findRestaurant(list1 []string, list2 []string) []string {
    l1 := make(map[string]int)

	for i, s := range list1 {
		l1[s] = i
	}
	
	out := 3000
	var res []string
	for i, s := range list2 {
		if _, ok := l1[s]; ok {
			sum := i + l1[s]
			if sum < out {
				out = sum
				res = []string{s}
			} else if out == sum {
				res = append(res, s)
			}

		}
	}

	return res
}