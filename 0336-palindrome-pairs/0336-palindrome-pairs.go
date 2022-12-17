import "strconv"

func palindromePairs(words []string) [][]int {
	ans := [][]int{}
	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		m[words[i]] = i
	}

	added := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		word := words[i]
		if word == "" {
			continue
		}

		dp := constructDP(word)
		reversedWord := reverse(word)
		for j := 0; j <= len(word); j++ {
			if j == 0 || dp[0][j-1] {
				r := len(word) - j
				substring := reversedWord[:r]
				if index, ok := m[substring]; ok && index != i {
					key := strconv.Itoa(index) + "-" + strconv.Itoa(i)
					if _, ok1 := added[key]; !ok1 {
						ans = append(ans, []int{index, i})
						added[key] = true
					}
				}
			}
		}

		for j := 0; j <= len(word); j++ {
			if j == 0 || dp[len(word)-j][len(word)-1] {
				substring := reversedWord[j:]
				if index, ok := m[substring]; ok && index != i {
					key := strconv.Itoa(i) + "-" + strconv.Itoa(index)
					if _, ok1 := added[key]; !ok1 {
						ans = append(ans, []int{i, index})
						added[key] = true
					}
				}
			}
		}
	}

	return ans
}

func constructDP(word string) [][]bool {
	dp := make([][]bool, len(word))
	for i := 0; i < len(word); i++ {
		dp[i] = make([]bool, len(word))
	}

	for i := 1; i <= len(word); i++ {
		for j := 0; j < len(word) && j+i-1 < len(word); j++ {
			k := i + j - 1
			if i == 1 {
				dp[j][k] = true
				continue
			} else if i == 2 {
				dp[j][k] = word[j] == word[k]
			} else {
				dp[j][k] = word[j] == word[k] && dp[j+1][k-1]
			}
		}
	}

	//fmt.Println(word, dp)
	return dp
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}