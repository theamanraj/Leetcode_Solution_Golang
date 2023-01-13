func beautifulArray(N int) []int {
	beautArray1024 := make([]int, 0, 1024)
	beautArray1024 = append(beautArray1024, 1, 2)
	for i := 2; i <= 1024; i*=2 {
		temp := make([]int, len(beautArray1024))
		copy(temp, beautArray1024)
		for index, value := range temp {
			temp[index] = 2*value
		}
		beautArray1024 = append(beautArray1024, temp...)
		for index := range temp {
			beautArray1024[index] = temp[index] - 1
		}
	}
	beautArray := make([]int, 0, N)
	for _, v := range beautArray1024 {
		if v <= N {
			beautArray = append(beautArray, v)
		}
	}
	return beautArray
}