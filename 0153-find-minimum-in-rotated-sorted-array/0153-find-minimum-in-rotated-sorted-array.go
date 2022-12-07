func findMin(nums []int) int {
    left := 0
    right := len(nums)-1
    for left < right {
        mid := int(uint(left+right)>>1)
        if nums[mid] <= nums[right] {
            right = mid
        } else {
            left = mid+1
        }
    }
    return nums[left]
}
