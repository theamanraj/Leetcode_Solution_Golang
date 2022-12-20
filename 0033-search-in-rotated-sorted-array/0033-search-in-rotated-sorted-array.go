func search(nums []int, target int) int {
    n := len(nums)
    if n == 1 {
        if nums[0] == target {
            return 0
        }
        return -1
    }
    pivot := findPivot(nums)
    inLeft := binarySearchIntSlice(nums, 0, pivot - 1, target)
    inRight := binarySearchIntSlice(nums, pivot, n - 1, target)
    if inLeft != -1 {
        return inLeft
    }
    return inRight
}

func findPivot(nums []int) int {
    n := len(nums)
    low, high, mid := 1, n - 1, 0
    for low < high {
        mid = (low + high) / 2
        if nums[mid] > nums[0] {
            low = mid + 1 // this index can't be a part of answer
        } else {
            high = mid // this index can be the answer
        }
    }
    return low
}

func binarySearchIntSlice(nums []int, low int, high int, target int) int {
    mid := 0
    for low <= high {
        mid = (low + high) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return -1
}