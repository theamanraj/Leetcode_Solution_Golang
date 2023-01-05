func totalHammingDistance(nums []int) int {
    sum := 0
    for i:= 0; i < 32; i++ {
        var ones uint
        ones = 0
        for _, c := range nums {
            v := uint(c)
            j := uint(i)
            ones += ((v >> j) & 0x01)
        }
        zeros:=(len(nums) - int(ones))
        sum += zeros * int(ones)
    } 
    return sum
}