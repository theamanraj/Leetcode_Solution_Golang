func isCapital(s byte) bool {
	return s >= 'A' && s <= 'Z'
}

func detectCapitalUse(word string) bool {
  	count := 0
	index := -1
	for i := 0; i < len(word); i++ {
		s := word[i]
		if isCapital(s) {
			count++
			index = i
		}
	}

	if count > 1 {
		return count == len(word)
	} else if count == 1 {
		return index == 0
	}
	return true  
}