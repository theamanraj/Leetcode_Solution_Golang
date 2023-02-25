func numRabbits(answers []int) int {
    sum := 0
    m := make(map[int]int)
    for _, ans := range answers {
        m[ans+1]++
    }

    for k, v := range m {
        s := v/k
        if v%k > 0 {
            s++
        }
        
        sum += s * k
    }
    return sum
}