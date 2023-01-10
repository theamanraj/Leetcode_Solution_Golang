func maxRepOpt1(text string) int {
	var consecMap, cntMap map[int]int = make(map[int]int), make(map[int]int)
	var maxLen int = 0
	for _, c := range text {
		if _, ok := cntMap[int(c)]; ok == false {
			cntMap[int(c)] = 1
		} else {
			cntMap[int(c)]++
		}

	}
	for i := 0; i < len(text); {
		var cnt int = 0
		c := text[i]
		for j := i; j < len(text); j++ {
			if text[j] == c {
				cnt++
			} else {
				break
			}
		}
		consecMap[i] = cnt
		if cnt > maxLen {
			maxLen = cnt
		}
		i += cnt
	}
	for k, v := range consecMap {
		var ch int = int(text[k])
		if k+v+1 < len(text) && text[k] == text[k+v+1] {
			var tmp int = 0
			if v+consecMap[k+v+1] < cntMap[ch] {
				tmp = v + consecMap[k+v+1] + 1
			} else {
				tmp = v + consecMap[k+v+1]
			}
			if tmp > maxLen {
				maxLen = tmp
			}
		}
		if v < cntMap[ch] && v+1 > maxLen {
			maxLen = v + 1
		}
	}
	return maxLen
}