func isLongPressedName(name string, typed string) bool {
	if name == typed {
		return true
	}

	var i, j, iCount, jCount int

	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}

		i, j, iCount, jCount = i+1, j+1, 1, 1
		// count the number of repeat bytes respectively.
		for ; i < len(name) && name[i] == name[i-1]; i, iCount = i+1, iCount+1 {}
		for ; j < len(typed) && typed[j] == typed[j-1]; j, jCount = j+1, jCount+1 {}
		
		// if name string has more repeat-bytes than typed string, typed isn't long enough to contain the repeat bytes of name.
		if iCount > jCount {
			return false
		}
	}
	
	// if any bytes left
	return i == len(name) && j == len(typed)
}