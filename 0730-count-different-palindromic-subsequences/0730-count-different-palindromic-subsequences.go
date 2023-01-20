type keys struct {
	char   string
	first  int
	second int
}

func countPalindromicSubsequences(s string) int {
	last := len(s) - 1
	letters := []rune{'a', 'b', 'c', 'd'}

	memo := make(map[keys]*big.Int)
	var helper func(char string, i int, j int) *big.Int

	helper = func(char string, i int, j int) *big.Int {
		if i > j {
			return big.NewInt(0)
		}

		check := keys{char, i, j}
		_, found := memo[check]
		if found == true {
			return memo[check]
		}

		ith := string(s[i])
		jth := string(s[j])

		if i == j && ith == char {
			return big.NewInt(1)
		}

		if ith == jth && ith == char {
			res := big.NewInt(2)
			for _, value := range letters {
				res.Add(res, helper(string(value), i+1, j-1))
			}
			memo[check] = res
			return res
		}

		if ith != char {
			return helper(char, i+1, j)
		}

		return helper(char, i, j-1)
	}
	result := big.NewInt(0)
	for _, character_rune := range letters {
		result.Add(result, helper(string(character_rune), 0, last))
	}
	test := math.Pow(10, 9) + 7
	base := big.NewInt(int64(test))
	temp_result := base.Mod(result, base)
	return int(temp_result.Int64())
}
