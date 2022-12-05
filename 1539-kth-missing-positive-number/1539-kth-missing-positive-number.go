func findKthPositive(arr []int, k int) int {
    nums := make([]int, 0)
	n := arr[len(arr)-1]
	l := 0
	i := 0
	for i = 1; i <= n+k; i++ {
		if i != arr[l] {
			nums = append(nums, i)
		} else if i == arr[l] {
			l++
			if l >= len(arr){
				l = l-1
			}
		}
	}
    return nums[k-1]
}