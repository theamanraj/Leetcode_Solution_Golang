var ans [][]int
func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    ans = [][]int{}
    Check(candidates, target, []int{})
    return ans
}

func Check(d []int, t int, c []int) {
    if t == 0 {
        ans = append(ans, c)
        return
    }
    for i, v := range d {
        if v == t {
            ans = append(ans, append(c, v))
            return
        }
        if v > t {
            return
        }
        tt := []int{}
        tt = append(tt, c...)
        tt = append(tt, v)
        Check(d[i:], t-v, tt)
    }
}