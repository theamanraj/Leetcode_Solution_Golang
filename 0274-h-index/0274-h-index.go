func hIndex(citations []int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(citations)))
    h:= 0;
    for i:=0; i<len(citations); i++ {
        if citations[i] >= i +1{
            h = i +1
        } else{
            break
        }
    }
    return h
}