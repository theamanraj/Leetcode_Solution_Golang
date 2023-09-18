func removeDuplicates(nums []int) int {
    // create an hashmap
    mapInput := make(map[int]bool)
    l := 0
    for i := 0; i< len(nums);i++ {
        if !mapInput[nums[i]]{
            mapInput[nums[i]] = true
            nums[l] = nums[i]
            l++
        }
    }
    return l
}