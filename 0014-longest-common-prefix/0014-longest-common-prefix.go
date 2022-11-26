func longestCommonPrefix(strs []string) string {
	newPrefix := ""
	for c, val := range strs {
		if c == 0 {
			newPrefix = val
		} else {
			prefix := ""
			for i := 0; i < len(val) && i < len(newPrefix); i++ {
				if newPrefix[i] == val[i] {
					prefix += string(val[i])
                } else {
                    break
                }
			}
			newPrefix = prefix
		}
	}
	return newPrefix
}