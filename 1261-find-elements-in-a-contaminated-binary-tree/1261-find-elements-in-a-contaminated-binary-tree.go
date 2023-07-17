/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    set map[int]bool
}

func (this *FindElements)constructTree(root *TreeNode, x int) {
    if root==nil {
        return
    }
    root.Val=x
    this.set[x]=true 
    this.constructTree(root.Left, (2*x)+1)
    this.constructTree(root.Right, (2*x)+2)
}
func Constructor(root *TreeNode) FindElements {
    m:=map[int]bool{}
    
    this:=FindElements{}
    this.set=m
    this.constructTree(root, 0)
    return this
}


func (this *FindElements) Find(target int) bool {
    return this.set[target];
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */