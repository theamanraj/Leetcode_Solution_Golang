import (
	"strconv"
)

func max(a int, b int) int {
    if a >= b {
        return a
    }
    
    return b
}

func maximalNetworkRank(n int, roads [][]int) int {
    graph := make([][]int, n)
    pairs:= make(map[string]bool)
    ranks := make([]int, n)

    for _, pair := range roads {
        i := pair[0]
        j := pair[1]
        a := strconv.Itoa(i)
        b := strconv.Itoa(j)
        graph[i] = append(graph[i], j)
        graph[j] = append(graph[j], i)
        pairs[a + ":" + b] = true
        pairs[b + ":" + a] = true
    }
    
    for i := 0; i < n; i++ {
        ranks[i] = len(graph[i])
    }
    
    ans := 0
    
    for i := 0; i < n; i++ {
        a := strconv.Itoa(i)

        for j := i + 1; j < n; j++ {  
            b := strconv.Itoa(j)

            res := ranks[i] + ranks[j]
            
            if pairs[a + ":" + b] || pairs[b + ":" + a] {
                res--
            }
            
            ans = max(ans, res)
        }
    }
    
    return ans
}