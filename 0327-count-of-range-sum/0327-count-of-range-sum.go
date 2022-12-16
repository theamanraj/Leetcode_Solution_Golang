func mergeSort(nums []int, start, end, low, high int) int {
	if start == end {
		return 0
	}
	if start == end -1 {
		count := 0
		if nums[end] - nums[start] <= high && nums[end] - nums[start] >= low {
			count++
		}
		if nums[start] > nums[end] {
			nums[start], nums[end] = nums[end], nums[start]
		}
		return count
	}
	mid := start + (end - start) / 2
	count := mergeSort(nums, start, mid, low, high)
	count += mergeSort(nums, mid+1, end, low, high)
	p, q := mid+1, mid+1
	for i:= start; i<=mid;i++ {
		// first item greater than or equal to nums[i] + low
		for p <= end && nums[i] + low > nums[p] {
			p++
		}
		// first item greater than nums[i]+high
		for q <= end && nums[i] + high >= nums[q] {
			q++
		}
		if p != q {
			count+= q - p
		}
	}
	temp := []int{}
	for i, j:= start, mid+1; i <= mid || j <= end; {
		if i > mid {
			temp = append(temp, nums[j])
			j++
		} else if j > end {
			temp = append(temp, nums[i])
			i++
		} else if nums[i] < nums[j] {
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[j])
			j++
		}
	}
	copy(nums[start:], temp)
	return count
}

func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}
	for i:= 1; i < len(nums);i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	nums = append([]int{0}, nums...)
	return mergeSort(nums, 0, len(nums)-1, lower, upper)
}