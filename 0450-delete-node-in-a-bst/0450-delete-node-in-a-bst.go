/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findNode(t, par *TreeNode, key int) (*TreeNode, *TreeNode) {
    if t == nil {
        return nil, par
    }
    
    if t.Val == key {
        return t, par
    }
    
    if t.Val < key {
        return findNode(t.Right, t, key)
    }
    
    return findNode(t.Left, t, key)
}

// adjustTree deletes the given node and adjusts the subtrees under the given node.
func adjustTree(t *TreeNode) *TreeNode {
    if t.Left == nil {
        return t.Right
    }
    
    if t.Right == nil {
        return t.Left
    }
    
    root := t.Left
    cur := root
    
    for cur.Right != nil {
        cur = cur.Right
    }
    
    cur.Right = t.Right
    
    return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
    cur, par := findNode(root, nil, key)
    if cur == nil {
        return root
    }
    
    if par == nil {
        return adjustTree(cur)
    }
    
    if par.Left == cur {
        par.Left = adjustTree(cur)
    } else {
        par.Right = adjustTree(cur)
    }
    
    return root
}