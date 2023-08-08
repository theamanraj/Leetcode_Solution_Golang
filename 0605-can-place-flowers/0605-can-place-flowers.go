func canPlaceFlowers(flowerbed []int, n int) bool {
    count := 0
    i := 0
    for i < len(flowerbed) {
        if flowerbed[i] == 0 {
            if i == 0 || flowerbed[i-1] == 0 {
                if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
                    flowerbed[i] = 1
                    count++
                }
            }
        }
        i++
    }
    return count >= n
}