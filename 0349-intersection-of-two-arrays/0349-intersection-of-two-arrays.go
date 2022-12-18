func intersection(nums1 []int, nums2 []int) []int {

	inputToMap := make(map[int]int)
	var outputToArray []int

	for _, content := range nums1 {
		inputToMap[content] = 1
	}

	for _, content := range nums2 {
		if inputToMap[content] == 1 {
			outputToArray = append(outputToArray, content)
			inputToMap[content] = 2
		}
	}
	return outputToArray
}