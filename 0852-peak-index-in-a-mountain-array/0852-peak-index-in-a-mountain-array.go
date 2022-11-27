func peakIndexInMountainArray(arr []int) int {
    max := 0
    index := 0
    if len(arr) < 3 {
        return 0
    } else if (arr[0] > arr[1]){
        return 0
    } else if (arr[len(arr)-1] > arr[len(arr)-2]){
        return 0
    }else {
        for i, v := range arr{
            if max < v {
                max = v
            }
            if arr[i] == max {
                index = i
            }
            
        }
        return index
    }
}