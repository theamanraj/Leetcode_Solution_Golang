func duplicateZeros(arr []int)  {
    for i := 0; i < len(arr); i++ {
	if arr[i] == 0 {
		arr = arr[:len(arr)-1]
		var sub []int = append(arr[:i], 0)
		arr = append(sub, arr[i:]...)
		i++
	}
}
}