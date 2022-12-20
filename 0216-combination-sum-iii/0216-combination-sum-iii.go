func combinationSum3(k int, n int) [][]int {
    result := make([][]int, 0)

    backtrack([]int{1,2,3,4,5,6,7,8,9}, []int{}, 0, k, n, &result)

    return result
}

func backtrack(candidates []int, comb []int, curr int, max int, remain int, result *[][]int) {
    if remain == 0 && len(comb) == max {
        *result = append(*result, append([]int{}, comb...))
        return
    }
    
    for i:=curr; i<len(candidates); i++ {
        if len(comb) >= max {
            break
        }
        
        if candidates[i] > remain {
            break
        }
        
        comb = append(comb, candidates[i])
        backtrack(candidates, comb, i + 1, max, remain - candidates[i], result)
        comb = comb[:len(comb)-1]
    }
}