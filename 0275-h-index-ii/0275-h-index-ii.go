func hIndex(inp []int) int {
    n := len(inp)
    s, e := 0, len(inp)-1
    for s <= e {
        mid := s + (e-s)/2
        elem := inp[mid]
        if elem < n-mid {
            s = mid+1
        } else {
            e = mid-1
        }
    }
    
    return n-s
}