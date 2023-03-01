func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])
	var newArray []int
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			newArray = append(newArray, right[r])
			r++
		} else {
			newArray = append(newArray, left[l])
			l++
		}
	}
	if l < len(left) {
		newArray = append(newArray, left[l:]...)
	} else if r < len(right) {
		newArray = append(newArray, right[r:]...)
	}
	return newArray
}