/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    var stack []*TreeNode
    for root != nil {
        stack = append(stack, root)
        root = root.Left
    }
    return BSTIterator{stack}
}


func (this *BSTIterator) Next() int {
    res := this.stack[len(this.stack) - 1]
    this.stack = this.stack[:len(this.stack) - 1]
    if res.Right != nil {
        cur := res.Right
        for cur != nil {
            this.stack = append(this.stack, cur)
            cur = cur.Left
        }
    }
    return res.Val
}


func (this *BSTIterator) HasNext() bool {
    return len(this.stack) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */