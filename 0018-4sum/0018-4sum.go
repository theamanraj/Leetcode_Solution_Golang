func fourSum(nums []int, target int) [][]int {

	result := make([][]int, 0)
	if len(nums) < 4 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
        if i >0 && nums[i] == nums[i-1]{
            continue
        }
		for j := i+1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			low := j + 1
			high := len(nums) - 1
			for low < high {
				sum := nums[i] + nums[j]+ nums[low] + nums[high]
				if sum == target {
					arr := []int{nums[i],nums[j],nums[low], nums[high]}
					result = append(result, arr)
					low++
					high--
					for low < high && nums[low] == nums[low-1] {
						low++
					}

					for low < high && nums[high] == nums[high+1] {
						high--
					}
				} else if sum > target {
					high--
				} else {
					low++
				}
			}

		}
	}
	return result
}