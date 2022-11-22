func findKthLargest(nums []int, k int) int {
    length := len(nums)
    recursive(&nums, 0, length-1, length-k)
    return nums[len(nums)-k]
}

func recursive(nums *[]int, start, end, k int) {
    if start >= end {
        return
    }
    
    pivotIndex := partition(nums, start, end)
    if pivotIndex == k {
        return
    }
    recursive(nums, start, pivotIndex-1, k)
    recursive(nums, pivotIndex+1, end, k)
}

func partition(nums *[]int, start, end int) int {
    pivot := (*nums)[end]
    for i:=start; i<end; i++ {
        if (*nums)[i] < pivot {
            swap(nums, start, i)
            start++
        }
    }
    swap(nums, start, end)
    return start
}

func swap(nums *[]int, i, j int) {
    (*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}