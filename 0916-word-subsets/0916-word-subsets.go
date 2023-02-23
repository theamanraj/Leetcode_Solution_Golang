func wordSubsets(A []string, B []string) []string {
    res := make([]string, 0)
    
    //generate the MIN superset of B
    minSuper := [26]int{}
    for i := 0; i < len(B); i++ {
        temp := conv(B[i])
        for i := 0; i < len(minSuper); i++ {
            minSuper[i] = max(minSuper[i], temp[i])
        }
    }
    
    //if A[i] is superset of minSuper, then append it into result
    for i := 0; i < len(A); i++ {
        temp := conv(A[i])
        valid := true
        for i := 0; i < len(temp); i++ {
            if temp[i] < minSuper[i] {
                valid = false
                break
            }
        }
        if valid {
            res = append(res, A[i])
        }
    }
    
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func conv(str string) [26]int {
    res := [26]int{}
    for i := 0; i < len(str); i++ {
        res[str[i] - 'a']++
    }
    return res
}