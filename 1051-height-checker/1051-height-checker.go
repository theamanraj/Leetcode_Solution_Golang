func heightChecker(heights []int) int {
    count := 0
    expected := make([]int, len(heights))
    copy(expected, heights)
    sort.Ints(expected)
    for i:=0;i<len(heights);i++ {
        if heights[i] != expected[i] {
            count++
        }
    }
    return count
}