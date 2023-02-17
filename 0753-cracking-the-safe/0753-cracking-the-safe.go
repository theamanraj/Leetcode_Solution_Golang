const N1_STR = "0123456789"

var N, K int
var sb []byte // string builder
var seen map[string]bool

func crackSafe(n int, k int) string {
	if n == 1 {
		return N1_STR[:k]
	}
	N = n
	K = k
	seen = make(map[string]bool)
	total := int(math.Pow(float64(k), float64(n)))
	sb = make([]byte, 0, total+n-1)
	for i := 1; i != n; i++ {
		sb = append(sb, '0')
	}
	dfs(total)
	return string(sb)
}

func dfs(depth int) bool {
	if depth == 0 {
		return true // we did it
	}
	for i := 0; i != K; i++ {
		sb = append(sb, byte('0'+i))
		cur := string(sb[len(sb)-N:])
		if _, haveseen := seen[cur]; haveseen != true {
			seen[cur] = true
			if dfs(depth - 1) {
				return true // only in this case do we not delete
			} else {
				delete(seen, cur)
			}
		}
		sb = sb[0 : len(sb)-1] // del
	}
	return false
}