func findLengthOfShortestSubarray(arr []int) int {
	shortest := len(arr)
	// tailIndex: [tailIndex,len(arr)) non-decreasing
	// headIndex: [0,headIndex] non-decreasing
	tailIndex, headIndex := -1, -1
	for i := len(arr)-2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			shortest = i+1
			tailIndex = i+1
			break
		}
	}
	if tailIndex == -1 {
		return 0
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			cur := len(arr)-(i+1)
			if cur < shortest {
				shortest = cur
			}
			headIndex = i
			break
		}
	}
	edge := tailIndex
	for i := 0; i <= headIndex && edge < len(arr); {
		if arr[i] <= arr[edge] {
			cut := edge - i - 1
			if cut < shortest {
				shortest = cut
			}
			i++
		} else {
			edge++
		}
	}
	return shortest
}