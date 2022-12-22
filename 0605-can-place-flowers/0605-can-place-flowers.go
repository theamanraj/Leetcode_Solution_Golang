func canPlaceFlowers(flowerbed []int, n int) bool {
    tmp := make([]int, len(flowerbed)+2)
    copy(tmp[1:], flowerbed)
    for i, empty := 0, 0; i < len(tmp); i++ {
        if tmp[i] == 1 {
            empty = 0
            continue
        }
        empty++
        if empty == 3 {
            n--
            empty = 1
        }
        if n == 0 {
            return true
        }
    }
    return false
}