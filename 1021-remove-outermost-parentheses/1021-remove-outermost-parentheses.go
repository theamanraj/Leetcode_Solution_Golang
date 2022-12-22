func removeOuterParentheses(S string) string {
    c := ""
	counter := 0
	for _, v := range S {
        if v == '(' { 
			counter++
		} else {
			counter--
		}
		if !(counter == 0 && v == ')') && !(counter == 1 && v == '(') {
            c = c + string(v)
		}
	}
	return c
}