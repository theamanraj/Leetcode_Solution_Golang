func maxNumber(nums1 []int, nums2 []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	res := []int{}
	for i := 0; i <= k; i++ {
		// get i numbers from nums1, get k - i numbers from nums2
		if len(nums1) < i || len(nums2) < k-i {
			continue
		}

		sub1, sub2 := getMaxSubArray(nums1, i), getMaxSubArray(nums2, k-i)
		temp := []int{}
		mergeTwo(&temp, sub1, sub2)
		if len(res) == 0 || bigger(temp, res) {
			res = temp
		}
	}
	return res
}

func getMaxSubArray(nums []int, k int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		for len(res) > 0 && nums[i] > res[len(res)-1] && len(res)+len(nums)-i-1 >= k {
			res = res[:len(res)-1]
		}
		res = append(res, nums[i])
	}

	for len(res) > k {
		res = res[:len(res)-1]
	}

	return res
}

func mergeTwo(res *[]int, nums1 []int, nums2 []int) {
	if len(nums1) == 0 && len(nums2) == 0 {
		return
	}

	if len(nums1) > 0 && bigger(nums1, nums2) {
		*res = append(*res, nums1[0])
		mergeTwo(res, nums1[1:], nums2)
	} else {
		*res = append(*res, nums2[0])
		mergeTwo(res, nums1, nums2[1:])
	}

}

func bigger(nums1 []int, nums2 []int) bool {
	i := 0
	for ; i < min(len(nums1), len(nums2)); i++ {
		if nums1[i] > nums2[i] {
			return true
		} else if nums1[i] < nums2[i] {
			return false
		}
	}

	return i != len(nums1)
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}