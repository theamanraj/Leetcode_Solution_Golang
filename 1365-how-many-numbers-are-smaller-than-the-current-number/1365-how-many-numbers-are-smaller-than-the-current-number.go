func smallerNumbersThanCurrent(nums []int) []int {
    var counts = make([]int, 0, len(nums))
    for i := 0; i < len(nums); i++ {
        var count int
        for j := 0; j < len(nums); j++ {
            if i != j && nums[i] > nums[j] {
                count++
            }
        }
        counts = append(counts, count)
    }
    
    return counts
}