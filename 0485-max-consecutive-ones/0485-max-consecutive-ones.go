func findMaxConsecutiveOnes(nums []int) int {

	A, B := 0, 0
	
	for _, v := range nums {
		if A + v > B {
			B = A + v
		}
        A = (A + v) * v
	}
	
	return B

}