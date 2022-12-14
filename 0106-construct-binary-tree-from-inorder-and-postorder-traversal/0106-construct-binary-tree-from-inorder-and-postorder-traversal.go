/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    
    // Build the inorder val -> idx map
    inOrderMap := make(map[int]int)
    for i, e := range inorder { inOrderMap[e] = i }

	var helper func(int, int) *TreeNode

	helper = func(l, r int) *TreeNode {
        
        //check if we have subtrees
		if l > r { return nil }
        
        // Why Go can't Pop! :)
		root := &TreeNode{postorder[len(postorder)-1], nil, nil}
		postorder = postorder[:len(postorder)-1]

		idx := inOrderMap[root.Val]
		root.Right = helper(idx+1, r)
		root.Left = helper(l, idx-1)
		return root

	}

	return helper(0, len(postorder)-1)

}