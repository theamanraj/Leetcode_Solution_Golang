func checkIfExist(arr []int) bool {
    hash := map[int]bool{}
    for _, i := range arr {
        if hash[i*2] || (hash[i/2] && i%2 == 0) {
            return true
        }
        hash[i] = true
    }
    return false
}