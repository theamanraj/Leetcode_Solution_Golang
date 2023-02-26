func replaceElements(arr []int) []int {
    max := -1
    for i := len(arr) - 1; i >= 0; i-- {
        if arr[i] > max {
            max, arr[i] = arr[i], max
        } else {
            arr[i] = max
        }
    }
    return arr
}