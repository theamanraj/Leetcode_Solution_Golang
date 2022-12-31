func findLUSlength(strs []string) int {
	hash := map[string]int{}
	for _, str := range strs {
		hash[str]++
	}
	distinct := make([]string, 0, len(hash))
	for k := range hash {
		distinct = append(distinct, k)
	}
	sort.Slice(distinct, func(i, j int) bool {
		return len(distinct[i]) > len(distinct[j])
	})
	mainLoop:
	for i, v := range distinct {
		if hash[v] == 1 {
			for j := 0; j < i; j++ {
                if hash[distinct[j]] > 1 && len(distinct[j]) > len(v) {
					if isSubSequence(distinct[j], v) {
						continue mainLoop
					}
				}
			}
			return len(v)
		}
	}
	return -1
}

func isSubSequence(str1, str2 string) bool {
	if len(str1) <= len(str2) {
		return false
	}
	i1, i2 := 0, 0
	for i1 < len(str1) {
        if len(str2)-i2 > len(str1)-i1 {
            return false
        }
		if str1[i1] == str2[i2] {
			i1++
			i2++
			if i2 == len(str2) {
				return true
			}
		} else {
			i1++
		}
	}
	return false
}