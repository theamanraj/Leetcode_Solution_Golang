func buddyStrings(A string, B string) bool {
	length := len(A)

	if length != len(B) {
		return false
	}

	diffs := 0
    // indexes of differences
	Ai, Bi := -1, -1
	for i := 0; i < length; i++ {
		if A[i] != B[i] {
			diffs++
			if Ai == -1 {
				Ai = i
			} else if Bi == -1 {
				Bi = i
			} else {
                // the third difference would return false
				return false
			}
		}
	}

	if diffs == 2 {
		// check wether 2 symbols from one string can replace
		// another two symbols from second string leaving both strings equal to each other
		return A[Ai] == B[Bi] && A[Bi] == B[Ai]
	}
    
    if diffs == 0 {
        //if strings are equal we should check wheather it consist two equievalent symbols
        flag := 0
        for i := 0; i < length; i++ {
            for z := 0; z < length; z++ {
                if A[i] == A[z] {
                    flag++
                }
                
                if flag == 2 {
                    return true
                }
            }
            //don't forget to drop flag after each iteration
            flag = 0
        }
    }
    
    return false
}