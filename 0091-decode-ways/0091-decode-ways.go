func ways(substr, orig string, count int, mem map[int]int) int {
	if val, ok := mem[count]; ok {
		return val
	}
	sum := 0
	if substr[0] == '0' {
		return 0
	}
	if count == len(orig) {
		return 1
	}
	sum = ways(orig[count:count+1], orig, count+1, mem)
	if count+2 <= len(orig) {
		if val, _ := strconv.Atoi(orig[count : count+2]); val >= 1 && val <= 26 {
			sum += ways(orig[count:count+2], orig, count+2, mem)
		}
	}
	mem[count] = sum
	return mem[count]
}
func numDecodings(s string) int {
	var mem = make(map[int]int)
	return ways(s, s, 0, mem)
}
