/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverseInorder(node *TreeNode, list *[]int, result *int, prev *int) {
    if node == nil {
        return
    }

    traverseInorder(node.Left, list, result, prev)        
    // since it is a inorder traversal therefore prev will always represent the smaller element in the sorted traversal order
    *list = append(*list, node.Val)
    if *result > node.Val - *prev {
        *result = node.Val - *prev
    }
    *prev = node.Val
    traverseInorder(node.Right, list, result, prev)
}

func minDiffInBST(root *TreeNode) int {
    
    // just trying to practice golang skills dont worry about it
    // we dont need to store traversal to solve this problem
    inorderTraversal := make([]int, 0)
    result := 10001;
    prev := -10001
    traverseInorder(root, &inorderTraversal, &result, &prev)
    fmt.Println(inorderTraversal)

    return result
}