func findSubsequences(nums []int) [][]int {
    var res [][]int
    same := func(t []int) bool{
        for _, r := range res {
            dup := true
            if len(t) != len(r) {
                dup = false
                continue
            }
            for i := 0; i < len(r); i ++{
                if r[i] != t[i] {
                    dup = false
                    break
                }
            }
            if dup {
                return true
            }
        }
        return false
    } 
	
    var DFS func(int, []int)
    DFS = func(n int, list []int) {
        if n >= len(nums) {
            var tmp []int
            if len(list) > 1 && !same(list) {
                res = append(res, append(tmp, list...))
            }
            return 
        }
        DFS(n + 1, list)
        if len(list) > 0 && nums[n] >= list[len(list) - 1] || len(list) == 0 {
            DFS(n + 1, append(list, nums[n]))
        }
    }
	
    DFS(0, nil)
    return res
}