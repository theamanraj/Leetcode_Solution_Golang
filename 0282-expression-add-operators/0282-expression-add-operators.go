func addOperators(num string, target int) []string {
	res := make([]string, 0)
	if len(num) == 0 {
		return res
	}

	path := make([]rune, len(num)*2-1)
	digits := []rune(num)
	n := 0
	for i := 0; i < len(digits); i++ {
		n = n*10 + int(digits[i]) - '0'
		path[i] = digits[i]
		if i < len(path) {
			path[i+1] = ' '
		}
		res = dfs(res, path, i+1, 0, n, digits, i+1, target)
		if n == 0 {
			break
		}
	}

	return res
}

func dfs(res []string, path []rune, length int, left int, cur int, digits []rune, pos int, target int) []string {
	if pos == len(digits) {
		if left+cur == target {
			addpath := string(path)
			addpath = strings.Replace(addpath, " ", "", -1)
			res = append(res, string(addpath))
		}
		return res
	}

	n := 0
	j := length + 1
	for i := pos; i < len(digits); i++ {
		n = n*10 + int(digits[i]) - '0'
		path[j] = digits[i]
		j++
		if j < len(path) {
			path[j] = ' '
		}

		path[length] = '+'
		res = dfs(res, path, j, left+cur, n, digits, i+1, target)

		path[length] = '-'
		res = dfs(res, path, j, left+cur, -n, digits, i+1, target)

		path[length] = '*'
		res = dfs(res, path, j, left, cur*n, digits, i+1, target)

		if digits[pos] == '0' {
			break
		}
	}

	return res
}