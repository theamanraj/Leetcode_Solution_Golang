func findPairs(nums []int, k int) int {
    if k < 0 {
        return 0
    }
    res := 0
    m := make(map[int]int)
    for _, n := range nums{
        m[n] ++
    }    
    for key, count := range m {
        if k != 0 {
            if m[key + k] != 0 {
                res ++
            }
        } else {
            if count > 1 {
                res ++
            }
        }
    }
    return res
}