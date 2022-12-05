func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var result []int
	for i := 0; i <= len(nums1)-1; i++ {
		var checked bool
		for j := 0; j <= len(nums2)-1; j++ {
			if nums1[i] == nums2[j] {
				for x := j; x <= len(nums2)-1; x++ {
					if nums2[x] > nums1[i] {
						result = append(result, nums2[x])
						checked = true
						break
					}
				}
				break
			}
		}
		if !checked {
			result = append(result, -1)
		}
	}
	return result
}