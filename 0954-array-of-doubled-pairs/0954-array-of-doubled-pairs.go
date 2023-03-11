func canReorderDoubled(nums []int) bool {
    m := map[int]int{}
    sort.Ints(nums)
    for _, num := range nums {
        if v, exists := m[num*2]; exists && v > 0 {
            m[num*2] = v-1
            continue
        } 
        if num %2 == 0 {
            if v, exists := m[num/2]; exists && v > 0 {
                m[num/2] = v-1
                continue
            }
        }
        if _, exists := m[num]; !exists {
            m[num] = 0
        }
        m[num] = m[num]+1
    }
    for _, v := range m {
        if v > 0 {
            return false
        }
    }
    return true
}