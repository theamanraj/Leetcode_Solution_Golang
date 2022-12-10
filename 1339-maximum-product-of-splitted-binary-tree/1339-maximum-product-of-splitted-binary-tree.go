/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxProduct(root *TreeNode) int {
    _ = helper(root)
    product := 0
    findMaxProduct(root, root.Val, &product)
    return product % 1000000007
}

func helper(node *TreeNode) int {
    if node == nil {
        return 0
    }
    node.Val += helper(node.Left) + helper(node.Right)
    return node.Val
}

func findMaxProduct(node *TreeNode, summary int, product *int) {
    if node.Left != nil {
        if current := node.Left.Val * (summary - node.Left.Val); current > *product {
            *product = current
        }
        findMaxProduct(node.Left, summary, product)
    }
    if node.Right != nil {
        if current := node.Right.Val * (summary - node.Right.Val); current > *product {
            *product = current
        }
        findMaxProduct(node.Right, summary, product)
    }
}