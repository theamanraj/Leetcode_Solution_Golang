/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
    dict := make(map[int]int)
    maxCnt :=0
    var dfs func(root *TreeNode)int
    dfs = func(root *TreeNode) int{
        if root == nil {return 0}
        sum := root.Val + dfs(root.Left) + dfs(root.Right)
        dict[sum] = dict[sum]+1
        maxCnt = max(maxCnt,dict[sum])
        return sum
    }
    dfs(root)
    res := make([]int,0)
    for k,v :=range dict{
        if  v == maxCnt{
            res = append(res,k)
        }
    }
    return res
}

func max(a int, b int) int{
    if a< b{return b}
    return a
}