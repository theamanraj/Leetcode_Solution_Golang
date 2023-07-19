func maxLength(arr []string) int {
    m := make(map[string]int)
    for _, a := range arr {
        m[a] = getSignature(a)
    }
    max := 0
    explore(arr, 0, 0, m, &max)
    return max
}

func getSignature(s string) int {
    var tmp [26]int
    for _, r := range []rune(s) {
        tmp[r-'a']++
        if tmp[r-'a'] > 1 {
            return 0
        }
    }
    ret := 0
    for i := 0; i < 26; i++ {
        ret = ret << 1 + tmp[i]
    }
    return ret
}

func explore(arr []string, index int, buf int, m map[string]int, max *int) {
    if index == len(arr) {
        if t := bits.OnesCount(uint(buf)); t > *max {
            *max = t
        }
        return
    }
    sig := m[arr[index]]
    // try not include this one
    explore(arr, index+1, buf, m, max)
    if buf & sig == 0 {
        explore(arr, index+1, buf | sig, m, max)
    }
}