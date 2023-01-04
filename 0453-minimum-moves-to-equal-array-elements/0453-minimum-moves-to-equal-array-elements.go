func minMoves(nums []int) int {
    min := int(^uint(1<<63))
    for _, e := range nums {
        if e<min {
            min = e
        }
    }
    
    s := 0
    for _, e := range nums {
        s += e-min
    }
    return s
}