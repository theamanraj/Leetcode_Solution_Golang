func arithmeticTriplets(nums []int, diff int) int {
	var l int = len(nums)
	var res int = 0
	var left int = 0
	var mid int = left + 1
	var right int = mid + 1
	for left < l && mid < l{
		for mid < l && (nums[mid] - nums[left]) < diff{
			mid++
		}
		if mid == l{
			break
		}
		if (nums[mid] - nums[left]) != diff{
			left++
			continue
		}
		right = max_int(right,mid + 1)
		for right < l && nums[right] - nums[mid] < diff{
			right++
		}
		if right == l{
			break
		}
		if (nums[right] - nums[mid]) == diff{
			res++
		}
		left++
	}
	return res
}

func max_int(a,b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}