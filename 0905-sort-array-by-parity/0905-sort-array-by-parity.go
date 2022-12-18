func sortArrayByParity(nums []int) []int {
    start,end:=0, len(nums)-1
    for start<end {
        if nums[start]%2==1 && nums[end]%2==0 {
            nums[start],nums[end]=nums[end],nums[start]
            start++
            end--
        } else if nums[start]%2==0 {
            start++
        } else if nums[end]%2==1 {
            end--
        }
    }
    return nums
}