func maxChunksToSorted(arr []int) int {
    ls := len(arr)
    aux := make([]int, ls)
    for i := 0; i < ls; i++ {
        aux[i] = i
    } 
    // fmt.Println(aux)
    res, i := 0, 0
    for i < ls {
        start, end := i, i
        for start < ls {
            ref1, ref2 := arr[start:end+1], aux[start:end+1]
            sort.Ints(ref1)
            ok := helper(ref1, ref2)
            if ok {
                res++
                i = end + 1
                break
            }
            end ++
        }
    }
    return res
    return 1
}

func helper(arr1, arr2 []int) bool {
    ls := len(arr1)
    for i := 0; i< ls; i++ {
        if arr1[i] != arr2[i] {
            return false
        }
    }
    return true
}