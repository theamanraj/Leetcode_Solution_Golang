func smallestEquivalentString(A string, B string, S string) string {
    parent := make(map[byte]byte)
	for i := 0; i < 26; i++ {
		parent[byte('a'+i)] = byte('a' + i)
	}
    
	for i := 0; i < len(A); i++ {
		l := find(A[i], parent)
		r := find(B[i], parent)
		if l < r {
			parent[r] = l
		}
		if l > r {
			parent[l] = r
		}
	}

	res := make([]byte, len(S))
	for i := 0; i < len(S); i++ {
		res[i] = find(S[i], parent)
	}
	return string(res)
}

func find(c byte, parent map[byte]byte) byte {
	if parent[c] != c {
		parent[c] = find(parent[c], parent)
	}
	return parent[c]
}