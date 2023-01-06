func findSecondMinimumValue(root *TreeNode) ( res int ){
    res = math.MaxInt32
    var find bool
    dfs( root , root.Val, &res, &find)      
    if !find{ return -1}
    return
}


func dfs ( root *TreeNode, min int, res *int, find *bool ){
    if root==nil{ return}  
    if root.Val== min{
        dfs(root.Left, min, res, find)
        dfs(root.Right, min, res, find)
        return
    }  
    *find=true
    if root.Val>min && root.Val < *res{        
        *res= root.Val
    }
}