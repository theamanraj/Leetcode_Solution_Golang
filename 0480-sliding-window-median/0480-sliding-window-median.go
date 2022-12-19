func medianSlidingWindow(nums []int, k int) []float64 {
    //if nums is empty return empty
    if len(nums) == 0 {
        return []float64{}
    }
    //if k = 1 return nums (as array of floats)
    if k == 1 {
        res := make([]float64, 0)
        for i := 0; i< len(nums); i++ {
            res = append(res, float64(nums[i]))
        }
        return res
    }
    //initialize result []float64
    res := make([]float64, 0)
    //sort elements in indexes 0...k-1 and store as: arr int[]
    arr := make([]int,k)
    copy(arr, nums[0:k])
    sort.Ints(arr)
    //find the median of arr and insert into result array
    res = append(res, findMedian(arr))
    //from index i=1 to (len(nums)-k):
    for i:=1;i<=(len(nums)-k);i++ {
        // drop first element of arr ERROR: should only drop the first 
        arr = dropNumber(arr, nums[i-1])
        // insert nums[i+k-1] into correct position using binary search - O(log(n))
        arr = insertSort(arr, nums[i+k-1])
        //find the median of arr and insert into result array
        res = append(res, findMedian(arr))
    }
    //return result array
    return res
}
func findMedian(arr []int) float64 {
    if (len(arr) % 2) == 1 {
        return float64(arr[len(arr)/2])
    }
    return float64(arr[len(arr)/2-1] + arr[len(arr)/2])/2
}
func insertSort(data []int, el int) []int {
	index := sort.SearchInts(data, el)
	data = append(data, el)
	copy(data[index+1:], data[index:])
	data[index] = el
	return data
}
func dropNumber(arr []int, x int) []int {
    index := sort.SearchInts(arr, x)
    return append(arr[:index], arr[(index+1):]...)
}