/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    var ans [][]int
    dfs(root, 0, &ans)
    for i := 0; i < len(ans)/2; i++ {
        ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
    }
    return ans
}

func dfs(node *TreeNode, depth int, ans *[][]int) {
    if len(*ans) <= depth {
        *ans = append(*ans, []int{})
    }
    (*ans)[depth] = append((*ans)[depth], node.Val)
    if node.Left != nil {
        dfs(node.Left, depth+1, ans)
    }
    if node.Right != nil {
        dfs(node.Right, depth+1, ans)
    }
}