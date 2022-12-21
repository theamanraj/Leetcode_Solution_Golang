func removeDuplicateLetters(s string) string {
	count := make(map[string]int,0)
	visited := make(map[string]bool)
	stack := make([]string,0)
	length := 0

	arr := strings.Split(s,"")
	for _, item := range arr {
		count[item]++
	}

for _, char := range arr {
	count[char]--
	if visited[char] {
		continue
	}

	for length > 0 && char < stack[length - 1] && count[stack[length-1]] > 0 {
		top := stack[length - 1]
		length--
		stack = stack[:length]
		visited[top] = false
	}
	stack = append(stack,char)
	visited[char] = true
	length++

}
return strings.Join(stack,"")
}