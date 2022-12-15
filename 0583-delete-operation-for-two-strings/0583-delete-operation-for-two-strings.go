func minDistance(word1 string, word2 string) int {

	if len(word1) == 0 || len(word2) ==0 {
		return len(word1) + len(word2)
	}

    word1 = " " + word1
	word2 = " " + word2
    
	dpArray := make([][] int, len(word1))
	for i:=0; i <len(word1); i++{
		dpArray[i] = make([]int, len(word2))
	}

	for i:=0; i<len(word1); i++{
		dpArray[i][0] = i
	}

	for i:=0; i<len(word2); i++{
		dpArray[0][i] = i
	}
	
	for i:=1; i<len(word1); i++{
		for j:=1; j<len(word2); j++ {
			if word1[i] != word2[j] {
				dpArray[i][j] = min(dpArray[i][j-1], dpArray[i-1][j]) + 1
			} else {
				dpArray[i][j] = dpArray[i-1][j-1]
			}
		}
	}

	return (dpArray[len(word1)-1][len(word2)-1])
}

func min(x int, y int) int{
	if x < y {
		return x
	}
	return y
}