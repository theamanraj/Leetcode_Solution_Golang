func findDisappearedNumbers(nums []int) []int {
    for i := range nums {
        for nums[nums[i]-1] != nums[i] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }
    
    ret := []int{}
    for i, v := range nums {
        if v != i+1 {
            ret = append(ret, i+1) 
        }
    }
    
    return ret
}