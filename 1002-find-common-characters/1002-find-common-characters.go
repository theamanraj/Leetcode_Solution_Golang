func commonChars(A []string) []string {
	var res []string
	if len(A) == 0 {
		return []string{}
	} else if len(A) == 1 {
		for i := 0; i < len(A[0]); i++ {
			res = append(res, fmt.Sprintf("%c", A[0][i]))
		}
	}

	pre := commonChars(A[:len(A)-1])
	cur := A[len(A)-1]

	for i := 0; i < len(pre); i++ {
		index := strings.Index(cur, pre[i])
		if index != -1 {
			res = append(res, pre[i])
			cur = fmt.Sprintf("%s#%s", cur[:index], cur[index+1:])
		}
	}
	return res
}