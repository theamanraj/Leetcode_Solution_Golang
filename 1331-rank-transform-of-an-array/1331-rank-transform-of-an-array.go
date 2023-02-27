func arrayRankTransform(arr []int) []int {    
    arr2 := make([]int, len(arr))
    copy (arr2, arr)
    sort.Ints(arr2)
    m := make(map[int]int)
    for i, r := 0, 1; i < len(arr2); i++ {        
        if _, ok := m[arr2[i]]; !ok {
            m[arr2[i]] = r        
            r++
        }              
    }
    for i := range arr {
        arr[i] = m[arr[i]]
    }
    return arr
}