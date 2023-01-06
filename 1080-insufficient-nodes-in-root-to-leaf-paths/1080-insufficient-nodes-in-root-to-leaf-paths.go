func sufficientSubset(root *TreeNode, limit int) *TreeNode {
  _, root = helper(root, 0, limit)
  return root
}

func helper(root *TreeNode, sum int, limit int) (int, *TreeNode) {
  sum += root.Val
  var maxVal int
  if root.Left == nil && root.Right == nil {
    // If this is the leaf node, the maxValue is going to be the sum so far
    maxVal = sum 
  } else {    
    // Otherwise, the max value is going to be the max of left and right subtree values
    leftMax, rightMax := math.MinInt32, math.MinInt32
    if root.Left != nil {
      leftMax, root.Left = helper(root.Left, sum, limit)
    }
    if root.Right != nil {
      rightMax, root.Right = helper(root.Right, sum, limit)
    }
    maxVal = max(leftMax, rightMax)
  }
  // If the max value of the root to leaf path sum is less than the limit, it means
  // all path sums are less than the limit
  if maxVal < limit {
    return maxVal, nil
  }
  return maxVal, root
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}