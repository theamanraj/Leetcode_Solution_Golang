func findMin(nums []int) int {
    l, r := 0, len(nums)-1
    
    for l < r {
        mid := l+(r-l)/2
        if nums[mid] < nums[r] {
            r = mid
        } else if nums[mid] > nums[r] {
            l = mid+1
        } else {
            r = r-1
        }
    }
    return nums[l]
}