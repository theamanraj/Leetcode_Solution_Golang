/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt32
    maxPathSumHelper(root, &maxSum)
    return maxSum
}

func maxPathSumHelper(node *TreeNode, maxSum *int) int {
    if node == nil { return 0 }
    left, right := maxPathSumHelper(node.Left, maxSum), maxPathSumHelper(node.Right, maxSum)
    maxSumEndingAtNode := max(node.Val, node.Val + max(left, right))
    *maxSum = max(*maxSum, max(node.Val + left + right, maxSumEndingAtNode))
    return maxSumEndingAtNode
}

func max(x, y int) int { if x > y { return x }; return y }